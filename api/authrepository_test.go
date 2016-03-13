package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Auths", func() {
	var createAuthRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/auths",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/auths/47631d42-25d1-4fde-a8b5-02d94f0d616d"},
			},
		},
	}

	var createFailedAuthRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/auths",
		Response: testnet.TestResponse{
			Status: 400,
			Header: http.Header{
				"accept": {"application/json"},
			},
			Body: "error",
		},
	}

	var deleteAuthRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/auths/47631d42-25d1-4fde-a8b5-02d94f0d616d",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}

	var authId = "47631d42-25d1-4fde-a8b5-02d94f0d616d"
	var userEmail = "ketsu@thoughtworks.com"
	var userPassword = "123456"

	var createAuthRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo AuthRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewAuthRepository(configRepo, gateway)
		return
	}

	var defaultAuthParams = func() UserParams {
		return UserParams{
			Email:    userEmail,
			Password: userPassword,
		}
	}

	It("should able to create an authorization with correct user and password", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{createAuthRequest})
		defer ts.Close()

		createdAuth, err := repo.Create(defaultAuthParams())
		Expect(err).To(BeNil())
		Expect(createdAuth.UserEmail()).To(Equal(userEmail))
		Expect(createdAuth.Id()).To(Equal(authId))
	})

	It("should show error with wrong user password", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{createFailedAuthRequest})
		defer ts.Close()

		createdAuth, err := repo.Create(defaultAuthParams())
		Expect(createdAuth).To(BeNil())
		Expect(err).ToNot(BeNil())
	})

	It("should able to delete an authorization", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{deleteAuthRequest})
		defer ts.Close()

		err := repo.Delete(authId)
		Expect(err).To(BeNil())
	})
})
