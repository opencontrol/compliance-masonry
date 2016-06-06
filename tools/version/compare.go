package version

import (
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

// Requirements is a set of version requirements needed for a given file with some helper information.
type Requirements struct {
	file       string
	fileType   string
	version    semver.Version
	minVersion semver.Version
	maxVersion semver.Version
}

// NewRequirements will create a new set of version requirements for a file.
func NewRequirements(file string, fileType string, version semver.Version,
	minVersion semver.Version, maxVersion semver.Version) Requirements {
	return Requirements{file: file, fileType: fileType, version: version,
		minVersion: minVersion, maxVersion: maxVersion}
}

// VerifyVersion will check if the version is compatible. If not, it will return a IncompatibleVersionError.
func (r Requirements) VerifyVersion() error {
	if r.minVersion.EQ(constants.VersionNotNeeded) && r.maxVersion.EQ(constants.VersionNotNeeded) {
		return nil
	} else if r.version.GTE(r.minVersion) && r.maxVersion.EQ(constants.VersionNotNeeded) {
		return nil
	} else if r.minVersion.EQ(constants.VersionNotNeeded) && r.version.LTE(r.maxVersion) {
		return nil
	} else if r.version.GTE(r.minVersion) && r.version.LTE(r.maxVersion) {
		return nil
	}
	return NewIncompatibleVersionError(r)
}
