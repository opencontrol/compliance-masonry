package base

import (
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"github.com/opencontrol/compliance-masonry/tools/mapset"
)

// SchemaParser is a generic interface that knows how parse different schema_versions.
type SchemaParser interface {
	ParseV1_0_0(data []byte) (OpenControl, error)
}

// BaseSchema is an interface that every schema should implement.
type OpenControl interface {
	Parse(data []byte) error
	GetSchemaVersion() string
	GetResources(string, string, *Worker) error
}

// Worker is a container of all COMMON things needed to do work on the configs.
type Worker struct {
	Parser      SchemaParser
	ResourceMap mapset.MapSet
	FSUtil      fs.Util
}
