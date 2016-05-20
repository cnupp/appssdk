package stub_test

import (
	. "github.com/sjkyspa/stacks/controller/api/testhelpers/stub"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"bytes"
	"io/ioutil"
)

var _ = Describe("Stub", func() {
	It("should able to start the stub server", func() {
		server, _ := NewStub([]TestRequest{
			TestRequest{
				Method: "GET",
				Path: "/path",
				Response: TestResponse{
					Status: 200,
					Header: http.Header{
						"Content-Type": {"application/json"},
						"Set-Cookie": {"cookie=cookie"},
					},
					Body: "content",
				},
			},
		})

		var body []byte;
		req, err := http.NewRequest("GET", server.URL + "/path", bytes.NewBuffer(body))
		Expect(err).To(BeNil())

		client := http.Client{}

		res, err := client.Do(req)
		Expect(err).To(BeNil())

		Expect(res.StatusCode).To(Equal(http.StatusOK))
		Expect(res.Header.Get("Content-Type")).To(Equal("application/json"))
		Expect(res.Header.Get("Set-Cookie")).To(Equal("cookie=cookie"))

		bodyInBytes, err := ioutil.ReadAll(res.Body)
		Expect(err).To(BeNil())
		Expect(string(bodyInBytes)).To(Equal("content"))
	})
})
