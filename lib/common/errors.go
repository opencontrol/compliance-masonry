package common

import "errors"

var (
	// ErrNoDataToParse represents the case that there is no data to be found to be parsed (either nil or empty).
	ErrNoDataToParse = errors.New("No data to parse")
	// ErrUnknownSchemaVersion is thrown when the schema version is unknown to the parser.
	ErrUnknownSchemaVersion = errors.New("Unknown schema version")
	// ErrCantParseSemver is thrown when the semantic versioning can not be parsed.
	ErrCantParseSemver = errors.New("Can't parse semantic versioning of schema_version")
)

