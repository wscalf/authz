// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1alpha/core.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CheckPermissionClient is the client API for CheckPermission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckPermissionClient interface {
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
}

type checkPermissionClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckPermissionClient(cc grpc.ClientConnInterface) CheckPermissionClient {
	return &checkPermissionClient{cc}
}

func (c *checkPermissionClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.CheckPermission/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckPermissionServer is the server API for CheckPermission service.
// All implementations should embed UnimplementedCheckPermissionServer
// for forward compatibility
type CheckPermissionServer interface {
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
}

// UnimplementedCheckPermissionServer should be embedded to have forward compatible implementations.
type UnimplementedCheckPermissionServer struct {
}

func (UnimplementedCheckPermissionServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}

// UnsafeCheckPermissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckPermissionServer will
// result in compilation errors.
type UnsafeCheckPermissionServer interface {
	mustEmbedUnimplementedCheckPermissionServer()
}

func RegisterCheckPermissionServer(s grpc.ServiceRegistrar, srv CheckPermissionServer) {
	s.RegisterService(&CheckPermission_ServiceDesc, srv)
}

func _CheckPermission_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckPermissionServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.CheckPermission/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckPermissionServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CheckPermission_ServiceDesc is the grpc.ServiceDesc for CheckPermission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CheckPermission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha.CheckPermission",
	HandlerType: (*CheckPermissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckPermission",
			Handler:    _CheckPermission_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha/core.proto",
}

// LicenseServiceClient is the client API for LicenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LicenseServiceClient interface {
	GetLicense(ctx context.Context, in *GetLicenseRequest, opts ...grpc.CallOption) (*GetLicenseResponse, error)
	ModifySeats(ctx context.Context, in *ModifySeatsRequest, opts ...grpc.CallOption) (*ModifySeatsResponse, error)
	GetSeats(ctx context.Context, in *GetSeatsRequest, opts ...grpc.CallOption) (*GetSeatsResponse, error)
	EntitleOrg(ctx context.Context, in *EntitleOrgRequest, opts ...grpc.CallOption) (*EntitleOrgResponse, error)
}

type licenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLicenseServiceClient(cc grpc.ClientConnInterface) LicenseServiceClient {
	return &licenseServiceClient{cc}
}

