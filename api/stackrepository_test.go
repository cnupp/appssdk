package api

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/net"
	testconfig "github.com/sjkyspa/stacks/controller/api/testhelpers/config"
	testnet "github.com/sjkyspa/stacks/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Stacks", func() {
	var unPublishStackRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6/unpublished",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}
	var publishStackRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6/published",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}
	var createStackRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/stacks",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6"},
			},
		},
	}

	var updateStackRequest = testnet.TestRequest{
		Method: "PUT",
		Path:   "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6",
		Response: testnet.TestResponse{
			Status: 200,
		},
	}

	var getStackResponse = `
	{
	  "name": "javajersey",
	  "id": "74a052c9-76b3-44a1-ac0b-666faa1223b6",
	  "services" :{
	  	"web":
	  	{
	  		"build":{
	  			"image":"192.168.99.100:5000/javajersey-test-build",
	  			"mem":1024,
	  			"cpus":1.0
	  		},
	  		"verify":{
	  			"image":"abc",
	  			"mem":512,
	  			"cpus":1.0
	  		},
	  		"name":"web",
	  		"main":true,
	  		"localhostPort":4876,
	  		"image":null,
	  		"links":["db"],
	  		"environment":{},
	  		"mem":512,
	  		"cpus":0.4,
	  		"instances":1,
	  		"expose":8088
		},
		"db":
		{
			"name":"db",
			"main":false,
			"localhostPort":26468,
			"image":"tutum/mysql",
			"links":[],
			"environment":{"EXTRA_OPTS":"--lower_case_table_names=1","MYSQL_PASS":"mysql","MYSQL_PASSWORD":"mysql","MYSQL_USER":"mysql","ON_CREATE_DB":"stacks"},
			"mem":256,
			"cpus":0.2,
			"instances":1,
			"expose":3306,
			"volumes":[{"container":"/var/lib/mysql","host":"db","mode":"RW"}]
		}
	  }
	}
	`
	var getStackRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getStackResponse,
		},
	}

	var getStacksResponse = `
	{
  		"count": 2,
  		"self": "/stacks?page=1&per_page=30",
  		"first": "/stacks?page=1&per_page=30",
  		"last": "/stacks?page=1&per_page=30",
  		"prev": null,
  		"next": null,
  		"items": [
    		{
    			"id": "74a052c9-76b3-44a1-ac0b-666faa1223b6",
      			"name": "javajersey",
      			"links": [
        			{
          			"rel": "self",
          			"uri": "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6"
	        		}
	      		]
	    	},
	    	{
	    		"id": "712w052c9-76b3-44a1-ac0b-666faa7638a8",
      			"name": "rubymongo",
      			"links": [
        			{
          			"rel": "self",
          			"uri": "/stacks/712w052c9-76b3-44a1-ac0b-666faa7638a8"
	        		}
	      		]
	    	}
 	 	]
	}
	`

	var getStacksRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getStacksResponse,
		},
	}

	var getStackByNameResponse = `
	{
  		"count": 1,
  		"self": "/stacks?name=java&page=1&per_page=30",
  		"first": "/stacks?name=java&page=1&per_page=30",
  		"last": "/stacks?name=java&page=1&per_page=30",
  		"prev": null,
  		"next": null,
  		"items": [
    		{
    			"id": "74a052c9-76b3-44a1-ac0b-666faa1223b6",
      			"name": "java",
      			"links": [
        			{
          			"rel": "self",
          			"uri": "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6"
	        		}
	      		]
	    	}
 	 	]
	}
	`

	var getNoStackByNameResponse = `
	{
  		"count": 0,
  		"self": "/stacks?name=not-exist&page=1&per_page=30",
  		"first": null,
  		"last": null,
  		"prev": null,
  		"next": null,
  		"items": [
 	 	]
	}
	`

	var getStackByNameRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks?name=java",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getStackByNameResponse,
		},
	}

	var getNoStackByNameRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/stacks?name=not-exist",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getNoStackByNameResponse,
		},
	}

	var deleteStackRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/stacks/74a052c9-76b3-44a1-ac0b-666faa1223b6",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var createStackRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo StackRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewStackRepository(configRepo, gateway)
		return
	}

	var prepareStackRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo StackRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewStackRepository(configRepo, gateway)
		return
	}

	var defaultStackParams = func() map[string]interface{} {
		stack := make(map[string]interface{})

		stack["name"] = "javajersey"
		return stack
	}

	It("should able to publish a stack", func() {
		ts, _, repo := prepareStackRepository([]testnet.TestRequest{publishStackRequest})
		defer ts.Close()

		err := repo.Publish("74a052c9-76b3-44a1-ac0b-666faa1223b6")
		Expect(err).To(BeNil())
	})

	It("should able to unpublish a stack", func() {
		ts, _, repo := prepareStackRepository([]testnet.TestRequest{unPublishStackRequest})
		defer ts.Close()

		err := repo.UnPublish("74a052c9-76b3-44a1-ac0b-666faa1223b6")
		Expect(err).To(BeNil())
	})

	It("should able to create an stack", func() {
		ts, _, repo := createStackRepository([]testnet.TestRequest{createStackRequest, getStackRequest})
		defer ts.Close()

		createdStack, err := repo.Create(defaultStackParams())
		Expect(err).To(BeNil())
		Expect(createdStack.Name()).To(Equal("javajersey"))
		Expect(createdStack.Links()).NotTo(BeNil())
	})

	It("should able to update an stack", func() {
		ts, _, repo := prepareStackRepository([]testnet.TestRequest{updateStackRequest})
		defer ts.Close()

		err := repo.Update("74a052c9-76b3-44a1-ac0b-666faa1223b6", defaultStackParams())
		Expect(err).To(BeNil())
	})

	It("should able to get an stack", func() {
		ts, _, repo := createStackRepository([]testnet.TestRequest{getStackRequest})
		defer ts.Close()

		stack, err := repo.GetStack("74a052c9-76b3-44a1-ac0b-666faa1223b6")
		Expect(err).To(BeNil())
		Expect(stack.Id()).To(Equal("74a052c9-76b3-44a1-ac0b-666faa1223b6"))
		Expect(stack.Name()).To(Equal("javajersey"))
		Expect(stack.Links()).NotTo(BeNil())
		Expect(stack.GetBuildImage()).NotTo(BeNil())
		buildImage := stack.GetBuildImage()
		Expect(buildImage.Cpus).To(Equal(1.0))
		Expect(buildImage.Mem).To(Equal(1024))
		Expect(buildImage.Name).To(Equal("192.168.99.100:5000/javajersey-test-build"))
		Expect(stack.GetVerifyImage()).NotTo(BeNil())
		verifyImage := stack.GetVerifyImage()
		Expect(verifyImage.Cpus).To(Equal(1.0))
		Expect(verifyImage.Mem).To(Equal(512))
		Expect(verifyImage.Name).To(Equal("abc"))

		services := stack.GetServices()
		Expect(len(services)).To(Equal(2))
		Expect(len(services["web"].GetEnv())).To(Equal(0))
		Expect(services["web"].GetLinks()[0]).To(Equal("db"))
		Expect(services["web"].GetExpose()[0]).To(Equal(8088))

		webVolumes := services["web"].GetVolumes()
		Expect(len(webVolumes)).To(Equal(0))

		dbVolumes := services["db"].GetVolumes()
		Expect(len(dbVolumes)).To(Equal(1))
		Expect(dbVolumes[0].ContainerPath).To(Equal("/var/lib/mysql"))
		Expect(dbVolumes[0].HostPath).To(Equal("db"))
		Expect(dbVolumes[0].Mode).To(Equal("RW"))

		env := services["db"].GetEnv()
		Expect(env["EXTRA_OPTS"]).To(Equal("--lower_case_table_names=1"))
		Expect(env["MYSQL_PASS"]).To(Equal("mysql"))
	})

	It("should able to get stacks", func() {
		ts, _, repo := createStackRepository([]testnet.TestRequest{getStacksRequest})
		defer ts.Close()

		stacks, err := repo.GetStacks()
		Expect(err).To(BeNil())
		Expect(stacks.Count()).To(Equal(2))
		Expect(stacks.Items()[0].Id()).To(Equal("74a052c9-76b3-44a1-ac0b-666faa1223b6"))
		Expect(stacks.Items()[0].Name()).To(Equal("javajersey"))
		Expect(stacks.Items()[0].Links()).NotTo(BeNil())
	})

	It("should able to get stack by stack name", func() {
		ts, _, repo := createStackRepository([]testnet.TestRequest{getStackByNameRequest, getNoStackByNameRequest})
		defer ts.Close()

		stacks, err := repo.GetStackByName("java")
		Expect(err).To(BeNil())
		Expect(stacks.Count()).To(Equal(1))
		Expect(stacks.Items()[0].Id()).To(Equal("74a052c9-76b3-44a1-ac0b-666faa1223b6"))
		Expect(stacks.Items()[0].Name()).To(Equal("java"))
		Expect(stacks.Items()[0].Links()).NotTo(BeNil())

		stacks, err = repo.GetStackByName("not-exist")
		Expect(err).ShouldNot(BeNil())

	})

	It("should able to delete an stack", func() {
		ts, _, repo := createStackRepository([]testnet.TestRequest{deleteStackRequest})
		defer ts.Close()

		err := repo.Delete("74a052c9-76b3-44a1-ac0b-666faa1223b6")
		Expect(err).To(BeNil())
	})
})
