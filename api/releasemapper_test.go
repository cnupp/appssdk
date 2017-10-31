package api_test

import (
	. "github.com/cnupp/cnup/controller/api/api"
	"github.com/cnupp/cnup/controller/api/net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	testconfig "github.com/cnupp/cnup/controller/api/testhelpers/config"
	testnet "github.com/cnupp/cnup/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("releases", func() {
	It("should able to create an release for app", func() {
		ts, _, mapper := createReleaseMapper([]testnet.TestRequest{createAppReleaseRequest, getAppReleaseRequest})
		defer ts.Close()
		_, err := mapper.Create(AppModel{
			ID: "ketsu",
		})

		Expect(err).To(BeNil())
	})

	It("should able to get all releases for app", func() {
		ts, _, mapper := createReleaseMapper([]testnet.TestRequest{getAppReleasesRequest})
		defer ts.Close()

		releases, err := mapper.GetReleases(AppModel{
			ID: "ketsu",
		})

		Expect(err).To(BeNil())
		Expect(releases.Count()).To(Equal(1))
		Expect(len(releases.Items())).To(Equal(1))
		Expect(releases.Items()[0].Id()).To(Equal("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"))
		Expect(releases.Items()[0].Envs()).To(Not(BeNil()))
		Expect(releases.Items()[0].Version()).To(Equal(1))
		Expect(releases.Items()[0].Links()).To(Not(BeNil()))
		Expect(releases.Items()[0].Status()).To(Equal("NEW"))
	})

	It("should able to get one release for app", func() {
		ts, _, mapper := createReleaseMapper([]testnet.TestRequest{getAppReleaseRequest})
		defer ts.Close()

		release, err := mapper.GetRelease(AppModel{
			ID: "ketsu",
		}, "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")

		Expect(err).To(BeNil())
		Expect(release.Id()).To(Equal("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"))
		Expect(release.Links()).To(Not(BeNil()))
		Expect(release.Status()).To(Equal("NEW"))
	})
})

var createAppReleaseRequest = testnet.TestRequest{
	Method: "POST",
	Path:   "/apps/ketsu/releases",
	Response: testnet.TestResponse{
		Status: 200,
		Header: http.Header{
			"Location": {"/apps/ketsu/releases/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"},
		},
	},
}

var getAppReleasesResponse = `
{
  "count": 1,
  "self": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases?page=1&per_page=30",
  "first": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases?page=1&per_page=30",
  "last": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases?page=1&per_page=30",
  "prev": null,
  "next": null,
  "items": [
    {
      "created": "1451953908",
      "id": "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
      "version": 1,
      "status": "NEW",
      "application": {
        "name": "ketsu"
      },
      "envs": {
      	"mysql": "www.tw.com:3019"
      },
      "links": [
        {
          "rel": "self",
          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
        },
        {
          "rel": "app",
          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
        },
        {
          "rel": "build",
          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
        }
      ]
    }
  ]
}
`

var getAppReleasesRequest = testnet.TestRequest{
	Method: "GET",
	Path:   "/apps/ketsu/releases",
	Response: testnet.TestResponse{
		Status: 200,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: getAppReleasesResponse,
	},
}

var getAppReleaseResponse = `
{
  "created": "1451953908",
  "id": "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
  "version": 1,
  "status": "NEW",
  "application": {
	"name": "ketsu"
  },
  "envs": {
   	"MYSQL": "WWW.TW.COM:3019"
  },
  "links": [
	{
	  "rel": "self",
	  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
	},
	{
	  "rel": "app",
	  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
	},
	{
	  "rel": "verifies",
	  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/verifies"
	}
  ]
}
`

var getAppReleaseRequest = testnet.TestRequest{
	Method: "GET",
	Path:   "/apps/ketsu/releases/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
	Response: testnet.TestResponse{
		Status: 200,
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
		Body: getAppReleaseResponse,
	},
}

func createReleaseMapper(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo ReleaseMapper) {
	ts, handler = testnet.NewServer(requests)
	configRepo := testconfig.NewRepositoryWithDefaults()
	configRepo.SetApiEndpoint(ts.URL)
	gateway := net.NewCloudControllerGateway(configRepo)
	repo = NewReleaseMapper(configRepo, gateway)
	return
}
