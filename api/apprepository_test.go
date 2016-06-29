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

var _ = Describe("Apps", func() {
	var createAppRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo AppRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewAppRepository(configRepo, gateway)
		return
	}

	var defaultAppParams = func() AppParams {
		name := "ketsu"

		return AppParams{
			Name:      name,
			Stack:     "/stacks/stackid",
		}
	}

	It("should able to create an app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuCreate(), fixtures.KetsuDetail()})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		Expect(err).To(BeNil())
		Expect(createdApp.Id()).To(Equal("ketsu"))
		Expect(createdApp.Links()).NotTo(BeNil())
	})

	It("should able to get an app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuDetail()})
		defer ts.Close()

		createdApp, err := repo.GetApp("ketsu")
		Expect(err).To(BeNil())
		Expect(createdApp.Id()).To(Equal("ketsu"))
		Expect(createdApp.GetEnvs()["ENV"]).To(Equal("PRODUCTION"))
		Expect(createdApp.Links()).NotTo(BeNil())
		Expect(createdApp.Links().Self()).NotTo(BeNil())
	})

	It("should able to get apps", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.AppList()})
		defer ts.Close()

		createdApps, err := repo.GetApps()
		Expect(err).To(BeNil())
		Expect(createdApps.Count()).To(Equal(2))
		Expect(createdApps.Items()[0].Id()).To(Equal("ketsu"))
		Expect(createdApps.Items()[0].Links()).NotTo(BeNil())
	})

	It("should able to delete apps", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuDestroy()})
		defer ts.Close()

		err := repo.Delete("ketsu")
		Expect(err).To(BeNil())
	})

	It("should able to get collaborators", func() {
		userId := "47631d4225d14fdea8b502d94f0d616d"

		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuCollaborators(), fixtures.EmptyCollaborators()})
		defer ts.Close()

		users, err := repo.GetCollaborators("ketsu")
		Expect(err).To(BeNil())
		Expect(len(users)).To(Equal(1))
		Expect(users[0].Id()).To(Equal(userId))
		Expect(users).NotTo(BeNil())

		users, err = repo.GetCollaborators("empty")
		Expect(err).Should(BeNil())
	})

	It("should able to create collaborator", func() {
		userEmail := "test@tw.com"

		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.CollaboratorsAdd()})
		defer ts.Close()

		err := repo.AddCollaborator("ketsu", CreateCollaboratorParams{
			Email: userEmail,
		})
		Expect(err).To(BeNil())
	})

	It("should able to remove collaborator", func() {
		userId := "abc"
		appId := "ketsu"

		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.CollaboratorsRemove()})
		defer ts.Close()

		err := repo.RemoveCollaborator(appId, userId)
		Expect(err).To(BeNil())
	})

	It("should able to transfer to other user", func() {
		userEmail := "otheruser@tw.com"
		appId := "ketsu"

		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuTransferToUser()})
		defer ts.Close()

		err := repo.TransferToUser(appId, userEmail)
		Expect(err).To(BeNil())
	})


	It("should able to transfer to org", func() {
		orgName := "tw-test"
		appId := "ketsu"

		ts, _, repo := createAppRepository([]testnet.TestRequest{fixtures.KetsuTransferToOrg()})
		defer ts.Close()

		err := repo.TransferToOrg(appId, orgName)
		Expect(err).To(BeNil())
	})
})
