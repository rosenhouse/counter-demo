package handlers_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/counter-demo/handlers"
	"github.com/rosenhouse/counter-demo/mocks"
)

var _ = Describe("Lines handler", func() {
	var handler *handlers.Lines
	var mockCounter *mocks.Counter

	BeforeEach(func() {
		mockCounter = &mocks.Counter{}
		handler = &handlers.Lines{
			Counter: mockCounter,
		}
	})

	It("responds with JSON encoding the number of lines", func() {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/some/url/path", nil)
		handler.ServeHTTP(resp, req)

		Expect(resp.Body.String()).To(MatchJSON(`{ "lines": 42 }`))
	})

	It("interprets the URL path suffix as a Golang package path", func() {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/lines/some/go/package", nil)
		handler.ServeHTTP(resp, req)

		Expect(mockCounter.CountCall.Receives.Package).To(Equal("some/go/package"))
	})
})
