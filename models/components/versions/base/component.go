package base

import (
	"github.com/opencontrol/compliance-masonry/models/common"
)

type Component interface {
	GetName() string
	GetKey() string
	SetKey(string)
	GetAllSatisfies() []Satisfies
	GetVerifications() *common.VerificationReferences
	GetReferences() *common.GeneralReferences
	GetVersion() float32
}

type Satisfies interface {
	GetStandardKey() string
	GetControlKey() string
	GetNarrative() string
	GetCoveredBy() common.CoveredByList
}

// Base is the bare minimum that every component YAML will have and is used to find the schema version.
// Complete implementations of component do not need to embed this struct or put it as a field in the component.
type Base struct {
	SchemaVersion float32 `yaml:"schema_version" json:"schema_version"`
}