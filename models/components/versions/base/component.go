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

// Base is the common struct that all component schemas must have.
type Base struct {
	SchemaVersion float32 `yaml:"schema_version" json:"schema_version"`
}