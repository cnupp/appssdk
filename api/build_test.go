package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/fixtures"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	. "github.com/sjkyspa/stacks/controller/api/testhelpers/matchers"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Build", func() {
	var createBuildMapper = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo BuildMapper) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewBuildMapper(configRepo, gateway)
		return
	}

	It("should able to success the build", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuild(), fixtures.SuccessKetsuBuild(func(r *http.Request) {}), fixtures.KetsuDetail()})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "86e03fc8b63941669a20dbae948bdfc8")
		Expect(build.IsSuccess()).To(BeFalse())

		err := build.Success()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to fail the build", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuild(), fixtures.FailKetsuBuild(func(r *http.Request) {}), fixtures.KetsuDetail()})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "86e03fc8b63941669a20dbae948bdfc8")

		err := build.Fail()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to set verify success", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuild(), fixtures.SuccessKetsuBuild(func(r *http.Request) {}), fixtures.SuccessKetsuVerify(func(r *http.Request) {}), fixtures.KetsuDetail()})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "86e03fc8b63941669a20dbae948bdfc8")
		Expect(build.IsVerifySuccess()).To(BeFalse())
		err := build.Success()
		Expect(err).To(BeNil())
		err = build.VerifySuccess()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})

	It("should able to set verify fail", func() {
		ts, handler, buildMapper := createBuildMapper([]testnet.TestRequest{fixtures.KetsuBuild(), fixtures.SuccessKetsuBuild(func(r *http.Request) {}), fixtures.FailKetsuVerify(func(r *http.Request) {}), fixtures.KetsuDetail()})
		defer ts.Close()

		build, _ := buildMapper.GetBuild(AppModel{
			ID: "ketsu",
		}, "86e03fc8b63941669a20dbae948bdfc8")
		Expect(build.IsVerifySuccess()).To(BeFalse())
		err := build.Success()
		Expect(err).To(BeNil())
		err = build.VerifyFail()
		Expect(handler).To(HaveAllRequestsCalled())
		Expect(err).To(BeNil())
	})
})
