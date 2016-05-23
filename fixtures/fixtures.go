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

func KetsuBuildCreate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/builds",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Location": {"/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8"},
			},
		},
	}
}

func KetsuCreate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/apps",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu"},
			},
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
			      "id": "86e03fc8b63941669a20dbae948bdfc8",
			      "status": "NEW",
			      "app": {
				"name": "ketsu"
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
			      ]
			    }
			  ]
			}`,
		},
	}
}

func SuccessKetsuBuild(matcher func(r *http.Request)) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path: "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/success",
		Matcher: matcher,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func FailKetsuBuild(matcher func(r *http.Request)) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path: "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/fail",
		Matcher: matcher,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func SuccessKetsuVerify() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/verify/success",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func FailKetsuVerify() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/verify/fail",
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

func KetsuStackDetail() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks/javajersey",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `{
				"id": "74a052c976b344a1ac0b666faa1223b6",
				"name": "javajersey",
				"links": [
				  {
				    "rel": "self",
				    "uri": "/stacks/javajersey"
				  }
				]
			}`,
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
				  "id": "b78dba518daf4fe99345c7ab582c3387",
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
					  "uri": "/stacks/74a052c976b344a1ac0b666faa1223b6"
					}
				  ]
				}
			  ]
			}`,
		},
	}
}

func KetsuRoutes() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/routes",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 31,
			  "self": "/apps/ketsu/routes?page=1&per_page=30",
			  "first": "/apps/ketsu/routes?page=1&per_page=30",
			  "last": "/apps/ketsu/routes?page=2&per_page=30",
			  "prev": "",
			  "next": "/apps/ketsu/routes?page=2&per_page=30",
			  "items": [
			    {
			      "id": "8399de76eeef418da56775253b03c4ec",
			      "path": "/path",
			      "domain": {
				"name": "deepi.cn"
			      },
			      "app": {
				"name": "ketsu"
			      },
			      "links": [
				{
				  "rel": "self",
				  "uri": "/apps/ketsu/routes/8399de76eeef418da56775253b03c4ec"
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

func KetsuRoutesSecondPage() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/routes?page=2&per_page=30",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 31,
			  "self": "/apps/ketsu/routes?page=2&per_page=30",
			  "first": "/apps/ketsu/routes?page=1&per_page=30",
			  "last": "/apps/ketsu/routes?page=2&per_page=30",
			  "prev": "/apps/ketsu/routes?page=1&per_page=30",
			  "next": "",
			  "items": [
			    {
			      "id": "f18045634c314aae9e2507e4d9088d2c",
			      "path": "/path",
			      "domain": {
				"name": "deepi.cn"
			      },
			      "app": {
				"name": "ketsu"
			      },
			      "links": [
				{
				  "rel": "self",
				  "uri": "/apps/ketsu/routes/f18045634c314aae9e2507e4d9088d2c"
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

func KetsuRoutesBind() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/routes",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu/routes/f18045634c314aae9e2507e4d9088d2c"},
			},
		},
	}
}

func KetsuRoutesUnbind() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "DELETE",
		Path:   "/apps/ketsu/routes/test.tw.com/path",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func KetsuStackUpdate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/switch-stack",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func KetsuEnvCreate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/env",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}
}

func KetsuEnvUpdate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/env",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}
}

func KetsuBuildLog() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/log?lines=15&log_type=build&offset=0",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "total": 2,
			  "size": 2,
			  "items": [
			    {
			      "message": "init successful"
			    },
			    {
			      "message": "success"
			    }
			  ]
			}`,
		},
	}
}


func KaylaPermissionOnKetsu() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path: "/apps/ketsu/permissions?user=abcd",
		Response: testnet.TestResponse{
			Status: 200,
			Body: `
			{
			  "write": true,
			  "read": false
			}`,
		},
	}
}