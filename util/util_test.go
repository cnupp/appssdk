package util_test

import (
	. "github.com/cnupp/cnup/controller/api/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	It("should able to split the uri correctly", func() {
		id, _ := IDFromURI("/apps/ketsu/id")
		Expect(id).To(Equal("id"))
	})
})
