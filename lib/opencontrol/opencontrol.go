package opencontrol

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Base is the common struct that all schemas must have.
type Base struct {
	// SchemaVersion contains the schema version.
	SchemaVersion string `yaml:"schema_version"`
}

// GetSchemaVersion is a simple getter function of the schema version.
func (b Base) GetSchemaVersion() string {
	return b.SchemaVersion
}

// SchemaParser is a generic interface that knows how parse different schema_versions.
type SchemaParser interface {
	Parse(data []byte) (common.OpenControl, error)
}
