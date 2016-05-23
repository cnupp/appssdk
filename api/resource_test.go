package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	"github.com/sjkyspa/stacks/controller/api/fixtures"
	"net/http"
)

var _ = Describe("Resource", func() {
	Context("Authorized User", func() {
		It("should able to get the build by the uri", func(done Done) {
			ts, _ := testnet.NewServer([]testnet.TestRequest{
				fixtures.KetsuBuild(), fixtures.KetsuDetail(), fixtures.SuccessKetsuBuild(func(r *http.Request) {  }),
			})
			defer ts.Close()

			configRepo := testconfig.NewConfigRepository()
			configRepo.SetApiEndpoint(ts.URL)
			configRepo.SetAuth("auth")
			gateway := net.NewCloudControllerGateway(configRepo)

			resource, err := NewResource(configRepo, gateway).GetResourceByURI("/apps/ketsu/builds/86e03fc8b63941669a20dbae948bdfc8")
			Expect(err).To(BeNil())
			Expect(resource).NotTo(BeNil())
			Expect(resource.(Build).GitSha()).To(Equal("60bc43aa"))
			Expect(resource.(Build).Success()).To(BeNil())
			close(done)
		}, 60)

		It("should able to get builds by the uri", func(done Done) {
			ts, _ := testnet.NewServer([]testnet.TestRequest{
				fixtures.KetsuBuilds(),
			})

			defer ts.Close()

			configRepo := testconfig.NewConfigRepository()
			configRepo.SetApiEndpoint(ts.URL)
			configRepo.SetAuth("auth")
			gateway := net.NewCloudControllerGateway(configRepo)

			resource, err := NewResource(configRepo, gateway).GetResourceByURI("/apps/ketsu/builds")
			Expect(err).To(BeNil())
			Expect(resource).NotTo(BeNil())
			Expect(resource.(Builds).Count()).To(Equal(1))
			close(done)
		}, 60)

		It("should able to get apps by the uri", func(done Done) {
			ts, _ := testnet.NewServer([]testnet.TestRequest{
				fixtures.AppList(),
			})

			defer ts.Close()

			configRepo := testconfig.NewConfigRepository()
			configRepo.SetApiEndpoint(ts.URL)
			configRepo.SetAuth("auth")
			gateway := net.NewCloudControllerGateway(configRepo)

			resource, err := NewResource(configRepo, gateway).GetResourceByURI("/apps")
			Expect(err).To(BeNil())
			Expect(resource).NotTo(BeNil())
			Expect(resource.(Apps).Count()).To(Equal(2))
			close(done)
		}, 60)
	})

})
