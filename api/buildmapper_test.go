package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"
	"github.com/sjkyspa/stacks/controller/api/net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Builds", func() {
	It("should able to create an build for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{createAppBuildRequest, getAppBuildRequest})
		defer ts.Close()

		build, err := mapper.Create(AppModel{
			ID: "ketsu",
		}, BuildParams{
			GitSha: "60bc43aa",
			User:   "user",
		})

		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
		Expect(build.Verify().StatusField).To(Equal("NEW"))
	})

	It("should able to get all builds for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{getAppBuildsRequest})
		defer ts.Close()

		builds, err := mapper.GetBuilds(AppModel{
			ID: "ketsu",
		})

		Expect(err).To(BeNil())
		Expect(builds.Count()).To(Equal(1))
		Expect(len(builds.Items())).To(Equal(1))
		Expect(builds.Items()[0].GitSha).To(Equal("60bc43aa"))
		Expect(builds.Items()[0].ID).To(Equal("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"))
		Expect(builds.Items()[0].Links).To(Not(BeNil()))
		Expect(builds.Items()[0].Status).To(Equal("NEW"))
	})

	It("should able to get one build for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{getAppBuildRequest})
		defer ts.Close()

		build, err := mapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")

		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
		Expect(build.Id()).To(Equal("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"))
		Expect(build.Links()).To(Not(BeNil()))
		Expect(build.Status()).To(Equal("NEW"))
	})
})

var createAppBuildRequest = testnet.TestRequest{
	Method: "POST",
	Path:   "/apps/ketsu/builds",
	Response: testnet.TestResponse{
		Status: 200,
		Header: http.Header{
			"Location": {"/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"},
		},
	},
}

var getAppBuildsResponse = `
{
  "count": 1,
  "self": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds?page=1&per_page=30",
  "first": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds?page=1&per_page=30",
  "last": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds?page=1&per_page=30",
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
          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
        },
        {
          "rel": "app",
          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
        }
      ]
    }
  ]
}
`

var getAppBuildsRequest = testnet.TestRequest{
	Method: "GET",
	Path:   "/apps/ketsu/builds",
	Response: testnet.TestResponse{
		Status: 200,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: getAppBuildsResponse,
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

func createBuildMapper(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo BuildMapper) {
	ts, handler = testnet.NewServer(requests)
	configRepo := testconfig.NewRepositoryWithDefaults()
	configRepo.SetApiEndpoint(ts.URL)
	gateway := net.NewCloudControllerGateway(configRepo)
	repo = NewBuildMapper(configRepo, gateway)
	return
}
