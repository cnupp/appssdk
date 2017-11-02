package api_test

import (
	. "github.com/cnupp/appssdk/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cnupp/appssdk/fixtures"
	"github.com/cnupp/appssdk/net"
	testconfig "github.com/cnupp/appssdk/testhelpers/config"
	. "github.com/cnupp/appssdk/testhelpers/matchers"
	testnet "github.com/cnupp/appssdk/testhelpers/net"
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

	var latestRootEvents = fixtures.Events("ReleaseSuccessEvent")
	var rootEventsSecondPage = fixtures.EventsOnPage("ReleaseSuccessEvent", 2, 2, 1)
	var rootEventsFirstPage = fixtures.EventsOnPage("ReleaseSuccessEvent", 2, 1, 1)

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
			rootEventsFirstPage,
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
			rootEventsFirstPage,
			rootEventsSecondPage,
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
