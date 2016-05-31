package version

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

// IncompatibleVersionError is a type of error that contains more specific information about the version that is
// incompatible with the current masonry toolchain.
type IncompatibleVersionError struct {
	file          string
	fileType      string
	actualVersion semver.Version
	minVersion    semver.Version
	maxVersion    semver.Version
}

// Error returns the string representation of the IncompatibleVersionError and satisfies the `error` interface.
func (e IncompatibleVersionError) Error() string {
	advice := ""
	if e.minVersion.NE(constants.VersionNotNeeded) {
		advice += fmt.Sprintf(" Min Version supported: %s", e.minVersion.String())
	}
	if e.maxVersion.NE(constants.VersionNotNeeded) {
		advice += fmt.Sprintf(" Max Version supported: %s", e.maxVersion.String())
	}
	return fmt.Sprintf("File: [%s] uses version %s. Filetype: [%s], %s",
		e.file,
		e.actualVersion.String(),
		e.fileType,
		advice,
	)
}

// NewIncompatibleVersionError is a constructor for the IncompatibleVersionError
func NewIncompatibleVersionError(file string, fileType string, actualVersion,
	minVersion, maxVersion semver.Version) IncompatibleVersionError {
	return IncompatibleVersionError{
		file:          file,
		fileType:      fileType,
		actualVersion: actualVersion,
		minVersion:    minVersion,
		maxVersion:    maxVersion,
	}
}
