package version

import (
	. "github.com/onsi/ginkgo/extensions/table"

	"github.com/blang/semver"
	. "github.com/onsi/ginkgo"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

const (
	file     = "file"
	fileType = "fileType"
)

var (
	badReqVersionLTMin            = NewRequirements(file, fileType, semver.MustParse("0.0.0"), semver.MustParse("1.0.0"), constants.VersionNotNeeded)
	badReqVersionGTMax            = NewRequirements(file, fileType, semver.MustParse("2.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"))
	badReqVersionOutsideMaxAndMin = NewRequirements(file, fileType, semver.MustParse("2.0.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0"))

	reqMinAndMaxNotNeeded       = NewRequirements(file, fileType, semver.MustParse("0.0.0"), constants.VersionNotNeeded, constants.VersionNotNeeded)
	reqVersionEQMinMaxNotNeeded = NewRequirements(file, fileType, semver.MustParse("0.0.0"), semver.MustParse("0.0.0"), constants.VersionNotNeeded)
	reqVersionGTMinMaxNotNeeded = NewRequirements(file, fileType, semver.MustParse("1.0.0"), semver.MustParse("0.0.0"), constants.VersionNotNeeded)
	reqVersionEQMaxMinNotNeeded = NewRequirements(file, fileType, semver.MustParse("1.0.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"))
	reqVersionLTMaxMinNotNeeded = NewRequirements(file, fileType, semver.MustParse("0.5.0"), constants.VersionNotNeeded, semver.MustParse("1.0.0"))
	reqVersionLTMaxGTMin        = NewRequirements(file, fileType, semver.MustParse("0.5.0"), semver.MustParse("0.0.0"), semver.MustParse("1.0.0"))
	reqVersionEQMaxEQMin        = NewRequirements(file, fileType, semver.MustParse("1.0.0"), semver.MustParse("1.0.0"), semver.MustParse("1.0.0"))
)

var _ = Describe("Compare", func() {

	DescribeTable("VerifyVersion",
		func(req Requirements, expectedError error) {
			assert.Equal(GinkgoT(), expectedError, req.VerifyVersion())
		},
		Entry("both the max and min are not needed", reqMinAndMaxNotNeeded, nil),
		Entry("the given version is equal to the min and the max is not needed", reqVersionEQMinMaxNotNeeded, nil),
		Entry("the given version is greater than the min and the max is not needed", reqVersionGTMinMaxNotNeeded, nil),
		Entry("the given version is less than the min and the max is not needed", badReqVersionLTMin, NewIncompatibleVersionError(badReqVersionLTMin)),
		Entry("the given version is equal to the max and the min is not needed", reqVersionEQMaxMinNotNeeded, nil),
		Entry("the given version is less than the max and the min is not needed", reqVersionLTMaxMinNotNeeded, nil),
		Entry("the given version is greater than the max and the min is not needed", badReqVersionGTMax, NewIncompatibleVersionError(badReqVersionGTMax)),
		Entry("the given version is less than the max and greater than the min", reqVersionLTMaxGTMin, nil),
		Entry("the given version is equal to both the max and the min", reqVersionEQMaxEQMin, nil),
		Entry("the given version is not within the min and max", badReqVersionOutsideMaxAndMin, NewIncompatibleVersionError(badReqVersionOutsideMaxAndMin)),
	)
})
