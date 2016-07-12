package inventory_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInventory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inventory Suite")
}
