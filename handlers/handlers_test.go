package handlers_test

import (
	"errors"
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

		mockCounter.CountCall.Returns.Lines = 1234
	})

	It("responds with JSON encoding the number of lines", func() {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/some/url/path", nil)
		handler.ServeHTTP(resp, req)

		Expect(resp.Body.String()).To(MatchJSON(`{ "lines": 1234 }`))
	})

	It("interprets the URL path suffix as a Golang package path", func() {
		resp := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/lines/some/go/package", nil)
		handler.ServeHTTP(resp, req)

		Expect(mockCounter.CountCall.Receives.Package).To(Equal("some/go/package"))
	})

	Context("when the counter fails", func() {
		BeforeEach(func() {
			mockCounter.CountCall.Returns.Error = errors.New("some error")
		})

		It("responds with status code 500 and the error message in JSON", func() {
			resp := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/some/url/path", nil)
			handler.ServeHTTP(resp, req)

			Expect(resp.Code).To(Equal(500))
			Expect(resp.Body.String()).To(MatchJSON(`{ "error": "some error" }`))
		})
	})
})
