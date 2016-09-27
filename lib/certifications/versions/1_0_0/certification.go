package certification

import (
	"sort"
	"vbom.ml/util/sortorder"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]map[string]interface{} `yaml:"standards" json:"standards"`
}

// GetKey returns the name of the certification.
func (certification Certification) GetKey() string {
	return certification.Key
}

// GetSortedStandards returns a list of sorted standard names
func (certification Certification) GetSortedStandards() []string {
	var standardNames []string
	for standardName := range certification.Standards {
		standardNames = append(standardNames, standardName)
	}
	sort.Sort(sortorder.Natural(standardNames))
	return standardNames
}

// GetStandard returns the standard for the given key.
func (certification Certification) GetControlKeysFor(standardKey string) []string {
	var controlNames []string
	for controlName := range certification.Standards[standardKey] {
		controlNames = append(controlNames, controlName)
	}
	sort.Sort(sortorder.Natural(controlNames))
	return controlNames
}