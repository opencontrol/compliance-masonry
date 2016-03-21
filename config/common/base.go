package common

// Base is the common struct that all schemas must have.
type Base struct {
	// SchemaVersion contains the schema version.
	SchemaVersion string `yaml:"schema_version"`
}

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
	GetResources(string, ConfigWorker) error
}

type ConfigWorker struct {
	Parser SchemaParser
	Downloader EntryDownloader
}
