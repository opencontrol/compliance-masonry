package version

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

// IncompatibleVersionError is a type of error that contains more specific information about the version that is
// incompatible with the current masonry toolchain.
type IncompatibleVersionError struct {
	req Requirements
}

// Error returns the string representation of the IncompatibleVersionError and satisfies the `error` interface.
func (e IncompatibleVersionError) Error() string {
	advice := ""
	if e.req.minVersion.NE(constants.VersionNotNeeded) {
		advice += fmt.Sprintf(" Min Version supported: %s", e.req.minVersion.String())
	}
	if e.req.maxVersion.NE(constants.VersionNotNeeded) {
		advice += fmt.Sprintf(" Max Version supported: %s", e.req.maxVersion.String())
	}
	return fmt.Sprintf("File: [%s] uses version %s. Filetype: [%s], %s",
		e.req.file,
		e.req.version.String(),
		e.req.fileType,
		advice,
	)
}

// NewIncompatibleVersionError is a constructor for the IncompatibleVersionError
func NewIncompatibleVersionError(r Requirements) IncompatibleVersionError {
	return IncompatibleVersionError{
		req: r,
	}
}
