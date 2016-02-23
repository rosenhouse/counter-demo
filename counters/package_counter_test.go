package counters_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/counter-demo/counters"
	"github.com/rosenhouse/counter-demo/mocks"
)

var _ = Describe("Counting lines in a package", func() {
	var (
		mockDirectoryLister  *mocks.DirectoryLister
		mockFileLinesCounter *mocks.FileLinesCounter
		counter              *counters.PackageLinesCounter
	)

	BeforeEach(func() {
		mockDirectoryLister = &mocks.DirectoryLister{}
		mockFileLinesCounter = &mocks.FileLinesCounter{}
		counter = &counters.PackageLinesCounter{
			GoPath:           "/some/gopath",
			DirectoryLister:  mockDirectoryLister,
			FileLinesCounter: mockFileLinesCounter,
		}

		mockDirectoryLister.ListFilesCall.Returns.Files =
			[]string{"file-1", "file-2", "file-3"}

		mockFileLinesCounter.CountLinesStub = func(filePath string) (int, error) {
			return 10 * mockFileLinesCounter.CountLinesCallCount(), nil
		}
	})

	It("returns the total line count", func() {
		total, err := counter.Count("some/package/path")
		Expect(err).NotTo(HaveOccurred())

		Expect(total).To(Equal(10 + 20 + 30))
	})

	It("expects the package to exist in the Go source root", func() {
		_, err := counter.Count("some/package/path")
		Expect(err).NotTo(HaveOccurred())

		Expect(mockDirectoryLister.ListFilesCall.Receives.DirPath).To(Equal(
			"/some/gopath/src/some/package/path"))
	})

	Context("when the package path is not clean", func() {
		It("returns an error", func() {
			_, err := counter.Count("../../../etc/passwd")
			Expect(err).To(MatchError("unclean path"))

			_, err = counter.Count("/etc/passwd")
			Expect(err).To(MatchError("unclean path"))
		})
	})

	Context("when listing files in the package fails", func() {
		BeforeEach(func() {
			mockDirectoryLister.ListFilesCall.Returns.Error = errors.New("boom!")
		})

		It("returns the error", func() {
			_, err := counter.Count("some/package/path")
			Expect(err).To(MatchError("listing files in package: boom!"))
		})
	})

	Context("when counting lines in a file fails", func() {
		BeforeEach(func() {
			mockFileLinesCounter.CountLinesStub = func(filePath string) (int, error) {
				if mockFileLinesCounter.CountLinesCallCount() == 2 {
					return 10, errors.New("banana")
				}
				return 20, nil
			}
		})

		It("returns the error", func() {
			_, err := counter.Count("some/package/path")
			Expect(err).To(MatchError(`counting lines in "file-2": banana`))
		})
	})
})
