package system_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("the webserver", func() {
	It("responds to GET /lines/:pkgPath with the line count", func() {
		pkgPath := "github.com/rosenhouse/counter-demo"
		url := fmt.Sprintf("http://%s/lines/%s", serverAddress, pkgPath)

		resp, err := http.Get(url)
		Expect(err).NotTo(HaveOccurred())

		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(200))

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())

		var result struct {
			Lines int `json:"lines"`
		}
		Expect(json.Unmarshal(bodyBytes, &result)).To(Succeed())

		expectedLineCount := unixLineCount(pkgPath)
		Expect(result.Lines).To(Equal(expectedLineCount))
	})
})

func unixLineCount(pkgPath string) int {
	cmd := exec.Command("/bin/sh", "-c",
		"find . -name '*.go' | xargs wc -l | tail -n1 | awk '{ print $1 }'")
	cmd.Dir = filepath.Join(os.Getenv("GOPATH"), "src", pkgPath)

	outputBytes, err := cmd.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())

	outputInt, err := strconv.Atoi(strings.TrimSpace(string(outputBytes)))
	Expect(err).NotTo(HaveOccurred())

	return outputInt
}
