package v1_0_0

import (
	"sort"
	"vbom.ml/util/sortorder"
	v1standards "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/lib/standards"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]v1standards.Standard `yaml:"standards" json:"standards"`
}

func (certification Certification) GetKey() string {
	return certification.Key
}

// GetSortedStandards returns a list of sorted standards
func (certification Certification) GetSortedStandards() []string {
	var standardNames []string
	for standardName := range certification.Standards {
		standardNames = append(standardNames, standardName)
	}
	sort.Sort(sortorder.Natural(standardNames))
	return standardNames
}

func (certification Certification) GetStandards() map[string]standards.Standard {
	m := make(map[string]standards.Standard)
	for key, value := range certification.Standards {
		m[key] = value
	}
	return m
}