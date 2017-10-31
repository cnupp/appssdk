package fixtures

import (
	testnet "github.com/cnupp/cnup/controller/api/testhelpers/net"
	"net/http"
)

func Login() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/auths",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/auths/47631d4225d14fdea8b502d94f0d616d"},
			},
		},
	}
}

func InvalidLogin() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/auths",
		Response: testnet.TestResponse{
			Status: 400,
			Header: http.Header{
				"accept": {"application/json"},
			},
			Body: "error",
		},
	}
}

func Logout() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "DELETE",
		Path:   "/auths/47631d4225d14fdea8b502d94f0d616d",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}
}

func Auths() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/auths",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "id": "47631d42-25d1-4fde-a8b5-02d94f0d616d",
			  "email": "ketsu@thoughtworks.com",
			  "links": [
				{
				  "rel": "self",
				  "uri": "/users/47631d4225d14fdea8b502d94f0d616d"
				},
				{
				  "rel": "keys",
				  "uri": "/users/47631d42-25d14fdea8b502d94f0d616d/keys"
				}
			  ]
			}`,
		},
	}
}
