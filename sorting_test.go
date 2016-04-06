package main_test

import (
	//"fmt"
	//"strings"
	"sort"
	"vbom.ml/util/sortorder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Natural Sort", func() {

	// Short example from http://blog.codinghorror.com/sorting-for-humans-natural-sort-order/
	list := []string{
		"AC-1", "AC-12", "AC-2 (1)", "AC-2 (11)", "AC-3 (1)",
		"z1", "z10", "z100", "z1000", "z101", "z11", "z2", "z3",
	}

	It("Should start with a computer-sorted example", func() {
		// copy the list from above and call golang's `sort`
		default_sorted := make([]string, len(list))
		copy(default_sorted, list)
		sort.Strings(default_sorted)
		Expect(list).To(Equal(default_sorted))
	})

	It("Should sort, Natch.", func() {
		sort.Sort(sortorder.Natural(list))
		Expect(list).To(Equal([]string{
			"AC-1", "AC-2 (1)", "AC-2 (11)", "AC-3 (1)", "AC-12",
			"z1", "z2", "z3", "z10", "z11", "z100", "z101", "z1000",
		}))
	})
})
