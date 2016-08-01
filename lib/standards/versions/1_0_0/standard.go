package v1_0_0

import (
	"sort"
	"vbom.ml/util/sortorder"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// Control struct stores data on a specific security requirement
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
type Control struct {
	Family string `yaml:"family" json:"family"`
	Name   string `yaml:"name" json:"name"`
}

// Standard struct is a collection of security requirements
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
type Standard struct {
	Name     string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}

// GetSortedData returns a list of sorted controls
func (standard Standard) GetSortedControls() []string {
	var controlNames []string
	for controlName := range standard.Controls {
		controlNames = append(controlNames, controlName)
	}
	sort.Sort(sortorder.Natural(controlNames))
	return controlNames
}

func (standard Standard) GetName() string {
	return standard.Name
}

func (standard Standard) GetControls() map[string]common.Control {
	m := make(map[string]common.Control)
	for key, value := range standard.Controls {
		m[key] = value
	}
	return m
}

func (standard Standard) GetControl(controlKey string) common.Control {
	return standard.Controls[controlKey]
}

func (control Control) GetFamily() string {
	return control.Family
}

func (control Control) GetName() string {
	return control.Name
}