func (c *licenseServiceClient) GetLicense(ctx context.Context, in *GetLicenseRequest, opts ...grpc.CallOption) (*GetLicenseResponse, error) {
	out := new(GetLicenseResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.LicenseService/GetLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) ModifySeats(ctx context.Context, in *ModifySeatsRequest, opts ...grpc.CallOption) (*ModifySeatsResponse, error) {
	out := new(ModifySeatsResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.LicenseService/ModifySeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) GetSeats(ctx context.Context, in *GetSeatsRequest, opts ...grpc.CallOption) (*GetSeatsResponse, error) {
	out := new(GetSeatsResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.LicenseService/GetSeats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *licenseServiceClient) EntitleOrg(ctx context.Context, in *EntitleOrgRequest, opts ...grpc.CallOption) (*EntitleOrgResponse, error) {
	out := new(EntitleOrgResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.LicenseService/EntitleOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LicenseServiceServer is the server API for LicenseService service.
// All implementations should embed UnimplementedLicenseServiceServer
// for forward compatibility
type LicenseServiceServer interface {
	GetLicense(context.Context, *GetLicenseRequest) (*GetLicenseResponse, error)
	ModifySeats(context.Context, *ModifySeatsRequest) (*ModifySeatsResponse, error)
	GetSeats(context.Context, *GetSeatsRequest) (*GetSeatsResponse, error)
	EntitleOrg(context.Context, *EntitleOrgRequest) (*EntitleOrgResponse, error)
}

// UnimplementedLicenseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLicenseServiceServer struct {
}

func (UnimplementedLicenseServiceServer) GetLicense(context.Context, *GetLicenseRequest) (*GetLicenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLicense not implemented")
}
func (UnimplementedLicenseServiceServer) ModifySeats(context.Context, *ModifySeatsRequest) (*ModifySeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeats not implemented")
}
func (UnimplementedLicenseServiceServer) GetSeats(context.Context, *GetSeatsRequest) (*GetSeatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeats not implemented")
}
func (UnimplementedLicenseServiceServer) EntitleOrg(context.Context, *EntitleOrgRequest) (*EntitleOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EntitleOrg not implemented")
}

// UnsafeLicenseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LicenseServiceServer will
// result in compilation errors.
type UnsafeLicenseServiceServer interface {
	mustEmbedUnimplementedLicenseServiceServer()
}

func RegisterLicenseServiceServer(s grpc.ServiceRegistrar, srv LicenseServiceServer) {
	s.RegisterService(&LicenseService_ServiceDesc, srv)
}

func _LicenseService_GetLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).GetLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.LicenseService/GetLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).GetLicense(ctx, req.(*GetLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_ModifySeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).ModifySeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.LicenseService/ModifySeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).ModifySeats(ctx, req.(*ModifySeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_GetSeats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSeatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).GetSeats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.LicenseService/GetSeats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).GetSeats(ctx, req.(*GetSeatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LicenseService_EntitleOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntitleOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LicenseServiceServer).EntitleOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.LicenseService/EntitleOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LicenseServiceServer).EntitleOrg(ctx, req.(*EntitleOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LicenseService_ServiceDesc is the grpc.ServiceDesc for LicenseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LicenseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha.LicenseService",
	HandlerType: (*LicenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLicense",
			Handler:    _LicenseService_GetLicense_Handler,
		},
		{
			MethodName: "ModifySeats",
			Handler:    _LicenseService_ModifySeats_Handler,
		},
		{
			MethodName: "GetSeats",
			Handler:    _LicenseService_GetSeats_Handler,
		},
		{
			MethodName: "EntitleOrg",
			Handler:    _LicenseService_EntitleOrg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha/core.proto",
}

// ImportServiceClient is the client API for ImportService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImportServiceClient interface {
	ImportOrg(ctx context.Context, in *ImportOrgRequest, opts ...grpc.CallOption) (*ImportOrgResponse, error)
}

type importServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImportServiceClient(cc grpc.ClientConnInterface) ImportServiceClient {
	return &importServiceClient{cc}
}

func (c *importServiceClient) ImportOrg(ctx context.Context, in *ImportOrgRequest, opts ...grpc.CallOption) (*ImportOrgResponse, error) {
	out := new(ImportOrgResponse)
	err := c.cc.Invoke(ctx, "/api.v1alpha.ImportService/ImportOrg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImportServiceServer is the server API for ImportService service.
// All implementations should embed UnimplementedImportServiceServer
// for forward compatibility
type ImportServiceServer interface {
	ImportOrg(context.Context, *ImportOrgRequest) (*ImportOrgResponse, error)
}

// UnimplementedImportServiceServer should be embedded to have forward compatible implementations.
type UnimplementedImportServiceServer struct {
}

func (UnimplementedImportServiceServer) ImportOrg(context.Context, *ImportOrgRequest) (*ImportOrgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportOrg not implemented")
}

// UnsafeImportServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImportServiceServer will
// result in compilation errors.
type UnsafeImportServiceServer interface {
	mustEmbedUnimplementedImportServiceServer()
}

func RegisterImportServiceServer(s grpc.ServiceRegistrar, srv ImportServiceServer) {
	s.RegisterService(&ImportService_ServiceDesc, srv)
}

func _ImportService_ImportOrg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportOrgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImportServiceServer).ImportOrg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1alpha.ImportService/ImportOrg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImportServiceServer).ImportOrg(ctx, req.(*ImportOrgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImportService_ServiceDesc is the grpc.ServiceDesc for ImportService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImportService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha.ImportService",
	HandlerType: (*ImportServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportOrg",
			Handler:    _ImportService_ImportOrg_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1alpha/core.proto",
}
