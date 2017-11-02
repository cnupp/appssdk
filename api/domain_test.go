package api_test

import (
	. "github.com/cnupp/appssdk/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cnupp/appssdk/fixtures"
	"github.com/cnupp/appssdk/net"
	testconfig "github.com/cnupp/appssdk/testhelpers/config"
	testnet "github.com/cnupp/appssdk/testhelpers/net"
	"net/http"
	"net/http/httptest"
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
