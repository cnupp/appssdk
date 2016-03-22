package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Domains", func() {
	var createDomainRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/domains",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/domains/tw.com"},
			},
		},
	}
	var getDomainResponse = `
	{
	  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
	  "name": "tw.com",
	  "links": [
		{
		  "rel": "self",
		  "uri": "/domains/tw.com"
		}
	  ]
	}
	`

	var getDomainRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/domains/tw.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getDomainResponse,
		},
	}

	var delteDomainRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/domains/tw.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: "",
		},
	}

	var getDomainsResponse = `
	{
	  "count": 1,
	  "self": "/domains?page=1&per_page=30",
	  "first": "/domains?page=1&per_page=30",
	  "last": "/domains?page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
		{
		  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
		  "name": "tw.com",
		  "links": [
			{
			  "rel": "self",
			  "uri": "/domains/tw.com"
			}
		  ]
		}
	  ]
	}
	`

	var getDomainsRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/domains",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getDomainsResponse,
		},
	}

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
		ts, _, repo := createDomainRepository([]testnet.TestRequest{createDomainRequest, getDomainRequest})
		defer ts.Close()

		createdDomain, err := repo.Create(defaultDomainParams())
		Expect(err).To(BeNil())
		Expect(createdDomain.Id()).To(Equal("b78dba51-8daf-4fe9-9345-c7ab582c3387"))
		Expect(createdDomain.Name()).To(Equal("tw.com"))
	})

	It("should able to get an domain", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{getDomainRequest})
		defer ts.Close()

		createdDomain, err := repo.GetDomain("tw.com")
		Expect(err).To(BeNil())
		Expect(createdDomain.Id()).To(Equal("b78dba51-8daf-4fe9-9345-c7ab582c3387"))
		Expect(createdDomain.Name()).To(Equal("tw.com"))
	})

	It("should able to get domains", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{getDomainsRequest})
		defer ts.Close()

		createdDomains, err := repo.GetDomains()
		Expect(err).To(BeNil())
		Expect(createdDomains.Count()).To(Equal(1))
		Expect(createdDomains.Items()[0].Id()).To(Equal("b78dba51-8daf-4fe9-9345-c7ab582c3387"))
		Expect(createdDomains.Items()[0].Name()).To(Equal("tw.com"))
		Expect(createdDomains.Items()[0].Links()).NotTo(BeNil())
	})

	It("should remove a domain", func() {
		ts, _, repo := createDomainRepository([]testnet.TestRequest{delteDomainRequest})
		defer ts.Close()

		err := repo.Delete("tw.com")
		Expect(err).To(BeNil())
	})
})
