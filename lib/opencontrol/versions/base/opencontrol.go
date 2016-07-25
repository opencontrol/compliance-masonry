package base

// Base is the common struct that all schemas must have.
type Base struct {
	// SchemaVersion contains the schema version.
	SchemaVersion string `yaml:"schema_version"`
}

// GetSchemaVersion is a simple getter function of the schema version.
func (b Base) GetSchemaVersion() string {
	return b.SchemaVersion
}
