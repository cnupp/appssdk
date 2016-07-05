package fixtures

import (
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"fmt"
	"net/url"
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

func SuccessKetsuVerify(matcher func(r *http.Request)) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/verify/success",
		Matcher: matcher,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func FailKetsuVerify(matcher func(r *http.Request)) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/verify/fail",
		Matcher: matcher,
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
			  "envs": {
			    "ENV": "PRODUCTION"
			  },
			  "needDeploy": true,
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
		Path:   "/apps",
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

func Domains() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/domains",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 1,
			  "self": "/domains?page=1&per_page=30",
			  "first": "/domains?page=1&per_page=30",
			  "last": "/domains?page=1&per_page=30",
			  "prev": null,
			  "next": null,
			  "items": [
				{
				  "id": "b78dba518daf4fe99345c7ab582c3387",
				  "name": "tw.com",
				  "links": [
					{
					  "rel": "self",
					  "uri": "/domains/tw.com"
					}
				  ]
				}
			  ]
			}`,
		},
	}
}

func DomainDelete() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "DELETE",
		Path:   "/domains/tw.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: "",
		},
	}
}

func DomainDetail() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/domains/tw.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "id": "b78dba518daf4fe99345c7ab582c3387",
			  "name": "tw.com",
			  "links": [
				{
				  "rel": "self",
				  "uri": "/domains/tw.com"
				}
			  ]
			}`,
		},
	}
}

func DomainAttachCert(matcher func(r *http.Request)) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path:   "/domains/tw.com/cert",
		Matcher: matcher,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: ``,
		},
	}
}

func DomainCreate() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "POST",
		Path:   "/domains",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/domains/tw.com"},
			},
		},
	}
}

func Events(eventType string) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/events?type=%s", eventType),
		Response: testnet.TestResponse{
			Status: 200,
			Body:   Render(`
			{
			    "self": "/events?page=2&per-page=1&type={{.Type}}",
			    "first": "/events?page=1&per-page=1&type={{.Type}}",
			    "last": "/events?page=2&per-page=1&type={{.Type}}",
			    "prev": "/events?page=1&per-page=1&type={{.Type}}",
			    "next": "",
			    "count": 2,
			    "items": [
			      {
				"id": "2",
				"type": "{{.Type}}",
				"content": {
				  "createdAt": 1453274984000,
				  "release": {
				    "createdAt": 1453274984000,
				    "application": {
				      "name": "javajersey-api2",
				      "id": "060113d0767946f090d7a3a21b3008d2"
				    },
				    "envs": {},
				    "links": [
				      {
					"rel": "self",
					"uri": "/apps/javajersey-api2/releases/1453274984822"
				      },
				      {
					"rel": "app",
					"uri": "/apps/javajersey-api2"
				      },
				      {
					"rel": "build",
					"uri": "/apps/javajersey-api2/builds/86e03fc8b63941669a20dbae948bdfc8"
				      }
				    ],
				    "id": "1453274984822",
				    "version": 0
				  }
				}
			      }
			    ]
			    }
			`, map[string]string{"Type": eventType}),
		},
	}
}

