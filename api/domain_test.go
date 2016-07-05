package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"github.com/sjkyspa/stacks/controller/api/fixtures"
	"net/http"
)

var _ = Describe("Domain", func() {

	var createDomainRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo DomainRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewDomainRepository(configRepo, gateway)
		return
	}

	It("should able to attach the domain cert", func() {
		DomainAttachCalled := false

		_, _, repo := createDomainRepository([]testnet.TestRequest{fixtures.DomainDetail(), fixtures.DomainAttachCert(func(r *http.Request) {
			DomainAttachCalled = true
		})})

		domain, err := repo.GetDomain("tw.com")
		err = domain.AttachCert(CertParams{Crt: "crt", Key: "key"})
		Expect(err).To(BeNil())
		Expect(DomainAttachCalled).To(BeTrue())
	})
})
