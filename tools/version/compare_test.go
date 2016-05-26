package version_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/opencontrol/compliance-masonry/tools/version"

	. "github.com/onsi/ginkgo"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Compare", func() {

	DescribeTable("VerifyVersion",
		func(version, min, max float32, expectedError error) {
			assert.Equal(GinkgoT(), expectedError, VerifyVersion("file", "type", version, min, max))
		},
		Entry("both the max and min are not needed", float32(0.0), constants.VersionNotNeeded, constants.VersionNotNeeded, nil),
		Entry("the given version is equal to the min and the max is not needed", float32(0.0), float32(0.0), constants.VersionNotNeeded, nil),
		Entry("the given version is greater than the min and the max is not needed", float32(1.0), float32(0.0), constants.VersionNotNeeded, nil),
		Entry("the given version is less than the min and the max is not needed", float32(0.0), float32(1.0), constants.VersionNotNeeded, NewIncompatibleVersionError("file", "type", 0, 1, constants.VersionNotNeeded)),
		Entry("the given version is equal to the max and the min is not needed", float32(1.0), constants.VersionNotNeeded, float32(1.0), nil),
		Entry("the given version is less than the max and the min is not needed", float32(0.5), constants.VersionNotNeeded, float32(1.0), nil),
		Entry("the given version is greater than the max and the min is not needed", float32(2.0), constants.VersionNotNeeded, float32(1.0), NewIncompatibleVersionError("file", "type", 2, constants.VersionNotNeeded, 1)),
		Entry("the given version is less than the max and greater than the min", float32(0.5), float32(0.0), float32(1.0), nil),
		Entry("the given version is equal to both the max and the min", float32(1.0), float32(1.0), float32(1.0), nil),
		Entry("the given version is not within the min and max", float32(2.0), float32(0.0), float32(1.0), NewIncompatibleVersionError("file", "type", 2, 0, 1)),
	)
})
