package version

import (
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

// VerifyVersion will check if the version is compatible. If not, it will return a IncompatibleVersionError.
func VerifyVersion(file string, fileType string, version, minVersion, maxVersion semver.Version) error {
	if minVersion.EQ(constants.VersionNotNeeded) && maxVersion.EQ(constants.VersionNotNeeded) {
		return nil
	} else if version.GTE(minVersion) && maxVersion.EQ(constants.VersionNotNeeded) {
		return nil
	} else if minVersion.EQ(constants.VersionNotNeeded) && version.LTE(maxVersion) {
		return nil
	} else if version.GTE(minVersion) && version.LTE(maxVersion) {
		return nil
	}
	return NewIncompatibleVersionError(file, fileType, version, minVersion, maxVersion)
}
