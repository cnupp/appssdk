package api_test

import (
	. "github.com/cnupp/cnup/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/cnupp/cnup/controller/api/net"
	testconfig "github.com/cnupp/cnup/controller/api/testhelpers/config"
	testnet "github.com/cnupp/cnup/controller/api/testhelpers/net"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Apps", func() {
	var getOrgMembersRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/orgs/tw-test/members",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"accept": {"application/json"},
			},
			Body: `
			[]
			`,
		},
	}
	var listOrgAppsRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/orgs/tw-test/apps",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"accept": {"application/json"},
			},
			Body: `
			[]
			`,
		},
	}
	var createOrgRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/orgs",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/orgs/tw-test"},
			},
		},
	}

	var addOrgAppRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/orgs/tw-test/apps",
		Response: testnet.TestResponse{
			Status: 201,
			Header: http.Header{
				"accept":   {"application/json"},
				"Location": {"/orgs/tw-test/apps/abc"},
			},
		},
	}

	var rmOrgMemberRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/orgs/tw/members/abc",
		Response: testnet.TestResponse{
			Status: 204,
		},
	}

	var addOrgMemberRequest = testnet.TestRequest{
		Method: "POST",
		Path:   "/orgs/tw/members",
		Response: testnet.TestResponse{
			Status: 201,
		},
	}

	var getOrgResponse = `
	{
	  "name": "tw-test",
	  "links": [
		{
		  "rel": "self",
		  "uri": "/orgs/tw-test"
		},
		{
		  "rel": "members",
		  "uri": "/orgs/tw-test/members"
		},
		{
		  "rel": "apps",
		  "uri": "/orgs/tw-test/apps"
		}
	  ]
	}
	`

	var getOrgRequest = testnet.TestRequest{
		Method: "GET",
		Path:   "/orgs/tw-test",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
			Body: getOrgResponse,
		},
	}

	var destroyOrgRequest = testnet.TestRequest{
		Method: "DELETE",
		Path:   "/orgs/tw-test",
		Response: testnet.TestResponse{
			Status: 200,
			Header: http.Header{
				"Content-Type": {"application/json"},
			},
		},
	}

	var createOrgRepository = func(requests []testnet.TestRequest) (ts *httptest.Server, handler *testnet.TestHandler, repo OrgRepository) {
		ts, handler = testnet.NewServer(requests)
		configRepo := testconfig.NewRepositoryWithDefaults()
		configRepo.SetApiEndpoint(ts.URL)
		gateway := net.NewCloudControllerGateway(configRepo)
		repo = NewOrgRepository(configRepo, gateway)
		return
	}

	var defaultOrgParams = func() OrgParams {
		name := "tw-test"

		return OrgParams{
			Name: name,
		}
	}

	It("should able to create an org", func() {
		ts, _, repo := createOrgRepository([]testnet.TestRequest{createOrgRequest, getOrgRequest})
		defer ts.Close()

		createdOrg, err := repo.Create(defaultOrgParams())
		Expect(err).To(BeNil())
		Expect(createdOrg.Name()).To(Equal("tw-test"))
		Expect(createdOrg.Links()).NotTo(BeNil())
	})

	It("should able to get an app", func() {
		ts, _, repo := createOrgRepository([]testnet.TestRequest{getOrgRequest})
		defer ts.Close()

		createdApp, err := repo.GetOrg("tw-test")
		Expect(err).To(BeNil())
		Expect(createdApp.Name()).To(Equal("tw-test"))
		Expect(createdApp.Links()).NotTo(BeNil())
	})

	It("should list all members", func() {
		ts, _, repo := createOrgRepository([]testnet.TestRequest{getOrgMembersRequest})
		defer ts.Close()

		_, err := repo.GetOrgMembers("tw-test")
		Expect(err).To(BeNil())
	})

	It("should add members", func() {
		userEmail := "user@tw.com"
		orgName := "tw"
		ts, _, repo := createOrgRepository([]testnet.TestRequest{addOrgMemberRequest})
		defer ts.Close()

		err := repo.AddMember(orgName, userEmail)
		Expect(err).To(BeNil())
	})

	It("should remove members", func() {
		userId := "abc"
		orgName := "tw"
		ts, _, repo := createOrgRepository([]testnet.TestRequest{rmOrgMemberRequest})
		defer ts.Close()

		err := repo.RmMember(orgName, userId)
		Expect(err).To(BeNil())
	})

	It("should list apps", func() {
		orgName := "tw-test"
		ts, _, repo := createOrgRepository([]testnet.TestRequest{listOrgAppsRequest})
		defer ts.Close()

		_, err := repo.GetApps(orgName)
		Expect(err).To(BeNil())
	})

	It("should add app", func() {
		orgName := "tw-test"
		appName := "abc"
		ts, _, repo := createOrgRepository([]testnet.TestRequest{addOrgAppRequest})
		defer ts.Close()

		err := repo.AddApp(orgName, appName)
		Expect(err).To(BeNil())
	})

	It("should able to delete org", func() {
		ts, _, repo := createOrgRepository([]testnet.TestRequest{destroyOrgRequest})
		defer ts.Close()

		err := repo.Delete("tw-test")
		Expect(err).To(BeNil())
	})
})
