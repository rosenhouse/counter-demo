package systemtest

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Counting lines", func() {
	var serverSession *gexec.Session
	var address string
	var serverIsAvailable = func() error {
		_, err := net.Dial("tcp", address)
		return err
	}

	BeforeEach(func() {
		address = "localhost:8000"
		command := exec.Command(pathToServerBinary)
		var err error
		serverSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Eventually(serverIsAvailable).Should(Succeed())
	})

	AfterEach(func() {
		serverSession.Interrupt()
		Eventually(serverSession).Should(gexec.Exit())
	})

	It("returns the number of lines of Go code in a source tree", func() {
		testTree := "github.com/rosenhouse/counter-demo"
		url := fmt.Sprintf("http://%s/%s", address, testTree)

		resp, err := http.Get(url)

		Expect(err).NotTo(HaveOccurred())
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(http.StatusOK))

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		Expect(err).NotTo(HaveOccurred())

		observedCount, err := strconv.Atoi(string(bodyBytes))
		Expect(err).NotTo(HaveOccurred())

		alternateCount := alternateMethod_UNIX(testTree)
		Expect(observedCount).To(Equal(alternateCount))
	})
})

func alternateMethod_UNIX(srcTree string) int {
	path := os.Getenv("GOPATH") + "/src/" + srcTree
	alternateMethod := exec.Command("/bin/bash", "-c",
		"find "+path+" -name '*.go' | xargs cat | wc -l")
	alternateCount, err := alternateMethod.CombinedOutput()
	Expect(err).NotTo(HaveOccurred())
	count, err := strconv.Atoi(strings.TrimSpace(string(alternateCount)))
	Expect(err).NotTo(HaveOccurred())
	return count
}
