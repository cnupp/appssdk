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

var _ = Describe("User", func() {

	var getUserResponse = `
	{
	  "id": "46208f69-0082-4db0-ba08-bfa39ccfdc2a",
	  "email": "ketsu@thoughtworks.com",
	  "links": [
		{
		  "rel": "self",
		  "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a"
		},
		{
		  "rel": "keys",
		  "uri": "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys"
		}
	  ]
	}
	`

	var getUserRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getUserResponse,
		},
	}

	var uploadKeyRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys/86e03fc8-b639-4166-9a20-dbae948bdfc8"},
			},
		},
	}

	var getKeyResponse = `
	{
	  "id": "86e03fc8-b639-4166-9a20-dbae948bdfc8",
	  "public": "ssh-rsa abe-23 xx@tw.com",
	  "fingerprint": "43:e8:e5:9b:bc:4c:c1:2e:60:ea:c8:cc:e0:b3:5a:d9",
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
	`

	var getKeysResponse = `
	{
	  "count": 1,
	  "self": "users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys?page=1&per_page=30",
	  "first": "users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys?page=1&per_page=30",
	  "last": "users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys?page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	    {
	      "id": "86e03fc8-b639-4166-9a20-dbae948bdfc8",
	      "public": "ssh-rsa abe-23 xx@tw.com",
	      "fingerprint": "43:e8:e5:9b:bc:4c:c1:2e:60:ea:c8:cc:e0:b3:5a:d9",
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

	var getKeyRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys/86e03fc8-b639-4166-9a20-dbae948bdfc8",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getKeyResponse,
		},
	}

	var defaultKeyParams = func() KeyParams {
		return KeyParams{
			Public: "ssh-rsa abe-23 xx@tw.com",
		}
	}

	var createUserRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo UserRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewUserRepository(configRepo, gateway)
		return
	}

	var deleteKeyRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys/86e03fc8-b639-4166-9a20-dbae948bdfc8",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}

	var getKeysRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users/46208f69-0082-4db0-ba08-bfa39ccfdc2a/keys",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getKeysResponse,
		},
	}

	It("should able to upload one key for user", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserRequest, uploadKeyRequest, getKeyRequest})
		defer ts.Close()

		user, err := repo.GetUser("46208f69-0082-4db0-ba08-bfa39ccfdc2a")
		_, err = user.UploadKey(defaultKeyParams())

		Expect(err).To(BeNil())
	})

	It("should able to delete key", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserRequest, deleteKeyRequest})
		defer ts.Close()

		user, err := repo.GetUser("46208f69-0082-4db0-ba08-bfa39ccfdc2a")
		err = user.DeleteKey("86e03fc8-b639-4166-9a20-dbae948bdfc8")

		Expect(err).To(BeNil())
	})

	It("should able to list keys", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserRequest, getKeysRequest})
		defer ts.Close()

		user, err := repo.GetUser("46208f69-0082-4db0-ba08-bfa39ccfdc2a")
		keys, err := user.Keys()

		Expect(err).To(BeNil())
		Expect(keys.Count()).To(Equal(1))
		Expect(keys.Items()[0].Public()).NotTo(BeNil())
		Expect(keys.Items()[0].Fingerprint()).NotTo(BeNil())
	})
})
