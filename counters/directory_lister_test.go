package counters_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rosenhouse/counter-demo/counters"
)

var _ = Describe("DirectoryLister", func() {
	var (
		expectedFiles []string
		pkgRootDir    string
		lister        *counters.DirectoryLister
	)

	BeforeEach(func() {
		var err error
		pkgRootDir, err = ioutil.TempDir("", "pkgRootDir-")
		Expect(err).NotTo(HaveOccurred())

		files := []string{
			"root_file.go",
			"another_root_file.go",
			"subdir/some_sub_file.go",
			"subdir/not_a_go_file.txt",
			"subdir/another-dir/foo.go",
			"subdir2/bar.go",
			"devious_directory.go/some-non-go-file.txt",
		}

		for _, f := range files {
			fullPath := filepath.Join(pkgRootDir, f)
			Expect(os.MkdirAll(filepath.Dir(fullPath), 0777)).To(Succeed())
			Expect(ioutil.WriteFile(fullPath, []byte{}, 0666)).To(Succeed())
			if strings.HasSuffix(f, ".go") {
				expectedFiles = append(expectedFiles, fullPath)
			}
		}

		lister = &counters.DirectoryLister{}
	})

	AfterEach(func() {
		Expect(os.RemoveAll(pkgRootDir)).To(Succeed())
	})

	It("returns all Go source files in the directory and its subdirectories", func() {
		allFiles, err := lister.ListFiles(pkgRootDir)
		Expect(err).NotTo(HaveOccurred())

		Expect(allFiles).To(ConsistOf(expectedFiles))
	})

	Context("when walking the filesystem returns an error", func() {
		BeforeEach(func() {
			Expect(os.RemoveAll(pkgRootDir)).To(Succeed())
		})

		It("returns the error", func() {
			_, err := lister.ListFiles(pkgRootDir)
			Expect(err).To(MatchError(ContainSubstring("no such file or directory")))
		})
	})
})
