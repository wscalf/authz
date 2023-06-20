package application

import (
	core "authz/api/gen/v1alpha"
	"authz/domain"
	"authz/domain/contracts"
	"authz/domain/services"
	"context"
	"sync/atomic"

	"github.com/golang/glog"
)

// LicenseAppService the handler for seat related endpoints.
type LicenseAppService struct {
	accessRepo    contracts.AccessRepository
	seatRepo      contracts.SeatLicenseRepository
	principalRepo contracts.PrincipalRepository
	subjectRepo   contracts.SubjectRepository
	orgRepo       contracts.OrganizationRepository
	ctx           context.Context
}

// GetSeatAssignmentRequest represents a request to get the users assigned seats on a license
type GetSeatAssignmentRequest struct {
	Requestor    string
	OrgID        string
	ServiceID    string
	IncludeUsers bool
	Assigned     bool
}

// ModifySeatAssignmentRequest represents a request to assign and/or unassign seat licenses
type ModifySeatAssignmentRequest struct {
	Requestor string
	OrgID     string
	ServiceID string
	Assign    []string
	Unassign  []string
}

// GetSeatAssignmentCountsRequest represents a request to get the seats limit and current allocation for a license
type GetSeatAssignmentCountsRequest struct {
	Requestor string
	OrgID     string
	ServiceID string
}

// OrgEntitledEvent represents an event where an organization has been entitled with a new license
type OrgEntitledEvent struct {
	OrgID     string
	ServiceID string
	MaxSeats  int
}

// ImportOrgEvent triggers new user import for an org
type ImportOrgEvent struct {
	OrgID string
}

// ImportUsersResult contains counters for imported and not imported users.
type ImportUsersResult struct {
	importedUsersCount    uint64
	notImportedUsersCount uint64
}

// NewLicenseAppService ctor.
func NewLicenseAppService(accessRepo contracts.AccessRepository, seatRepo contracts.SeatLicenseRepository, principalRepo contracts.PrincipalRepository, subjectRepo contracts.SubjectRepository, orgRepo contracts.OrganizationRepository) *LicenseAppService {
	return &LicenseAppService{
		accessRepo:    accessRepo,
		seatRepo:      seatRepo,
		principalRepo: principalRepo,
		subjectRepo:   subjectRepo,
		orgRepo:       orgRepo,
		ctx:           context.Background(),
	}
}

// GetSeatAssignmentCounts gets the seat limit and current allocation for a license
func (s *LicenseAppService) GetSeatAssignmentCounts(req GetSeatAssignmentCountsRequest) (limit int, available int, err error) {
	evt := domain.GetLicenseEvent{
		OrgID:     req.OrgID,
		ServiceID: req.ServiceID,
	}

	evt.Requestor = domain.SubjectID(req.Requestor)

	seatsService := services.NewSeatLicenseService(s.seatRepo, s.accessRepo)

	lic, err := seatsService.GetLicense(evt)
	if err != nil {
		return 0, 0, err
	}

	limit = lic.MaxSeats
	available = lic.GetAvailableSeats()
	err = nil
	return
}

// GetSeatAssignments gets the subjects assigned to seats in a license
func (s *LicenseAppService) GetSeatAssignments(req GetSeatAssignmentRequest) ([]domain.Principal, error) {
	evt := domain.GetLicenseEvent{
		OrgID:     req.OrgID,
		ServiceID: req.ServiceID,
	}

	evt.Requestor = domain.SubjectID(req.Requestor)

	seatService := services.NewSeatLicenseService(s.seatRepo, s.accessRepo)

	var resultIds []domain.SubjectID
	var err error
	if req.Assigned {
		resultIds, err = seatService.GetAssignedSeats(evt)
	} else {
		resultIds, err = seatService.GetAssignableSeats(evt)
	}

	if err != nil {
		return nil, err
	}

	if req.IncludeUsers {
		return s.principalRepo.GetByIDs(resultIds)
	}

	principals := make([]domain.Principal, len(resultIds))
	for i, id := range resultIds {
		principals[i] = domain.Principal{ID: id}
	}
	return principals, nil
}

// ModifySeats Assign and/or unassign a number of users for a given org and service
func (s *LicenseAppService) ModifySeats(req ModifySeatAssignmentRequest) error {
	evt := domain.ModifySeatAssignmentEvent{
		Org:     domain.Organization{ID: req.OrgID},
		Service: domain.Service{ID: req.ServiceID},
	}

	evt.Requestor = domain.SubjectID(req.Requestor)

	evt.Assign = make([]domain.SubjectID, len(req.Assign))
	for i, id := range req.Assign {
		evt.Assign[i] = domain.SubjectID(id)
	}

	evt.UnAssign = make([]domain.SubjectID, len(req.Unassign))
	for i, id := range req.Unassign {
		evt.UnAssign[i] = domain.SubjectID(id)
	}

	seatService := services.NewSeatLicenseService(s.seatRepo, s.accessRepo)

	return seatService.ModifySeats(evt)
}

// ProcessOrgEntitledEvent handles the OrgEntitledEvent by storing the license and importing users
func (s *LicenseAppService) ProcessOrgEntitledEvent(evt OrgEntitledEvent) error {

	err := s.seatRepo.ApplyLicense(&domain.License{
		OrgID:     evt.OrgID,
		ServiceID: evt.ServiceID,
		MaxSeats:  evt.MaxSeats,
		Version:   "",
		InUse:     0,
	})

	if err != nil {
		return err
	}

	// always run import.
	_, e := s.importUsers(evt.OrgID)
	if e != nil {
		return err
	}
	return nil
}

// ImportUsersForOrg imports users for a given orgID and returns a result containing a count of imported and not imported users
func (s *LicenseAppService) ImportUsersForOrg(evt ImportOrgEvent) (*core.ImportOrgResponse, error) {
	// always run import.
	result, err := s.importUsers(evt.OrgID)
	if err != nil {
		return nil, err
	}

	return &core.ImportOrgResponse{
		ImportedUsersCount:    result.importedUsersCount,
		NotImportedUsersCount: result.notImportedUsersCount,
	}, nil
}

func (s *LicenseAppService) importUsers(orgID string) (*ImportUsersResult, error) {
	// always run import.
	subjects, errors := s.subjectRepo.GetByOrgID(orgID)

	var importedUsersCount uint64
	var failedUserImportCount uint64

loop:
	for {
		select {
		case subject, ok := <-subjects:
			if ok {
				err := s.orgRepo.AddSubject(orgID, subject)
				if err != nil {
					atomic.AddUint64(&failedUserImportCount, 1)
					glog.Errorf("Failed to import user %s to org %s", subject.SubjectID, orgID)

					if errorShouldBeRetried(err) { // TODO: add test to test 'true' path
						return nil, err // TODO: add retry mechanism (but for now it's fine to bomb out and retry the whole processing of the event since all ops are idempotent)
					}
				} else {
					atomic.AddUint64(&importedUsersCount, 1)
				}
			} else {
				break loop
			}
		case err, ok := <-errors:
			if !ok {
				break loop
			}

			glog.Errorf(err.Error()) // TODO: think more about the contract. Is it possible to reason about individual errors in a channel and what they refer to and whether we should stop or continue?

			return nil, err
		}
	}

	return &ImportUsersResult{
		importedUsersCount:    importedUsersCount,
		notImportedUsersCount: failedUserImportCount,
	}, nil
}

func errorShouldBeRetried(err error) bool {
	return err != domain.ErrSubjectAlreadyExists
}
