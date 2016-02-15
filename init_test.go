package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSystemTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Test Suite")
}
