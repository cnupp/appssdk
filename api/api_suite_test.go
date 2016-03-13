package api_test

import (
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}
