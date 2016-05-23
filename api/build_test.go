package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	. "github.com/sjkyspa/stacks/controller/api/testhelpers/matchers"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Build", func() {
	var createBuildMapper = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo BuildMapper) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewBuildMapper(configRepo, gateway)
		return
	}

	var successBuildRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/success",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var failBuildRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/fail",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var successVerifyRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/verify/success",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var failVerifyRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/verify/fail",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var getAppBuildResponse = `
	{
	  "git_sha": "60bc43aa",
	  "created_at": 1456333105000,
	  "verify": {
		"id": "47de9390-03a6-4f27-a8f3-e2739c5c5e4a",
		"status": "NEW"
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
	  ],
	  "id": "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
	  "status": "NEW"
	}
	`

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

	var getAppBuildRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppBuildResponse,
		},
	}

	It("should able to success the build", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{getAppBuildRequest, successBuildRequest,getAppRequest})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")
		Expect(build.IsSuccess()).To(BeFalse())

		err := build.Success()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to fail the build", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{getAppBuildRequest, failBuildRequest, getAppRequest})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")

		err := build.Fail()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to set verify success", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{getAppBuildRequest, successBuildRequest, successVerifyRequest, getAppRequest})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")
		Expect(build.IsVerifySuccess()).To(BeFalse())
		err := build.Success()
		Expect(err).To(BeNil())
		err = build.VerifySuccess()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to set verify fail", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{getAppBuildRequest, successBuildRequest, failVerifyRequest, getAppRequest})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")
		Expect(build.IsVerifySuccess()).To(BeFalse())
		err := build.Success()
		Expect(err).To(BeNil())
		err = build.VerifyFail()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})
})
