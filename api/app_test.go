package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/cnupp/cnup/controller/api/api"
	"github.com/cnupp/cnup/controller/api/fixtures"
	"github.com/cnupp/cnup/controller/api/net"
	testconfig "github.com/cnupp/cnup/controller/api/testhelpers/config"
	testnet "github.com/cnupp/cnup/controller/api/testhelpers/net"
	"net/http/httptest"
)

var _ = Describe("App", func() {
	var getAppPermissionRequest = fixtures.KaylaPermissionOnKetsu()
	var createAppRequest = fixtures.KetsuCreate()
	var getAppRequest = fixtures.KetsuDetail()
	var getStackRequest = fixtures.KetsuStackDetail()
	var getAppBuildRequest = fixtures.KetsuBuild()
	var bindRouteWithAppRequest = fixtures.KetsuRoutesBind()
	var getRoutesWithAppRequest = fixtures.KetsuRoutes()
	var getRoutesOnNextPageWithAppRequest = fixtures.KetsuRoutesSecondPage()
	var switchToAnotherStackRequest = fixtures.KetsuStackUpdate()
	var getAppBuildLog = fixtures.KetsuBuildLog()
	var unbindRouteWithAppRequest = fixtures.KetsuRoutesUnbind()
	var createEnvRequest = fixtures.KetsuEnvCreate()
	var deleteEnvRequest = fixtures.KetsuEnvUpdate()

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
			Name:  name,
			Stack: "/stacks/stackid",
		}
	}

	var defaultBuildParams = func() BuildParams {
		return BuildParams{
			GitSha: "60bc43aa",
			User:   "/users/ketsu",
		}
	}

	var defaultRouteParams = func() AppRouteParams {
		return AppRouteParams{
			Route: "test.deepi.cn/path",
		}
	}

	It("should able to get builds for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest, fixtures.KetsuBuilds()})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		builds, err := createdApp.GetBuilds()
		Expect(err).To(BeNil())
		Expect(builds.Count()).To(Equal(1))
		Expect(builds.Items()[0].GitSha).To(Equal("60bc43aa"))
	})

	It("should able to get one build for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest, getAppBuildRequest})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		build, err := createdApp.GetBuild("86e03fc8b63941669a20dbae948bdfc8")
		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
	})

	It("should able to get one build from uri for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest, getAppBuildRequest})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		build, err := createdApp.GetBuildByURI("/apps/ketsu/86e03fc8b63941669a20dbae948bdfc8")
		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
	})

	It("should able to create one build for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, fixtures.KetsuBuildCreate(), getAppBuildRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		build, err := remoteApp.CreateBuild(defaultBuildParams())

		Expect(err).To(BeNil())
		Expect(build.Id()).To(Equal("86e03fc8b63941669a20dbae948bdfc8"))
		Expect(build.GitSha()).To(Equal("60bc43aa"))
	})

	It("should able to bind route with the app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, bindRouteWithAppRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		err = remoteApp.BindWithRoute(defaultRouteParams())

		Expect(err).To(BeNil())
	})

	It("should able to get routes associated with the app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getRoutesWithAppRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		routes, err := remoteApp.GetRoutes()

		Expect(err).To(BeNil())
		Expect(routes.Count()).To(Equal(31))
		Expect(routes.Items()[0].PathField).NotTo(BeNil())
		Expect(routes.Items()[0].IDField).NotTo(BeNil())
	})

	It("should able to get second page of routes associated with the app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getRoutesWithAppRequest, getRoutesOnNextPageWithAppRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		routesOnFirstPage, err := remoteApp.GetRoutes()
		routesOnSecondPage, err := routesOnFirstPage.Next()

		Expect(err).To(BeNil())
		Expect(routesOnSecondPage.Count()).To(Equal(31))
		Expect(routesOnSecondPage.Items()[0].PathField).NotTo(BeNil())
		Expect(routesOnSecondPage.Items()[0].IDField).NotTo(BeNil())
	})

	It("should able to get prev page of routes associated with the app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getRoutesWithAppRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		routesOnFirstPage, err := remoteApp.GetRoutes()
		routesOnPreviousPage, err := routesOnFirstPage.Prev()

		Expect(err).To(BeNil())
		Expect(routesOnPreviousPage).To(BeNil())
	})

	It("should able to delete bind association between route and app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, unbindRouteWithAppRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		err = remoteApp.UnbindRoute("test.tw.com/path")

		Expect(err).To(BeNil())
	})

	It("should able to get stack of app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getStackRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		stack, err := remoteApp.GetStack()

		Expect(err).To(BeNil())
		Expect(stack.Name()).To(Equal("javajersey"))
		Expect(stack.Links()).NotTo(BeNil())
	})

	It("should able to create env", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, createEnvRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		envs := make(map[string]interface{})
		envs["ENV"] = "PRODUCTION"
		err = remoteApp.SetEnv(envs)
		Expect(err).To(BeNil())
	})

	It("should able to delete env", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, deleteEnvRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		err = remoteApp.UnsetEnv([]string{"ENV"})
		Expect(err).To(BeNil())
	})

	It("should able to switch to another stack for an application", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, switchToAnotherStackRequest})
		defer ts.Close()

		app, err := repo.GetApp("ketsu")
		newStack := UpdateStackParams{
			Stack: "newStack",
		}
		err = app.SwitchStack(newStack)
		Expect(err).To(BeNil())

	})

	It("should able to get build logs", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getAppBuildLog})
		defer ts.Close()

		app, err := repo.GetApp("ketsu")
		Expect(err).To(BeNil())

		log, err := app.GetLogForTests("86e03fc8b63941669a20dbae948bdfc8", "build", 15, 0)
		Expect(err).To(BeNil())
		Expect(log.ItemsField[0].MessageField).To(Equal("init successful"))

	})

	It("should able to get app permissions", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, getAppPermissionRequest})
		defer ts.Close()

		app, err := repo.GetApp("ketsu")
		Expect(err).To(BeNil())

		permission, err := app.GetPermissions("abcd")
		Expect(err).To(BeNil())
		Expect(permission.Write).To(BeTrue())
	})
})
