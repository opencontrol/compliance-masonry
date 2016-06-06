package version

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"

	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Errors", func() {

	Describe("NewIncompatibleVersionError", func() {
		It("should return a IncompatibleVersionError", func() {
			err := NewIncompatibleVersionError(NewRequirements("file", "type", semver.MustParse("2.0.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0")))
			assert.IsType(GinkgoT(), IncompatibleVersionError{}, err)
		})
	})

	DescribeTable("String() for IncompatibleVersionError",
		func(err IncompatibleVersionError, output string) {
			assert.Equal(GinkgoT(), output, err.Error())
		},
		Entry("Min version not needed", NewIncompatibleVersionError(NewRequirements("file", "type", semver.MustParse("2.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"))), "File: [file] uses version 2.0.0. Filetype: [type],  Max Version supported: 1.0.0"),
		Entry("Max version not needed", NewIncompatibleVersionError(NewRequirements("file", "type", semver.MustParse("1.0.0"), semver.MustParse("2.0.0"), constants.VersionNotNeeded)), "File: [file] uses version 1.0.0. Filetype: [type],  Min Version supported: 2.0.0"),
		Entry("Both max and min versions needed", NewIncompatibleVersionError(NewRequirements("file", "type", semver.MustParse("4.0.0"), semver.MustParse("2.0.0"), semver.MustParse("3.0.0"))), "File: [file] uses version 4.0.0. Filetype: [type],  Min Version supported: 2.0.0 Max Version supported: 3.0.0"),
	)
})
