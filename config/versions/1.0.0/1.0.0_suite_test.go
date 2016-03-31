package schema_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func Test1_0_0(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "1.0.0 Suite")
}
