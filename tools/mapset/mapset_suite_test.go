package mapset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMapset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapset Suite")
}
