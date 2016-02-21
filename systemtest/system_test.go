package system_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the webserver", func() {
	It("responds to GET / with a hello", func() {
		url := fmt.Sprintf("http://%s/", serverAddress)

		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())

		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(200))

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(bodyBytes).To(Equal([]byte("hello")))
	})
})
