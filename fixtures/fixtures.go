package fixtures

import (
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
)

func KetsuBuild() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path: "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
				"git_sha": "60bc43aa",
				"created_at": 1456333105000,
				"verify": {
					"id": "66e03fc8b63941669a20dbae948bdfc8",
					"status": "NEW"
				},
				"links": [
					{
						"rel": "self",
						"uri": "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8"
					},
					{
						"rel": "app",
						"uri": "/apps/ketsu"
					}
				],
				"id": "86e03fc8b63941669a20dbae948bdfc8",
				"status": "NEW"
			}`,
		},
	}
}

func KetsuBuilds() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path: "/apps/ketsu/builds",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 1,
			  "self": "/apps/ketsu/builds?page=1&per_page=30",
			  "first": "/apps/ketsu/builds?page=1&per_page=30",
			  "last": "/apps/ketsu/builds?page=1&per_page=30",
			  "prev": null,
			  "next": null,
			  "items": [
			    {
			      "created": "1451953908",
			      "git_sha": "60bc43aa",
			      "id": "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
			      "status": "NEW",
			      "app": {
				"name": "ketsu"
			      },
			      "links": [
				{
				  "rel": "self",
				  "uri": "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
				},
				{
				  "rel": "app",
				  "uri": "/apps/ketsu"
				}
			      ]
			    }
			  ]
			}`,
		},
	}
}

func SuccessKetsuBuild() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path: "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/success",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func KetsuDetail() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
		{
		  "id": "b78dba518daf4fe99345c7ab582c3387",
		  "name": "ketsu",
		  "memory": 30,
		  "disk": 30,
		  "instances": 1,
		  "links": [
			{
			  "rel": "self",
			  "uri": "/apps/ketsu"
			},
			{
			  "rel": "env",
			  "uri": "/apps/ketsu/env"
			},
			{
			  "rel": "routes",
			  "uri": "/apps/ketsu/routes"
			},
			{
			  "rel": "builds",
			  "uri": "/apps/ketsu/builds"
			},
			{
			  "rel": "releases",
			  "uri": "/apps/ketsu/releases"
			},
			{
			  "rel": "stack",
			  "uri": "/stacks/javajersey"
			}
		  ]
		}
		`,
		},
	}
}

func AppList() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps?page=1&per-page=30",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 2,
			  "self": "/apps?page=1&per-page=30",
			  "first": "/apps?page=1&per-page=30",
			  "last": "/apps?page=1&per-page=30",
			  "prev": null,
			  "next": null,
			  "items": [
				{
				  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
				  "name": "ketsu",
				  "links": [
					{
					  "rel": "self",
					  "uri": "/apps/ketsu"
					},
					{
					  "rel": "env",
					  "uri": "/apps/ketsu/env"
					},
					{
					  "rel": "routes",
					  "uri": "/apps/ketsu/routes"
					},
					{
					  "rel": "builds",
					  "uri": "/apps/ketsu/builds"
					},
					{
					  "rel": "releases",
					  "uri": "/apps/ketsu/releases"
					},
					{
					  "rel": "stack",
					  "uri": "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6"
					}
				  ]
				}
			  ]
			}`,
		},
	}
}

