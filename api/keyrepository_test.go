package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Keys", func() {

	var getKeysResponse = `
	{
	  "count": 1,
	  "self": "/keys?page=1&per_page=30",
	  "first": "/keys?page=1&per_page=30",
	  "last": "/keys?page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	    {
	      "id": "86e03fc8-b639-4166-9a20-dbae948bdfc8",
	      "public": "ssh-rsa abe-23 xx@tw.com",
	      "fingerprint": "43:e8:e5:9b:bc:4c:c1:2e:60:ea:c8:cc:e0:b3:5a:d9",
	      "name": "id_rsa",
	      "created": "1451953908",
	      "owner": "ketsu@thoughtworks.com",
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys/86e03fc8-b639-4166-9a20-dbae948bdfc8"
	        },
	        {
	          "rel": "owner",
	          "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a"
	        }
	      ]
	    }
	  ]
	}
	`

	var getKeysRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/keys",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getKeysResponse,
		},
	}

	var createKeyRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo KeyRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewKeyRepository(configRepo, gateway)
		return
	}

	It("should able to get keys", func() {
		ts, _, repo := createKeyRepository([]testnet.TestRequest{getKeysRequest})
		defer ts.Close()

		keys, err := repo.GetKeys()
		Expect(err).To(BeNil())
		Expect(keys.Count()).To(Equal(1))
		Expect(keys.Items()[0].ID()).To(Equal("86e03fc8-b639-4166-9a20-dbae948bdfc8"))
		Expect(keys.Items()[0].Public()).To(Equal("ssh-rsa abe-23 xx@tw.com"))
		Expect(keys.Items()[0].Fingerprint()).To(Equal("43:e8:e5:9b:bc:4c:c1:2e:60:ea:c8:cc:e0:b3:5a:d9"))
		Expect(keys.Items()[0].Owner()).To(Equal("ketsu@thoughtworks.com"))
		Expect(keys.Items()[0].Links()).NotTo(BeNil())
	})
})
