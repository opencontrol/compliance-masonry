package version_test

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/opencontrol/compliance-masonry/tools/version"

	"github.com/blang/semver"
	. "github.com/onsi/ginkgo"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Compare", func() {

	DescribeTable("VerifyVersion",
		func(version, min, max semver.Version, expectedError error) {
			assert.Equal(GinkgoT(), expectedError, VerifyVersion("file", "type", version, min, max))
		},
		Entry("both the max and min are not needed", semver.MustParse("0.0.0"), constants.VersionNotNeeded, constants.VersionNotNeeded, nil),
		Entry("the given version is equal to the min and the max is not needed", semver.MustParse("0.0.0"), semver.MustParse("0.0.0"), constants.VersionNotNeeded, nil),
		Entry("the given version is greater than the min and the max is not needed", semver.MustParse("1.0.0"), semver.MustParse("0.0.0"), constants.VersionNotNeeded, nil),
		Entry("the given version is less than the min and the max is not needed", semver.MustParse("0.0.0"), semver.MustParse("1.0.0"), constants.VersionNotNeeded, NewIncompatibleVersionError("file", "type", semver.MustParse("0.0.0"), semver.MustParse("1.0.0"), constants.VersionNotNeeded)),
		Entry("the given version is equal to the max and the min is not needed", semver.MustParse("1.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"), nil),
		Entry("the given version is less than the max and the min is not needed", semver.MustParse("0.5.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"), nil),
		Entry("the given version is greater than the max and the min is not needed", semver.MustParse("2.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"), NewIncompatibleVersionError("file", "type", semver.MustParse("2.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"))),
		Entry("the given version is less than the max and greater than the min", semver.MustParse("0.5.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0"), nil),
		Entry("the given version is equal to both the max and the min", semver.MustParse("1.0.0"), semver.MustParse("1.0.0"), semver.MustParse("1.0.0"), nil),
		Entry("the given version is not within the min and max", semver.MustParse("2.0.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0"), NewIncompatibleVersionError("file", "type", semver.MustParse("2.0.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0"))),
	)
})
