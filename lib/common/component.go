package common

import "github.com/blang/semver"

type Component interface {
	GetName() string
	GetKey() string
	SetKey(string)
	GetAllSatisfies() []Satisfies
	GetVerifications() *VerificationReferences
	GetReferences() *GeneralReferences
	GetVersion() semver.Version
	SetVersion(semver.Version)
	GetResponsibleRole() string
}

type Satisfies interface {
	GetStandardKey() string
	GetControlKey() string
	GetNarratives() []Section
	GetParameters() []Section
	GetCoveredBy() CoveredByList
	GetControlOrigin() string
	GetControlOrigins() []string
	GetImplementationStatus() string
	GetImplementationStatuses() []string
}

type Section interface {
	GetKey() string
	GetText() string
}
