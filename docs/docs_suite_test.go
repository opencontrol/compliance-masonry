package docs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDocs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docs Suite")
}
