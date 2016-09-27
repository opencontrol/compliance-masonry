package common

// Certification is the interface for getting all the attributes for a given certification.
// Schema info: https://github.com/opencontrol/schemas#certifications
//
// GetKey returns the the unique key that represents the name of the certification.
//
// GetStandards returns the standards.
type Certification interface {
	GetKey() string
	GetSortedStandards() []string
	GetControlKeysFor(string) []string
}