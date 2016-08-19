package lib

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
	"vbom.ml/util/sortorder"
	v1 "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]v1.Standard `yaml:"standards" json:"standards"`
}

// GetSortedData returns a list of sorted standards
func (certification Certification) GetSortedData(callback func(string, string)) {
	var standardNames []string
	for standardName := range certification.Standards {
		standardNames = append(standardNames, standardName)
	}
	sort.Sort(sortorder.Natural(standardNames))
	for _, standardKey := range standardNames {
		controlKeys := certification.Standards[standardKey].GetSortedControls()
		for _, controlKey := range controlKeys {
			callback(standardKey, controlKey)
		}
	}
}

// LoadCertification struct loads certifications into a Certification struct
// and add it to the main object.
func (ws *LocalWorkspace) LoadCertification(certificationFile string) error {
	var certification Certification
	certificationData, err := ioutil.ReadFile(certificationFile)
	if err != nil {
		return common.ErrReadFile
	}
	err = yaml.Unmarshal(certificationData, &certification)
	if err != nil {
		return common.ErrCertificationSchema
	}
	ws.Certification = &certification
	return nil
}
