package main_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/rosenhouse/counter-demo"
	"github.com/rosenhouse/counter-demo/mocks"
)

var _ = Describe("CountHandler", func() {
	var (
		resp    *httptest.ResponseRecorder
		handler *main.CountHandler
		counter *mocks.Counter
	)

	BeforeEach(func() {
		resp = httptest.NewRecorder()
		counter = &mocks.Counter{}
		handler = &main.CountHandler{
			Counter: counter,
		}
	})

	It("passes the URL path into the Counter", func() {
		url := "http://example.com/some/go/package"
		req, _ := http.NewRequest("GET", url, nil)

		handler.ServeHTTP(resp, req)

		Expect(counter.CountLinesCall.CallCount).To(Equal(1))

		Expect(
			counter.CountLinesCall.Receives.PackageName,
		).To(Equal("some/go/package"))
	})

	It("returns the line count as a string", func() {
		counter.CountLinesCall.Returns.Count = 42

		req, _ := http.NewRequest("GET", "some-url", nil)

		handler.ServeHTTP(resp, req)

		Expect(resp.Body.String()).To(Equal("42"))
	})
})
