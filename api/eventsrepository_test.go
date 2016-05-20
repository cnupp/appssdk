package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	. "github.com/sjkyspa/stacks/controller/api/testhelpers/matchers"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http/httptest"
)

var _ = Describe("Eventsrepository", func() {
	var createEventRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo EventRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewEventRepository(configRepo, gateway)
		return
	}

	var rootEventsLatestPage = `
	{
  	    "next": "",
  	    "last": "/events?type=ReleaseSuccessEvent&page=2&per-page=1",
  	    "prev": "/events?type=ReleaseSuccessEvent&page=1&per-page=1",
  	    "self": "/events?type=ReleaseSuccessEvent&page=2&per-page=1",
  	    "count": 2,
  	    "items": [
  	      {
  	        "id": "2",
  	        "type": "ReleaseSuccessEvent",
  	        "content": {
  	          "createdAt": 1453274984000,
  	          "release": {
  	            "createdAt": 1453274984000,
  	            "application": {
  	              "name": "javajersey-api2",
  	              "id": "060113d0-7679-46f0-90d7-a3a21b3008d2"
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
  	                "uri": "/apps/javajersey-api2/builds/3ff77dab-9d97-4a53-a13c-1b61958b3d7b"
  	              }
  	            ],
  	            "id": "1453274984822",
  	            "version": 0
  	          }
  	        }
  	      }
  	    ],
  	    "first": "/events?type=ReleaseSuccessEvent&page=1&per-page=1"
	    }
	`

	var rootEventsResponseFirstPage = `
	{
  	    "next": "/events?type=ReleaseSuccessEvent&page=2&per-page=1",
  	    "last": "/events?type=ReleaseSuccessEvent&page=2&per-page=1",
  	    "prev": "",
  	    "self": "/events?type=ReleaseSuccessEvent&page=1&per-page=1",
  	    "count": 2,
  	    "items": [
  	      {
  	        "id": "1",
  	        "type": "ReleaseSuccessEvent",
  	        "content": {
  	          "createdAt": 1453274984000,
  	          "release": {
  	            "createdAt": 1453274984000,
  	            "application": {
  	              "name": "javajersey-api2",
  	              "id": "060113d0-7679-46f0-90d7-a3a21b3008d2"
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
  	                "uri": "/apps/javajersey-api2/builds/3ff77dab-9d97-4a53-a13c-1b61958b3d7b"
  	              }
  	            ],
  	            "id": "1453274984822",
  	            "version": 0
  	          }
  	        }
  	      }
  	    ],
  	    "first": "/events?type=ReleaseSuccessEvent&page=1&per-page=1"
	}
	`
	var latestRootEvents = testnet.TestRequest{
		Method: "GET",
		Path:   "/events?type=ReleaseSuccessEvent",
		Response: testnet.TestResponse{
			Status: 200,
			Body:   rootEventsLatestPage,
		},
	}

	var secondRootEvents = testnet.TestRequest{
		Method: "GET",
		Path:   "/events?type=ReleaseSuccessEvent&page=2&per-page=1",
		Response: testnet.TestResponse{
			Status: 200,
			Body:   rootEventsLatestPage,
		},
	}

	var firstRootEvents = testnet.TestRequest{
		Method: "GET",
		Path:   "/events?type=ReleaseSuccessEvent&page=1&per-page=1",
		Response: testnet.TestResponse{
			Status: 200,
			Body:   rootEventsResponseFirstPage,
		},
	}

	It("should able to get latest the events", func() {
		ts, handler, er := createEventRepository([]testnet.TestRequest{
			latestRootEvents,
		})
		defer ts.Close()

		events, err := er.GetEvents("ReleaseSuccessEvent")
		Expect(err).To(BeNil())
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(events.Count()).To(Equal(2))
		Expect(len(events.Items())).To(Equal(1))
		Expect(events.Items()[0].Type()).To(Equal("ReleaseSuccessEvent"))
		Expect(events.Items()[0].Links()).NotTo(BeNil())
	})

	It("should able to get prev page the events", func() {
		ts, handler, er := createEventRepository([]testnet.TestRequest{
			latestRootEvents,
			firstRootEvents,
		})
		defer ts.Close()

		events, _ := er.GetEvents("ReleaseSuccessEvent")
		prev, _ := events.Prev()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(len(prev.Items())).To(Equal(1))
	})

	It("should able to get next events", func() {
		ts, handler, er := createEventRepository([]testnet.TestRequest{
			latestRootEvents,
			firstRootEvents,
			secondRootEvents,
		})
		defer ts.Close()

		events, _ := er.GetEvents("ReleaseSuccessEvent")
		prev, _ := events.Prev()
		next, _ := prev.Next()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(len(prev.Items())).To(Equal(1))
		Expect(len(next.Items())).To(Equal(1))
	})
})
