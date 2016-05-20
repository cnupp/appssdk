package stub

import (
	"net/http"
	"net/http/httptest"
	"github.com/onsi/ginkgo"
)

type TestRequest struct {
	Method   string
	Path     string
	Header   http.Header
	Matcher  RequestMatcher
	Response TestResponse
}

type RequestMatcher func(*http.Request)

type TestResponse struct {
	Body   string
	Status int
	Header http.Header
}

type TestHandler struct {
	Requests  []TestRequest
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer ginkgo.GinkgoRecover()

}

func NewStub(requests []TestRequest) (*httptest.Server, *TestHandler) {
	handler := &TestHandler{Requests: requests}
	return httptest.NewServer(handler), handler
}