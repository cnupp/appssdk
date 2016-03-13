package util_test

import (
	. "github.com/sjkyspa/stacks/controller/api/util"

	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	It("should able to split the uri correctly", func() {
		id, _ := IDFromURI("/apps/ketsu/id")
		Expect(id).To(Equal("id"))
	})
})
