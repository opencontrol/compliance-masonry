package version

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

// IncompatibleVersionError is a type of error that contains more specific information about the version that is
// incompatible with the current masonry toolchain.
type IncompatibleVersionError struct {
	file          string
	fileType      string
	actualVersion float32
	minVersion    float32
	maxVersion    float32
}

// Error returns the string representation of the IncompatibleVersionError and satisfies the `error` interface.
func (e IncompatibleVersionError) Error() string {
	advice := ""
	if e.minVersion != constants.VersionNotNeeded {
		advice += fmt.Sprintf(" Min Version supported: %.2f", e.minVersion)
	}
	if e.maxVersion != constants.VersionNotNeeded {
		advice += fmt.Sprintf(" Max Version supported: %.2f", e.maxVersion)
	}
	return fmt.Sprintf("File: [%s] uses version %.2f. Filetype: [%s], %s",
		e.file,
		e.actualVersion,
		e.fileType,
		advice,
	)
}

// NewIncompatibleVersionError is a constructor for the IncompatibleVersionError
func NewIncompatibleVersionError(file string, fileType string, actualVersion float32,
	minVersion float32, maxVersion float32) IncompatibleVersionError {
	return IncompatibleVersionError{
		file:          file,
		fileType:      fileType,
		actualVersion: actualVersion,
		minVersion:    minVersion,
		maxVersion:    maxVersion,
	}
}
