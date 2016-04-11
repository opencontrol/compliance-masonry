package models

import (
	"io/ioutil"
	"sort"

	"gopkg.in/yaml.v2"
	"vbom.ml/util/sortorder"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]Standard `yaml:"standards" json:"standards"`
}

// GetSortedData returns a list of sorted standards
func (certification Certification) GetSortedData(callback func(string, string)) {
	var standardNames []string
	for standardName := range certification.Standards {
		standardNames = append(standardNames, standardName)
	}
	sort.Sort(sortorder.Natural(standardNames))
	for _, standardKey := range standardNames {
		certification.Standards[standardKey].GetSortedData(func(controlKey string) {
			callback(standardKey, controlKey)
		})
	}
}

// LoadCertification struct loads certifications into a Certification struct
// and add it to the main object.
func (openControl *OpenControl) LoadCertification(certificationFile string) error {
	var certification Certification
	certificationData, err := ioutil.ReadFile(certificationFile)
	if err != nil {
		return ErrReadFile
	}
	err = yaml.Unmarshal(certificationData, &certification)
	if err != nil {
		return ErrCertificationSchema
	}
	openControl.Certification = &certification
	return nil
}
