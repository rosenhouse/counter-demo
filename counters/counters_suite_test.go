package counters_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCounters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Counters Suite")
}
