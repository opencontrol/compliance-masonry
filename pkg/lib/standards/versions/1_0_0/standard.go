package standard

import (
	"sort"

	"github.com/opencontrol/compliance-masonry/lib/common"
	"vbom.ml/util/sortorder"
)

// Control struct stores data on a specific security requirement
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
type Control struct {
	Family      string `yaml:"family" json:"family"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
}

// Standard struct is a collection of security requirements
// Schema info: https://github.com/opencontrol/schemas#standards-documentation
type Standard struct {
	Name     string             `yaml:"name" json:"name"`
	Controls map[string]Control `yaml:",inline"`
}

// GetSortedControls returns a list of sorted controls
func (standard Standard) GetSortedControls() []string {
	var controlNames []string
	for controlName := range standard.Controls {
		controlNames = append(controlNames, controlName)
	}
	sort.Sort(sortorder.Natural(controlNames))
	return controlNames
}

// GetName returns the name of the standard.
func (standard Standard) GetName() string {
	return standard.Name
}

// GetControls returns all controls associated with the standard
func (standard Standard) GetControls() map[string]common.Control {
	m := make(map[string]common.Control)
	for key, value := range standard.Controls {
		m[key] = value
	}
	return m
}

// GetControl returns a particular control
func (standard Standard) GetControl(controlKey string) common.Control {
	return standard.Controls[controlKey]
}

// GetFamily returns which family the control belongs to.
func (control Control) GetFamily() string {
	return control.Family
}

// GetName returns the string representation of the control.
func (control Control) GetName() string {
	return control.Name
}

// GetDescription returns the string description of the control.
func (control Control) GetDescription() string {
	return control.Description
}
