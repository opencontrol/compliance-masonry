package version

import "github.com/opencontrol/compliance-masonry/tools/constants"

// VerifyVersion will check if the version is compatible. If not, it will return a IncompatibleVersionError.
func VerifyVersion(file string, fileType string, version, minVersion, maxVersion float32) error {
	if minVersion == constants.VersionNotNeeded && maxVersion == constants.VersionNotNeeded {
		return nil
	} else if version >= minVersion && maxVersion == constants.VersionNotNeeded {
		return nil
	} else if minVersion == constants.VersionNotNeeded && version <= constants.VersionNotNeeded {
		return nil
	} else if version >= minVersion && version <= maxVersion {
		return nil
	}
	return NewIncompatibleVersionError(file, fileType, version, minVersion, maxVersion)
}
