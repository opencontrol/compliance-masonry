package lib

import (
	"io/ioutil"
	"sort"
	"sync"

	"gopkg.in/yaml.v2"
	"vbom.ml/util/sortorder"
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

// Standards struct is a thread save mapping of Standards
type Standards struct {
	mapping map[string]*Standard
	sync.RWMutex
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

func (standard Standard) GetControl(controlKey string) Control {
	return standard.Controls[controlKey]
}

// NewStandards creates an instance of Components struct
func NewStandards() *Standards {
	return &Standards{mapping: make(map[string]*Standard)}
}

// Add adds a standard to the standards mapping
func (standards *Standards) Add(standard *Standard) {
	standards.Lock()
	standards.mapping[standard.Name] = standard
	standards.Unlock()
}

// Get retrieves a standard
func (standards *Standards) Get(standardName string) *Standard {
	standards.Lock()
	defer standards.Unlock()
	return standards.mapping[standardName]
}

// GetAll retrieves all the standards
func (standards *Standards) GetAll() []*Standard {
	standards.RLock()
	defer standards.RUnlock()
	s := make([]*Standard, len(standards.mapping))
	idx := 0
	for _, value := range standards.mapping {
		s[idx] = value
		idx++
	}
	return s
}

// LoadStandard imports a standard into the Standard struct and adds it to the
// main object.
func (ws *LocalWorkspace) LoadStandard(standardFile string) error {
	var standard Standard
	standardData, err := ioutil.ReadFile(standardFile)
	if err != nil {
		return ErrReadFile
	}
	err = yaml.Unmarshal(standardData, &standard)
	if err != nil {
		return ErrStandardSchema
	}
	ws.standards.Add(&standard)
	return nil
}
