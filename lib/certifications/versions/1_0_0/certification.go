package certification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"vbom.ml/util/sortorder"
)

// Certification struct is a collection of specific standards and controls
// Schema info: https://github.com/opencontrol/schemas#certifications
type Certification struct {
	//	Key       string                            `yaml:"name" json:"name"`
	//	Standards map[string]map[string]interface{} `yaml:"standards" json:"standards"`
	Key       string                            `json:"name" yaml:"name"`
	Standards map[string]map[string]interface{} `json:"standards" yaml:"standards"`
}

// MarshalJSON provides JSON support
func (p *Certification) MarshalJSON() (b []byte, e error) {
	// start the marshaling
	buffer := bytes.NewBufferString("{")

	// write data
	buffer.WriteString(fmt.Sprintf("\"key\":\"%s\",\"standards\":", p.Key))
	bytes, err := json.Marshal(p.GetSortedStandards())
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytes))

	// done with marshaling
	buffer.WriteString("}")
	return buffer.Bytes(), nil
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

// GetControlKeysFor returns the control keys for the given standard key.
func (certification Certification) GetControlKeysFor(standardKey string) []string {
	var controlNames []string
	for controlName := range certification.Standards[standardKey] {
		controlNames = append(controlNames, controlName)
	}
	sort.Sort(sortorder.Natural(controlNames))
	return controlNames
}
