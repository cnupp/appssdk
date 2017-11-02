package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cnupp/appssdk/net"
	testconfig "github.com/cnupp/appssdk/testhelpers/config"
	testnet "github.com/cnupp/appssdk/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Routes", func() {

	var createRouteRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/routes",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/routes/8399de76-eeef-418d-a567-75253b03c4ec"},
			},
		},
	}

	var getRoutesResponse = `
	{
	  "count": 2,
	  "self": "",
	  "first": "",
	  "last": "",
	  "prev": "",
	  "next": "",
	  "items": [
	    {
	      "id": "8399de76-eeef-418d-a567-75253b03c4ec",
	      "path": "/path",
	      "domain": {
	        "name": "deepi.cn"
	      },
	      "created": "1451953908",
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/routes/8399de76-eeef-418d-a567-75253b03c4ec"
	        },
	        {
	          "rel": "domain",
	          "uri": "/domains/5576178a-2e0e-4ed7-842d-26de65a6992f"
	        },
	        {
	          "rel": "apps",
	          "uri": "/routes/8399de76-eeef-418d-a567-75253b03c4ec/apps"
	        }
	      ]
	    }
	  ]
	}
	`
	var getRoutesRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/routes",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getRoutesResponse,
		},
	}

	var getAppsUnderRouteResponse = `
	{
	  "count": 1,
	  "self": "/routes/8399de76-eeef-418d-a567-75253b03c4ec/apps?page=1&per_page=30",
	  "first": "/routes/8399de76-eeef-418d-a567-75253b03c4ec/apps?page=1&per_page=30",
	  "last": "/routes/8399de76-eeef-418d-a567-75253b03c4ec/apps?page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	    {
	      "name": "ketsu",
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
	        },
	        {
	          "rel": "routes",
	          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/routes"
	        },
	        {
	          "rel": "builds",
	          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds"
	        },
	        {
	          "rel": "releases",
	          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/releases"
	        },
	        {
	          "rel": "env",
	          "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/env"
	        }
	      ]
	    }
	  ]
	}

	`

	var getAppsUnderRouteRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/routes/8399de76-eeef-418d-a567-75253b03c4ec/apps",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppsUnderRouteResponse,
		},
	}

	var createRouteRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo RouteRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewRouteRepository(configRepo, gateway)
		return
	}

	var defaultRouteParams = func() RouteParams {
		return RouteParams{
			Domain: "test.tw.com",
			Path:   "/path",
		}
	}

	It("should able to define an route", func() {
		ts, _, repo := createRouteRepository([]testnet.TestRequest{createRouteRequest})
		defer ts.Close()

		err := repo.Create(defaultRouteParams())
		Expect(err).To(BeNil())
	})

	It("should able to list routes", func() {
		ts, _, repo := createRouteRepository([]testnet.TestRequest{getRoutesRequest})
		defer ts.Close()

		routes, err := repo.GetRoutes()
		Expect(err).To(BeNil())
		Expect(routes.Count()).To(Equal(2))
		Expect(routes.Items()[0].ID()).To(Equal("8399de76-eeef-418d-a567-75253b03c4ec"))
		Expect(routes.Items()[0].Path()).To(Equal("/path"))
		Expect(routes.Items()[0].Domain().Name).To(Equal("deepi.cn"))
		Expect(routes.Items()[0].Links()).NotTo(BeNil())
	})

	It("should able to list apps under a route", func() {
		ts, _, repo := createRouteRepository([]testnet.TestRequest{getAppsUnderRouteRequest})
		defer ts.Close()

		apps, err := repo.GetAppsForRoute("8399de76-eeef-418d-a567-75253b03c4ec")
		Expect(err).To(BeNil())
		Expect(apps.Count()).To(Equal(1))
		Expect(apps.Items()[0].Id()).To(Equal("ketsu"))
		Expect(apps.Items()[0].Links()).NotTo(BeNil())
	})
})
