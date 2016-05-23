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

var _ = Describe("Auths", func() {
	var authId = "47631d4225d14fdea8b502d94f0d616d"
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

	It("should able to get user by auth", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{fixtures.Auths()})
		defer ts.Close()

		createdUser, err := repo.Get()
		Expect(err).To(BeNil())
		Expect(createdUser.Email()).To(Equal(userEmail))
		Expect(createdUser.Links()).NotTo(BeNil())
	})

	It("should able to create an authorization with correct user and password", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{fixtures.Login()})
		defer ts.Close()

		createdAuth, err := repo.Create(defaultAuthParams())
		Expect(err).To(BeNil())
		Expect(createdAuth.UserEmail()).To(Equal(userEmail))
		Expect(createdAuth.Id()).To(Equal(authId))
	})

	It("should show error with wrong user password", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{fixtures.InvalidLogin()})
		defer ts.Close()

		createdAuth, err := repo.Create(defaultAuthParams())
		Expect(createdAuth).To(BeNil())
		Expect(err).ToNot(BeNil())
	})

	It("should able to delete an authorization", func() {
		ts, _, repo := createAuthRepository([]testnet.TestRequest{fixtures.Logout()})
		defer ts.Close()

		err := repo.Delete(authId)
		Expect(err).To(BeNil())
	})
})
