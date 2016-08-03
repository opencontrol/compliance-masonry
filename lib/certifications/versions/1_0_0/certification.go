package certification

import (
	"sort"
	"vbom.ml/util/sortorder"
	v1standards "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]v1standards.Standard `yaml:"standards" json:"standards"`
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

// GetStandards returns a map of all the standard names and their corresponding standard.
func (certification Certification) GetStandards() map[string]common.Standard {
	m := make(map[string]common.Standard)
	for key, value := range certification.Standards {
		m[key] = value
	}
	return m
}