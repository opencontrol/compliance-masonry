package schema_tools

import "fmt"

type IncompatibleSchemaError struct {
	file                string
	fileType            string
	actualSchemaVersion float32
	minSchemaVersion    float32
	maxSchemaVersion    float32
}

func (e IncompatibleSchemaError) Error() string {
	advice := ""
	if e.minSchemaVersion != SchemaVersionNotNeeded {
		advice += fmt.Sprintf(" Min Version supported: %.2f", e.minSchemaVersion)
	}
	if e.maxSchemaVersion != SchemaVersionNotNeeded {
		advice += fmt.Sprintf(" Max Version supported: %.2f", e.maxSchemaVersion)
	}
	return fmt.Sprintf("File: [%s] uses schema version %.2f. Filetype: [%s], %s",
		e.file,
		e.actualSchemaVersion,
		e.fileType,
		advice,
	)
}

func NewIncompatibleSchemaError(file string, fileType string, actualSchemaVersion float32,
	minSchemaVersion float32, maxSchemaVersion float32) IncompatibleSchemaError {
	return IncompatibleSchemaError{
		file:                file,
		fileType:            fileType,
		actualSchemaVersion: actualSchemaVersion,
		minSchemaVersion:    minSchemaVersion,
		maxSchemaVersion:    maxSchemaVersion,
	}
}
