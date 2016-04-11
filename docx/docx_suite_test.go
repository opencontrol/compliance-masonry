package docx_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDocx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docx Suite")
}
