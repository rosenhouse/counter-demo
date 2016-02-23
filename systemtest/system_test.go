package system_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the webserver", func() {
	It("responds to GET /lines/:pkgPath with the line count", func() {
		pkgPath := "github.com/golang/protobuf"
		url := fmt.Sprintf("http://%s/lines/%s", serverAddress, pkgPath)

		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())

		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(200))

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())
		Expect(bodyBytes).To(MatchJSON(`{ "lines": 26071 }`))
	})
})
