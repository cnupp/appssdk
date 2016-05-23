package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http/httptest"
	"github.com/sjkyspa/stacks/controller/api/fixtures"
)

var _ = Describe("Domains", func() {
	var createDomainRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo DomainRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewDomainRepository(configRepo, gateway)
		return
	}

	var defaultDomainParams = func() DomainParams {
		name := "tw.com"

		return DomainParams{
			Name: name,
		}
	}

	It("should able to create an domain", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{fixtures.DomainCreate(), fixtures.DomainDetail()})
		defer ts.Close()

		createdDomain, err := repo.Create(defaultDomainParams())
		Expect(err).To(BeNil())
		Expect(createdDomain.Id()).To(Equal("b78dba518daf4fe99345c7ab582c3387"))
		Expect(createdDomain.Name()).To(Equal("tw.com"))
	})

	It("should able to get an domain", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{fixtures.DomainDetail()})
		defer ts.Close()

		createdDomain, err := repo.GetDomain("tw.com")
		Expect(err).To(BeNil())
		Expect(createdDomain.Id()).To(Equal("b78dba518daf4fe99345c7ab582c3387"))
		Expect(createdDomain.Name()).To(Equal("tw.com"))
	})

	It("should able to get domains", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{fixtures.Domains()})
		defer ts.Close()

		createdDomains, err := repo.GetDomains()
		Expect(err).To(BeNil())
		Expect(createdDomains.Count()).To(Equal(1))
		Expect(createdDomains.Items()[0].Id()).To(Equal("b78dba518daf4fe99345c7ab582c3387"))
		Expect(createdDomains.Items()[0].Name()).To(Equal("tw.com"))
		Expect(createdDomains.Items()[0].Links()).NotTo(BeNil())
	})

	It("should remove a domain", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{fixtures.DomainDelete()})
		defer ts.Close()

		err := repo.Delete("tw.com")
		Expect(err).To(BeNil())
	})
})
