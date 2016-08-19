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

// GetStandard returns the standard for the given key.
func (certification Certification) GetStandard(key string) common.Standard {
	return certification.Standards[key]
}