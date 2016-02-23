package counters_test

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/counter-demo/counters"
)

var _ = Describe("Counting lines in a single file", func() {
	var (
		fileLinesCounter *counters.FileLinesCounter
		filePath         string
		expectedLines    int
	)

	BeforeEach(func() {
		fileLinesCounter = &counters.FileLinesCounter{}

		tempFile, err := ioutil.TempFile("", "test-file-")
		Expect(err).NotTo(HaveOccurred())
		filePath = tempFile.Name()

		expectedLines = rand.Intn(100)
		for i := 0; i < expectedLines; i++ {
			_, err := tempFile.WriteString(fmt.Sprintf("this is line %d\n", i+1))
			Expect(err).NotTo(HaveOccurred())
		}
		tempFile.Close()
	})

	AfterEach(func() {
		Expect(os.RemoveAll(filePath)).To(Succeed())
	})

	It("returns the number of lines", func() {
		lines, err := fileLinesCounter.CountLines(filePath)
		Expect(err).NotTo(HaveOccurred())

		Expect(lines).To(Equal(expectedLines))
	})

	Context("when opening the file fails", func() {
		BeforeEach(func() {
			Expect(os.Remove(filePath)).To(Succeed())
		})

		It("returns the error", func() {
			_, err := fileLinesCounter.CountLines(filePath)
			Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
		})
	})
})
