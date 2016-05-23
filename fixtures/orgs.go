package fixtures

import (
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
)

func CollaboratorsRemove() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "DELETE",
		Path:   "/apps/ketsu/collaborators/abc",
		Response: testnet.TestResponse{
			Status: 204,
		},
	}
}
func CollaboratorsAdd() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/collaborators",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu/collaborators/123"},
			},
		},
	}
}

func KetsuCollaborators() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/collaborators",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			[
				{
				  "id": "47631d4225d14fdea8b502d94f0d616d",
				  "email": "ketsu@thoughtworks.com",
				  "links": [
					{
					  "rel": "self",
					  "uri": "/users/47631d4225d14fdea8b502d94f0d616d"
					}
				  ]
				}
			]`,
		},
	}
}

func EmptyCollaborators() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/empty/collaborators",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `[]`,
		},
	}
}



