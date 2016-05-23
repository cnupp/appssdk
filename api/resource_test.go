package api_test

import (
	. "github.com/sjkyspa/stacks/controller/api/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sjkyspa/stacks/controller/api/testhelpers/config"
)

var _ = Describe("Resource", func() {
	Context("Authorized User", func() {
		config := config.NewConfigRepository()
		config.SetAuth("auth")

		It("should able to get the build by the uri", func(done Done) {

			resource, err := NewResource(config).GetResourceByURI("/apps/test/builds")
			Expect(err).To(BeNil())
			Expect(resource).NotTo(BeNil())
			close(done)
		}, 60)
	})

})
