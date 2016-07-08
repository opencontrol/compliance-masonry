package mapset

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Map", func() {
	var (
		m = Init()
	)
	DescribeTable("Reserve", func(key, value string, expectedError error, expectedSuccess bool, expectedValue string) {
		res := m.Reserve(key, value)
		assert.Equal(GinkgoT(), expectedError, res.Error)
		assert.Equal(GinkgoT(), expectedSuccess, res.Success)
		assert.Equal(GinkgoT(), expectedValue, res.Value)
	},
		Entry("Regular reservation", "key1", "value", nil, true, "value"),
		Entry("Repeat reservation", "key1", "value", nil, false, "value"),
		Entry("no key", "", "value", ErrEmptyInput, false, ""),
		Entry("no value", "key1", "", ErrEmptyInput, false, ""),
	)

})
