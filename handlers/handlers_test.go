package handlers_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/counter-demo/handlers"
)

var _ = Describe("Lines handler", func() {
	It("responds with JSON encoding the number of lines", func() {
		handler := handlers.Lines{}

		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/some/url/path", nil)
		handler.ServeHTTP(resp, req)

		Expect(resp.Body.String()).To(MatchJSON(`{ "lines": 42 }`))
	})
})
