package common

import (
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
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
	ParseV1_0_0(data []byte) (BaseSchema, error)
}

// BaseSchema is an interface that every schema should implement.
type BaseSchema interface {
	Parse(data []byte) error
	GetSchemaVersion() string
	GetResources(string, string, *ConfigWorker) error
}

// ConfigWorker is a container of all COMMON things needed to do work on the configs.
type ConfigWorker struct {
	Parser      SchemaParser
	Downloader  EntryDownloader
	ResourceMap mapset.MapSet
	FSUtil      fs.Util
}
