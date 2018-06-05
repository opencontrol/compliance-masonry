package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExampleplugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exampleplugin Suite")
}
