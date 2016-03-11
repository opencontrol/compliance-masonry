package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Certification struct is a collection of specific standards and controls
type Certification struct {
	Key       string              `yaml:"name" json:"name"`
	Standards map[string]Standard `yaml:"standards" json:"standards"`
}

// LoadCertification struct loads certifications into a Certification struct
// and add it to the main object.
func (openControl *OpenControl) LoadCertification(certificationFile string) {
	var certification Certification
	certificationData, err := ioutil.ReadFile(certificationFile)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(certificationData, &certification)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Certification = &certification
}
