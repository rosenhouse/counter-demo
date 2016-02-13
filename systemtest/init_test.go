package systemtest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestSystemTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "System Test Suite")
}

var pathToServerBinary string

var _ = BeforeSuite(func() {
	var err error
	pathToServerBinary, err = gexec.Build("github.com/rosenhouse/counter-demo")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
