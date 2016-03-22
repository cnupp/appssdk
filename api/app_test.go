package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sjkyspa/stacks/controller/api/api"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
	"strings"
)

var _ = Describe("App", func() {
	var createAppRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/apps",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu"},
			},
		},
	}
	var getAppResponse = `
	{
	  "id": "b78dba51-8daf-4fe9-9345-c7ab582c3387",
	  "name": "ketsu",
	  "memory": 30,
	  "disk": 30,
	  "instances": 1,
	  "links": [
		{
		  "rel": "self",
		  "uri": "/apps/ketsu"
		},
		{
		  "rel": "env",
		  "uri": "/apps/ketsu/env"
		},
		{
		  "rel": "routes",
		  "uri": "/apps/ketsu/routes"
		},
		{
		  "rel": "builds",
		  "uri": "/apps/ketsu/builds"
		},
		{
		  "rel": "releases",
		  "uri": "/apps/ketsu/releases"
		},
		{
		  "rel": "stack",
		  "uri": "/stacks/javajersey"
		}
	  ]
	}
	`

	var getAppRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppResponse,
		},
	}

	var getStackResponse = `
	{
        "id": "74a052c9-76b3-44a1-ac0b-666faa1223b6",
        "name": "javajersey",
        "links": [
          {
            "rel": "self",
            "uri": "/stacks/javajersey"
          }
        ]
	}
	`

	var getStackRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks/javajersey",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getStackResponse,
		},
	}

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

	var getAppBuildResponse = `
	{
	  "created": "1451953908",
	  "git_sha": "60bc43aa",
	  "id": "1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
	  "status": "NEW",
	  "app": {
		"name": "ketsu"
	  },
	  "links": [
		{
		  "rel": "self",
		  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"
		},
		{
		  "rel": "app",
		  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387"
		},
		{
		  "rel": "verifies",
		  "uri": "/apps/b78dba51-8daf-4fe9-9345-c7ab582c3387/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/verifies"
		}
	  ]
	}
	`
	var getRoutesResponse = `
	{
	  "count": 2,
	  "self": "/apps/ketsu/routes",
	  "first": "/apps/ketsu/routes?page=1&per_page=30",
	  "last": "/apps/ketsu/routes?page=2&per_page=30",
	  "prev": "",
	  "next": "/apps/ketsu/routes?page=2&per_page=30",
	  "items": [
	    {
	      "id": "8399de76-eeef-418d-a567-75253b03c4ec",
	      "path": "/path",
	      "domain": {
	        "name": "deepi.cn"
	      },
	      "app": {
	        "name": "ketsu"
	      },
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/apps/ketsu/routes/8399de76-eeef-418d-a567-75253b03c4ec"
	        },
	        {
	          "rel": "app",
	          "uri": "/apps/ketsu"
	        }
	      ]
	    }
	  ]
	}

	`

	var getRoutesOnSecondPageResponse = `
	{
	  "count": 2,
	  "self": "/apps/ketsu/routes?page=2&per_page=30",
	  "first": "/apps/ketsu/routes?page=1&per_page=30",
	  "last": "/apps/ketsu/routes?page=2&per_page=30",
	  "prev": "/apps/ketsu/routes?page=1&per_page=30",
	  "next": "",
	  "items": [
	    {
	      "id": "8399de76-eeef-418d-a567-75253b03c4ec",
	      "path": "/path",
	      "domain": {
	        "name": "deepi.cn"
	      },
	      "app": {
	        "name": "ketsu"
	      },
	      "links": [
	        {
	          "rel": "self",
	          "uri": "/apps/ketsu/routes/8399de76-eeef-418d-a567-75253b03c4ec"
	        },
	        {
	          "rel": "app",
	          "uri": "/apps/ketsu"
	        }
	      ]
	    }
	  ]
	}
	`

	var getAppBuildRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getAppBuildResponse,
		},
	}

	var bindRouteWithAppRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/routes",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/apps/ketsu/routes/8399de76-eeef-418d-a567-75253b03c4ec"},
			},
		},
	}

	var getRoutesWithAppRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/routes",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getRoutesResponse,
		},
	}

	var getRoutesOnNextPageWithAppRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/routes",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getRoutesOnSecondPageResponse,
		},
	}

	var switchToAnotherStackRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/switch-stack",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var logBody = "2016-03-17 05:38:19 DEBUG PooledDataSource:316  - PooledDataSource forcefully closed/removed all connections."

	var getAppBuildLog = testnet.TestRequest{
		Method: "GET",
		Path:   "/apps/ketsu/builds/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45/log",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"text/plain"},
			},
			Body: logBody,
		},
	}

	var unbindRouteWithAppRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/apps/ketsu/routes/test.tw.com/path",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var createEnvRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/apps/ketsu/env",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}

	var deleteEnvRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/apps/ketsu/env",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}

	It("should able to get builds for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest, getAppBuildsRequest})
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
		build, err := createdApp.GetBuild("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")
		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
	})

	It("should able to get one build from uri for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{createAppRequest, getAppRequest, getAppBuildRequest})
		defer ts.Close()

		createdApp, err := repo.Create(defaultAppParams())
		build, err := createdApp.GetBuildByURI("/apps/ketsu/1a5abd6c-49b6-4c6a-b47c-d75fedec0a45")
		Expect(err).To(BeNil())
		Expect(build.GitSha()).To(Equal("60bc43aa"))
	})

	It("should able to create one build for app", func() {
		ts, _, repo := createAppRepository([]testnet.TestRequest{getAppRequest, createAppBuildRequest, getAppBuildRequest})
		defer ts.Close()

		remoteApp, err := repo.GetApp("ketsu")
		build, err := remoteApp.CreateBuild(defaultBuildParams())

		Expect(err).To(BeNil())
		Expect(build.Id()).To(Equal("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45"))
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
		Expect(routes.Count()).To(Equal(2))
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
		Expect(routesOnSecondPage.Count()).To(Equal(2))
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

		log, err := app.GetLogForTests("1a5abd6c-49b6-4c6a-b47c-d75fedec0a45", "build", 15, "2016-03-16T09:53:23.594Z")
		Expect(err).To(BeNil())
		Expect(strings.TrimSpace(log)).To(Equal(logBody))

	})
})
