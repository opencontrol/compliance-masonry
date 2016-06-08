package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/version"
	"gopkg.in/yaml.v2"
)

// Components struct is a thread-safe structure mapping for components
type Components struct {
	mapping map[string]*Component
	sync.RWMutex
}

// Component struct is an individual component requiring documentation
// Schema info: https://github.com/opencontrol/schemas#component-yaml
type Component struct {
	Name          string                  `yaml:"name" json:"name"`
	Key           string                  `yaml:"key" json:"key"`
	References    *GeneralReferences      `yaml:"references" json:"references"`
	Verifications *VerificationReferences `yaml:"verifications" json:"verifications"`
	Satisfies     *SatisfiesList          `yaml:"satisfies" json:"satisfies"`
	SchemaVersion semver.Version          `yaml:"-" json:"-"`
}

type componentLoadError struct {
	message string
}

// Error implements the error interface by simply returning the message as a string.
func (e componentLoadError) Error() string {
	return e.message
}

// UnmarshalYAML is a overridden implementation of YAML parsing the component.yaml
// This method is similar to the one found here: http://choly.ca/post/go-json-marshalling/
// This is necessary because we want to have backwards compatibility with parsing the old types of version 2.0
// (type =float).
// To compensate for that, we have to hand roll our own UnmarshalYAML that can decide what to do for parsing
// the older version of type float and converting it into semver. In addition, we will use this logic to parse strings
// into semver.
func (c *Component) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// When we call "unmarshal" callback on an object, it will call that object's "UnmarshalYAML" if defined.
	// Since we are currently in the implementation of Component's "UnmarshalYAML", when finally we call
	// unmarshal again, if it's on type Component, we would end up in a recursive infinite loop.
	// To prevent this, we create a separate type, called Alias.
	type Alias Component
	// Create an anonymous struct with an interface{} type for the schema_version that we want to parse
	aux := &struct {
		SchemaVersion interface{} `yaml:"schema_version" json:"schema_version"`
		Alias         `yaml:",inline"`
	}{
		Alias: (Alias)(*c),
	}

	// Call unmarshal on the new Alias type. Don't return the error yet because we want to gather more information
	// if we can below.
	err := unmarshal(&aux)

	// Create a placeholder variable for the converted semver.
	var ver semver.Version
	// Create a placeholder variable for the error.
	var versionErr error

	// Store the version value for conciseness.
	value := aux.SchemaVersion

	// Try to cast the value from interface{} to certain types.
	switch v := value.(type) {
	// For float types, which are the old types, we need to upcast it to semver if it's an older version.
	case float32, float64:
		switch v {
		// Schema Version started being documented with "2.0".
		// We should be able to parse it for backwards compatibility.
		// All future versioning should be in semver format already.
		case 2.0:
			ver = semver.MustParse("2.0.0")
		// If not the older version, it needs to be in semver format, send an error.
		default:
			return componentLoadError{fmt.Sprintf(constants.ErrVersionNotInSemverFormatf, v)}

		}
	// The interface type will default to string if not numeric which is what all semver types will be initially.
	case string:
		ver, versionErr = semver.Parse(v)
		if versionErr != nil {
			return componentLoadError{fmt.Sprintf(constants.ErrMissingVersion)}
		}
	// In the case, it's just missing completely.
	default:
		return componentLoadError{fmt.Sprintf(constants.ErrMissingVersion)}
	}
	// Copy everything from the Alias back to the original component.
	*c = (Component)(aux.Alias)

	// Get the version
	c.SchemaVersion = ver
	return err
}

// VerifySchemaCompatibility will check that the current component schema version is
// compatible with the current masonry toolchain.
func (c *Component) VerifySchemaCompatibility(fileName string) error {
	if c != nil {
		requirements := version.NewRequirements(fileName, "component", c.SchemaVersion,
			constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion)
		return requirements.VerifyVersion()
	}
	return nil
}

// SatisfiesList is a list of Satisfies
type SatisfiesList []Satisfies

// Satisfies struct contains data demonstrating why a specific component meets
// a control
// This struct is a one-to-one mapping of a `satisfies` item in the component.yaml schema
// https://github.com/opencontrol/schemas#component-yaml
type Satisfies struct {
	ControlKey  string             `yaml:"control_key" json:"control_key"`
	StandardKey string             `yaml:"standard_key" json:"standard_key"`
	Narrative   []NarrativeSection `yaml:"narrative" json:"narrative"`
	CoveredBy   CoveredByList      `yaml:"covered_by" json:"covered_by"`
}

// NarrativeSection contains the key and text for a particular narrative section.
type NarrativeSection struct {
	Key  string `yaml:"key,omitempty" json:"key,omitempty"`
	Text string `yaml:"text" json:"text"`
}

// NewComponents creates an instance of Components struct
func NewComponents() *Components {
	return &Components{mapping: make(map[string]*Component)}
}

// Add adds a new component to the component map
func (components *Components) Add(component *Component) {
	components.Lock()
	components.mapping[component.Key] = component
	components.Unlock()
}

// Get retrieves a new component from the component map
func (components *Components) Get(key string) *Component {
	components.Lock()
	defer components.Unlock()
	return components.mapping[key]
}

// GetAndApply get a component and apply the callback function inside while locking
// components
func (components *Components) GetAndApply(key string, callback func(component *Component)) {
	components.Lock()
	callback(components.mapping[key])
	components.Unlock()
}

// CompareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *Components) CompareAndAdd(component *Component) bool {
	components.Lock()
	defer components.Unlock()
	added := false
	if _, exists := components.mapping[component.Key]; !exists {
		components.mapping[component.Key] = component
		added = true
	} else {
		log.Fatalf("Component: %s exisits!\n", component.Key)
	}
	return added
}

// GetAll retrieves all the components
func (components *Components) GetAll() map[string]*Component {
	return components.mapping
}

// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (openControl *OpenControl) LoadComponent(componentDir string) error {
	fileName := filepath.Join(componentDir, "component.yaml")
	_, err := os.Stat(fileName)
	if err != nil {
		return ErrComponentFileDNE
	}
	var component *Component
	componentData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return ErrReadFile
	}
	err = yaml.Unmarshal(componentData, &component)
	// If we have a user friendly error via componentLoadError return it.
	if err != nil {
		switch errValue := err.(type) {
		// If we a user friendly error, let's return it now.
		case componentLoadError:
			return errValue
		}
	}
	// If we don't have a user friendly error yet...
	// Check the component version to give a better error before the generic "ErrControlSchema"
	if versionErr := component.VerifySchemaCompatibility(fileName); versionErr != nil {
		return versionErr
	}

	// If no specific errors were found, but a general error was found, return that.
	if err != nil {
		return ErrControlSchema
	}

	if component.Key == "" {
		component.Key = getKey(componentDir)
	}
	if openControl.Components.CompareAndAdd(component) {
		openControl.Justifications.LoadMappings(component)
	}
	return nil
}

// Len retruns the length of a SatisfiesList struct
func (slice SatisfiesList) Len() int {
	return len(slice)
}
