package common

import "errors"

var (
	// ErrNoDataToParse represents the case that there is no data to be found to be parsed (either nil or empty).
	ErrNoDataToParse = errors.New("No data to parse")
	// ErrUnknownSchemaVersion is thrown when the schema version is unknown to the parser.
	ErrUnknownSchemaVersion = errors.New("Unknown schema version")
	// ErrCantParseSemver is thrown when the semantic versioning can not be parsed.
	ErrCantParseSemver = errors.New("Can't parse semantic versioning of schema_version")
	// ErrReadFile is raised when a file can not be read
	ErrReadFile = errors.New("Unable to read the file")
	// ErrCertificationSchema is raised a certification cannot be parsed
	ErrCertificationSchema = errors.New("Unable to parse certification")
	// ErrStandardSchema is raised a standard cannot be parsed
	ErrStandardSchema = errors.New("Unable to parse standard")
)
