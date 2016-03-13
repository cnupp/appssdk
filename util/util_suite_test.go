package util_test

import (
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util Suite")
}
