package stub_test

import (
	. "github.com/sjkyspa/stacks/controller/api/testhelpers/stub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"bytes"
)

var _ = Describe("Stub", func() {
	It("should able to start the stub server", func() {
		server, _ := NewStub([]TestRequest{
			TestRequest{
				Method: "GET",
				Path: "/path",
				Response: TestResponse{
					Status: 200,
				},
			},
		})

		var body []byte;
		req, err :=http.NewRequest("GET", server.URL + "/path", bytes.NewBuffer(body))
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusOK))
	})
})
