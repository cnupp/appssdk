package stub_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStub(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stub Suite")
}
