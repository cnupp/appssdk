package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"
	"github.com/sjkyspa/stacks/controller/api/net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/fixtures"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http/httptest"
)

var _ = Describe("Builds", func() {
	It("should able to create an build for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuildCreate(), fixtures.KetsuBuild()})
		defer ts.Close()

		build, err := mapper.Create(AppModel{
			ID: "ketsu",
		}, BuildParams{
			GitSha: "60bc43aa",
			User:   "user",
		})

		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
		Expect(build.Verify().Status()).To(Equal("NEW"))
	})

	It("should able to get all builds for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuilds()})
		defer ts.Close()

		builds, err := mapper.GetBuilds(AppModel{
			ID: "ketsu",
		})

		Expect(err).To(BeNil())
		Expect(builds.Count()).To(Equal(1))
		Expect(len(builds.Items())).To(Equal(1))
		Expect(builds.Items()[0].GitSha).To(Equal("60bc43aa"))
		Expect(builds.Items()[0].ID).To(Equal("86e03fc8b63941669a20dbae948bdfc8"))
		Expect(builds.Items()[0].Links).To(Not(BeNil()))
		Expect(builds.Items()[0].Status).To(Equal("NEW"))
	})

	It("should able to get one build for app", func() {
		ts, _, mapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuild()})
		defer ts.Close()

		build, err := mapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "86e03fc8b63941669a20dbae948bdfc8")

		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
		Expect(build.Id()).To(Equal("86e03fc8b63941669a20dbae948bdfc8"))
		Expect(build.Links()).To(Not(BeNil()))
		Expect(build.Status()).To(Equal("NEW"))
	})
})

func createBuildMapper(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo BuildMapper) {
	ts, handler = testnet.NewServer(requests)
	configRepo := testconfig.NewRepositoryWithDefaults()
	configRepo.SetApiEndpoint(ts.URL)
	gateway := net.NewCloudControllerGateway(configRepo)
	repo = NewBuildMapper(configRepo, gateway)
	return
}
