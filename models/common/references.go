package common

// GeneralReference struct contains data for the name and path of a
// compliance reference.
// This struct is a one-to-one mapping of `references` in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type GeneralReference struct {
	Name string `yaml:"name" json:"name"`
	Path string `yaml:"path" json:"path"`
	Type string `yaml:"type" json:"type"`
}

//GeneralReferences a slice of type GeneralReference
type GeneralReferences []GeneralReference

// VerificationReference struct is a general reference that verifies a specific
// control, it can be pointed to in the control documentation.
// This struct is a one-to-one mapping of `verifications` in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type VerificationReference struct {
	GeneralReference `yaml:",inline"`
	Key              string `yaml:"key" json:"key"`
}

//VerificationReferences a slice of type VerificationReference
type VerificationReferences []VerificationReference

// CoveredBy struct is the pointing mechanism for for referring to
// VerificationReferences in the documentation.
// This struct is a one-to-one mapping of `covered_by` in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type CoveredBy struct {
	ComponentKey    string `yaml:"component_key" json:"component_key"`
	VerificationKey string `yaml:"verification_key" json:"verification_key"`
}

//CoveredByList a slice of type CoveredBy
type CoveredByList []CoveredBy

// Len returns the length of the GeneralReferences slice
func (slice GeneralReferences) Len() int {
	return len(slice)
}

// Less returns true if a GeneralReference is less than another reference
func (slice GeneralReferences) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

// Swap swaps the two GeneralReferences
func (slice GeneralReferences) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Len returns the length of the VerificationReference slice
func (slice VerificationReferences) Len() int {
	return len(slice)
}

// Less returns true if a VerificationReference is less than another reference
func (slice VerificationReferences) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name
}

// Swap swaps the two VerificationReferences
func (slice VerificationReferences) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// Get returns a VerificationReference of the given key
func (slice VerificationReferences) Get(key string) VerificationReference {
	for _, reference := range slice {
		if reference.Key == key {
			return reference
		}
	}
	return VerificationReference{}
}
