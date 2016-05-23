package net

import (
	"net/http"
	"net/http/httptest"
	"github.com/onsi/ginkgo"
	"fmt"
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
	Requests []TestRequest
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer ginkgo.GinkgoRecover()

	request := Filter(h.Requests, func(tq TestRequest) bool {
		return tq.Method == r.Method && tq.Path == r.RequestURI;
	})

	if(len(request) == 0) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	matched := request[0]

	if (matched.Matcher != nil) {
		matched.Matcher(r)
	}

	for key, values := range matched.Response.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(matched.Response.Status)
	fmt.Fprint(w, matched.Response.Body)
}

func Filter(vs []TestRequest, f func(tq TestRequest) bool) []TestRequest {
	vsf := make([]TestRequest, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func NewServer(requests []TestRequest) (*httptest.Server, *TestHandler) {
	handler := &TestHandler{Requests: requests}
	return httptest.NewServer(handler), handler
}