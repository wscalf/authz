baseUri=$1
idpDiscoveryUri=$2
token=""

client_id=cloud-services
scopes=openid
port=8000

function login() {
    metadata=$(curl -s "$idpDiscoveryUri")
    authzEndpoint=$(echo "$metadata" | jq -r ".authorization_endpoint")
    tokenEndpoint=$(echo "$metadata" | jq -r ".token_endpoint")

    echo "Please open this url in a browser: $authzEndpoint?response_type=code&client_id=$client_id&redirect_uri=http://127.0.0.1:$port/&scope=$scopes"
    
    echo "Listening for response from identity provider.."
    #The response that should be sent to the browser gets piped into netcat, which listens on the given port and send the response to the first connection, then grep searches the request for code=yourcodehere to extract the authorization code)
    authzcode=$(echo -e "HTTP/1.1 200\r\nContent-Length: 26\r\nContent-Type: text/plain\r\nConnection: close\r\n\r\nPlease return to terminal." | nc -l 8000 | grep -oP "(?<=code=)[^&\s]+")

    echo "Exchanging code for token"
    token_response=$(curl -s -X POST -d "grant_type=authorization_code&client_id=$client_id&code=$authzcode&redirect_uri=http://127.0.0.1:$port/" $tokenEndpoint)
    token=$(echo "$token_response" | jq -r '.access_token') #Parse the access token out of the response and set it to a global
}

function fail() {
    echo "$1"
    exit 1
}

function assert() {
    condition=$1
    message=$2

    if [ ! $condition ]; then
        fail "FAIL! Condition: $condition Detail: $message"
    fi
}

function cleanup() {
    #always unassign synthetic_test on exit to avoid pollution on the next run. Also ignore all output- this will fail except when the test is exiting early before the second quantization interval
    curl -s -X POST $baseUri/v1alpha/orgs/o1/licenses/smarts -H "Origin: http://smoketest.test" -H "Content-Type: application/json" -H "Authorization:Bearer $token" -d '{"unassign": ["synthetic_test"]}' > /dev/null
}

login

msg='Granting license to synthetic_test (should succeed)'
curl --fail -X POST $baseUri/v1alpha/orgs/o1/licenses/smarts -H "Origin: http://smoketest.test" -H "Content-Type: application/json" -H "Authorization:Bearer $token" -d '{"assign": ["synthetic_test"]}' || fail "Failed request: $msg"

trap cleanup EXIT #once the synthetic_test user has been assigned, ensure unassign gets called on exit. This will always be redundant (and cause an error) on successful runs but ensures that it's not left behind on failure.

msg='Getting number of seats available - should be less than the license allows'
previousAvailable=`( curl --silent --fail $baseUri/v1alpha/orgs/o1/licenses/smarts -H "Origin: http://smoketest.test" -H "Authorization:Bearer $token" || fail "Failed request: $msg") | jq ".seatsAvailable"`
assert "$previousAvailable -lt 10" "$msg"

echo "Waiting for quantization interval"
sleep 5

msg='Checking access for synthetic_test (should succeed)'
ret=`( curl --silent --fail -X POST $baseUri/v1alpha/check -H "Origin: http://smoketest.test" -H "Content-Type: application/json" -H "Authorization:Bearer $token" -d '{"subject": "synthetic_test", "operation": "access", "resourcetype": "license", "resourceid": "o1/smarts"}' || fail "Failed request: $msg" ) | jq ".result"`
assert "$ret = true" "$msg"

msg='Checking if synthetic_test is included in the list of assigned users'
ret=`( curl --silent --fail -H "Origin: http://smoketest.test" -H "Authorization:Bearer $token" $baseUri/v1alpha/orgs/o1/licenses/smarts/seats || fail "Failed request: $msg") | jq 'any(.users[]; .id == "synthetic_test")'`
assert "$ret = true" "$msg"

msg='Revoking license for synthetic_test (should succeed)'
curl --fail -X POST $baseUri/v1alpha/orgs/o1/licenses/smarts -H "Origin: http://smoketest.test" -H "Content-Type: application/json" -H "Authorization:Bearer $token" -d '{"unassign": ["synthetic_test"]}' || fail "Failed request: $msg"

echo "Waiting for quantization interval"
sleep 5

msg='Getting license counts again - one more should be available'
newAvailable=`( curl --silent --fail $baseUri/v1alpha/orgs/o1/licenses/smarts -H "Origin: http://smoketest.test" -H "Authorization:Bearer $token" || fail "Failed request: $msg" ) | jq ".seatsAvailable"`
assert "$previousAvailable -lt $newAvailable" "$msg"

msg="Checking access for synthetic_test again (should return false)"
ret=`( curl --silent --fail -X POST $baseUri/v1alpha/check -H "Origin: http://smoketest.test" -H "Content-Type: application/json" -H "Authorization:Bearer $token" -d '{"subject": "synthetic_test", "operation": "access", "resourcetype": "license", "resourceid": "o1/smarts"}' || fail "Failed request: $msg" ) | jq ".result"`
assert "$ret = false" "$msg"

echo "PASS"
