package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Control struct stores data on a specific security requirement
type Control struct {
	Family string `yaml:"family" json:"family"`
	Name   string `yaml:"name" json:"name"`
}

// Standard struct is a collection of security requirements
type Standard struct {
	Key      string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}

// LoadStandard imports a standard into the Standard struct and adds it to the
// main object.
func (openControl *OpenControl) LoadStandard(standardFile string) {
	var standard Standard
	standardData, err := ioutil.ReadFile(standardFile)
	if err != nil {
		log.Println(err.Error())
	}
	err = yaml.Unmarshal(standardData, &standard)
	if err != nil {
		log.Println(err.Error())
	}
	openControl.Standards[standard.Key] = &standard
}
