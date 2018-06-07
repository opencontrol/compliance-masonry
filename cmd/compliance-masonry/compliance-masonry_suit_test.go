package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestComplianceMasonryGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ComplianceMasonryGo Suite")
}