func EventsOnPage(eventType string, total, page, perPage int) testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   fmt.Sprintf("/events?page=%d&per-page=%d&type=%s", page, perPage, eventType),
		Response: testnet.TestResponse{
			Status: 200,
			Body:   Render(`
			{
			    "self": "{{.Self}}",
			    "first": "{{.First}}",
			    "last": "{{.Last}}",
			    "prev": "{{.Prev}}",
			    "next": "{{.Next}}",
			    "count": {{.Total}},
			    "items": [
			      {
				"id": "2",
				"type": "{{.Type}}",
				"content": {
				  "createdAt": 1453274984000,
				  "release": {
				    "createdAt": 1453274984000,
				    "application": {
				      "name": "javajersey-api2",
				      "id": "060113d0767946f090d7a3a21b3008d2"
				    },
				    "envs": {},
				    "links": [
				      {
					"rel": "self",
					"uri": "/apps/javajersey-api2/releases/1453274984822"
				      },
				      {
					"rel": "app",
					"uri": "/apps/javajersey-api2"
				      },
				      {
					"rel": "build",
					"uri": "/apps/javajersey-api2/builds/86e03fc8b63941669a20dbae948bdfc8"
				      }
				    ],
				    "id": "1453274984822",
				    "version": 0
				  }
				}
			      }
			    ]
			    }
			`, map[string]string{
				"Type": eventType,
				"Self": PageGenerator(fmt.Sprintf("/events?type=%s", eventType), total, page, perPage).Self,
				"First": PageGenerator(fmt.Sprintf("/events?type=%s", eventType), total, page, perPage).First,
				"Last": PageGenerator(fmt.Sprintf("/events?type=%s", eventType), total, page, perPage).Last,
				"Prev": PageGenerator(fmt.Sprintf("/events?type=%s", eventType), total, page, perPage).Prev,
				"Next": PageGenerator(fmt.Sprintf("/events?type=%s", eventType), total, page, perPage).Next,
				"Total": fmt.Sprintf("%d", total),
			}),
		},
	}
}

func Keys() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "GET",
		Path:   "/keys",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: `
			{
			  "count": 1,
			  "self": "/keys?page=1&per_page=30",
			  "first": "/keys?page=1&per_page=30",
			  "last": "/keys?page=1&per_page=30",
			  "prev": null,
			  "next": null,
			  "items": [
			    {
			      "id": "86e03fc8-b639-4166-9a20-dbae948bdfc8",
			      "public": "ssh-rsa abe-23 xx@tw.com",
			      "fingerprint": "43:e8:e5:9b:bc:4c:c1:2e:60:ea:c8:cc:e0:b3:5a:d9",
			      "name": "id_rsa",
			      "created": "1451953908",
			      "owner": "ketsu@thoughtworks.com",
			      "links": [
				{
				  "rel": "self",
				  "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys/86e03fc8-b639-4166-9a20-dbae948bdfc8"
				},
				{
				  "rel": "owner",
				  "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a"
				}
			      ]
			    }
			  ]
			}`,
		},
	}
}

func KetsuDestroy() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "DELETE",
		Path:   "/apps/ketsu",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}
}

func KetsuTransferToUser() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path: "/apps/ketsu/transferred",
		Response: testnet.TestResponse{
			Status: 204,
		},
	}
}

func KetsuTransferToOrg() testnet.TestRequest {
	return testnet.TestRequest{
		Method: "PUT",
		Path: "/apps/ketsu/transferred",
		Response: testnet.TestResponse{
			Status: 204,
		},
	}}

type Page struct {
	First string
	Last  string
	Prev  string
	Next  string
	Self  string
}

func PageGenerator(prefix string, total, page, perPage int) Page {
	uri := func(page, perPage int) string {
		u, _ := url.Parse(prefix)
		query := u.Query()
		query.Set("type", fmt.Sprintf("%s", query.Get("type")))
		query.Add("page", fmt.Sprintf("%d", page))
		query.Add("per-page", fmt.Sprintf("%d", perPage))
		u.RawQuery = query.Encode()
		return u.String()
	}

	var last = func() string {
		last := (total / perPage) + 1
		return uri(last, perPage)
	}

	var first = func() string {
		return uri(1, perPage)
	}

	var self = func() string {
		return uri(page, perPage)
	}

	var next = func() string {
		last := (total / perPage) + 1
		next := page + 1
		if (next >= last) {
			return ""
		} else {
			return uri(next, perPage)
		}
	}

	var prev = func() string {
		prev := page - 1
		if (prev <= 1) {
			return ""
		} else {
			return uri(prev, perPage)
		}
	}

	return Page{
		Self: self(),
		First: first(),
		Last: last(),
		Prev: prev(),
		Next: next(),
	}
}



