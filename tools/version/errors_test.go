package version_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/opencontrol/compliance-masonry/tools/version"

	. "github.com/onsi/ginkgo"

	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

var _ = Describe("Errors", func() {

	Describe("NewIncompatibleVersionError", func() {
		It("should return a IncompatibleVersionError", func() {
			err := NewIncompatibleVersionError("file", "type", 2, 0, 1)
			assert.IsType(GinkgoT(), IncompatibleVersionError{}, err)
		})
	})

	DescribeTable("String() for IncompatibleVersionError",
	func(err IncompatibleVersionError, output string){
		assert.Equal(GinkgoT(), output, err.Error())
	},
		Entry("Min version not needed", NewIncompatibleVersionError("file", "type", 2, constants.VersionNotNeeded, 1), "File: [file] uses version 2.00. Filetype: [type],  Max Version supported: 1.00"),
		Entry("Max version not needed", NewIncompatibleVersionError("file", "type", 1, 2, constants.VersionNotNeeded), "File: [file] uses version 1.00. Filetype: [type],  Min Version supported: 2.00"),
		Entry("Both max and min versions needed", NewIncompatibleVersionError("file", "type", 4, 2, 3), "File: [file] uses version 4.00. Filetype: [type],  Min Version supported: 2.00 Max Version supported: 3.00"),

	)
})
