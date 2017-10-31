package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/cnupp/cnup/controller/api/api"
	"github.com/cnupp/cnup/controller/api/net"
	testconfig "github.com/cnupp/cnup/controller/api/testhelpers/config"
	testnet "github.com/cnupp/cnup/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Users", func() {
	var createUserRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/users",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/users/47631d42-25d1-4fde-a8b5-02d94f0d616d"},
			},
		},
	}
	var getUserResponse = `
	{
	  "id": "47631d42-25d1-4fde-a8b5-02d94f0d616d",
	  "email": "ketsu@thoughtworks.com",
	  "links": [
		{
		  "rel": "self",
		  "uri": "/users/47631d42-25d1-4fde-a8b5-02d94f0d616d"
		},
		{
		  "rel": "keys",
		  "uri": "/users/47631d42-25d1-4fde-a8b5-02d94f0d616d/keys"
		}
	  ]
	}
	`

	var getUserRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users/47631d42-25d1-4fde-a8b5-02d94f0d616d",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getUserResponse,
		},
	}

	var userId = "47631d42-25d1-4fde-a8b5-02d94f0d616d"
	var userEmail = "ketsu@thoughtworks.com"

	var getUserByEmailResponse = `
	{
	  "count": 1,
	  "self": "/users?email=ketsu@thoughtworks.com&page=1&per_page=30",
	  "first": "/users?email=ketsu@thoughtworks.com&page=1&per_page=30",
	  "last": "/users?email=ketsu@thoughtworks.com&page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	    {
	      "id": "47631d42-25d1-4fde-a8b5-02d94f0d616d",
	      "email": "ketsu@thoughtworks.com",
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/users/47631d42-25d1-4fde-a8b5-02d94f0d616d"
	        }
	      ]
	    }
	  ]
	}
	`
	var getNoUserByEmailResponse = `
	{
	  "count": 0,
	  "self": "/users?email=noexist@thoughtworks.com&page=1&per_page=30",
	  "first": "/users?email=noexist@thoughtworks.com&page=1&per_page=30",
	  "last": "/users?email=noexist@thoughtworks.com&page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	  ]
	}
	`

	var getUserByFingerprintResponse = `
	{
	  "count": 1,
	  "self": "/users?fingerprint=ketsu@thoughtworks.com&page=1&per_page=30",
	  "first": "/users?fingerprint=ketsu@thoughtworks.com&page=1&per_page=30",
	  "last": "/users?fingerprint=ketsu@thoughtworks.com&page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	    {
	      "id": "47631d42-25d1-4fde-a8b5-02d94f0d616d",
	      "email": "ketsu@thoughtworks.com",
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/users/47631d42-25d1-4fde-a8b5-02d94f0d616d"
	        }
	      ]
	    }
	  ]
	}
	`
	var getNoUserByFingerprintResponse = `
	{
	  "count": 0,
	  "self": "/users?fingerprint=noexist@thoughtworks.com&page=1&per_page=30",
	  "first": "/users?fingerprint=noexist@thoughtworks.com&page=1&per_page=30",
	  "last": "/users?fingerprint=noexist@thoughtworks.com&page=1&per_page=30",
	  "prev": null,
	  "next": null,
	  "items": [
	  ]
	}
	`

	var getUserByEmailRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users?email=" + userEmail,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getUserByEmailResponse,
		},
	}

	var getNoUserByEmailRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users?email=noexist@thoughtworks.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getNoUserByEmailResponse,
		},
	}

	var getUserByFingerprintRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users?fingerprint=" + userEmail,
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getUserByFingerprintResponse,
		},
	}

	var getNoUserByFingerprintRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/users?fingerprint=noexist@thoughtworks.com",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getNoUserByFingerprintResponse,
		},
	}

	var createUserRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo UserRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewUserRepository(configRepo, gateway)
		return
	}

	var defaultUserParams = func() UserParams {
		return UserParams{
			Email:    userEmail,
			Password: "123456",
		}
	}

	It("should able to create an user", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{createUserRequest, getUserRequest})
		defer ts.Close()

		err := repo.Create(defaultUserParams())
		Expect(err).To(BeNil())
	})

	It("should able to get an user", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserRequest})
		defer ts.Close()

		createdUser, err := repo.GetUser(userId)
		Expect(err).To(BeNil())
		Expect(createdUser.Email()).To(Equal(userEmail))
		Expect(createdUser.Links()).NotTo(BeNil())
	})

	It("should able to get an user by email", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserByEmailRequest, getNoUserByEmailRequest})
		defer ts.Close()

		users, err := repo.GetUserByEmail(userEmail)
		Expect(err).To(BeNil())
		Expect(users.Count()).To(Equal(1))
		Expect(users.Items()[0].Id()).To(Equal(userId))
		Expect(users.Items()).NotTo(BeNil())

		users, err = repo.GetUserByEmail("noexist@thoughtworks.com")
		Expect(err).ShouldNot(BeNil())
	})

	It("should able to get an user by fingerprint", func() {
		ts, _, repo := createUserRepository([]testnet.TestRequest{getUserByFingerprintRequest, getNoUserByFingerprintRequest})
		defer ts.Close()

		users, err := repo.GetUserByFingerprint(userEmail)
		Expect(err).To(BeNil())
		Expect(users.Count()).To(Equal(1))
		Expect(users.Items()[0].Id()).To(Equal(userId))
		Expect(users.Items()).NotTo(BeNil())

		users, err = repo.GetUserByFingerprint("noexist@thoughtworks.com")
		Expect(err).ShouldNot(BeNil())
	})
})
