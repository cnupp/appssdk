package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	"net/http"
)

var _ = Describe("Resource", func() {
	Context("Authorized User", func() {
		var buildRequest = testnet.TestRequest{
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
					}
					`,
			},
		}

		var putBuildRequest = testnet.TestRequest{
			Method: "PUT",
			Path: "/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8/success",
			Response: testnet.TestResponse{
				Status: 200,
				Header: http.Header{
					"Content-Type": {"application/json"},
				},
			},
		}

		var getAppResponse = `
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
		`

		var getAppRequest = testnet.TestRequest{
			Method: "GET",
			Path:   "/apps/ketsu",
			Response: testnet.TestResponse{
				Status: 200,
				Header: http.Header{
					"Content-Type": {"application/json"},
				},
				Body: getAppResponse,
			},
		}

		It("should able to get the build by the uri", func(done Done) {
			ts, _ := testnet.NewServer([]testnet.TestRequest{
				buildRequest, getAppRequest, putBuildRequest,
			})
			defer ts.Close()

			configRepo := testconfig.NewConfigRepository()
			configRepo.SetApiEndpoint(ts.URL)
			configRepo.SetAuth("auth")
			gateway := net.NewCloudControllerGateway(configRepo)

			resource, err := NewResource(configRepo, gateway).GetResourceByURI("/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8")
			Expect(err).To(BeNil())
			Expect(resource).NotTo(BeNil())
			Expect(resource.(Build).GitSha()).To(Equal("60bc43aa"))
			Expect(resource.(Build).Success()).To(BeNil())
			close(done)
		}, 60)
	})

})
