package base

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

type Component interface {
	GetName() string
	GetKey() string
	SetKey(string)
	GetAllSatisfies() []Satisfies
	GetVerifications() common.VerificationReferences
	GetReferences() common.GeneralReferences
}

type Satisfies interface {
	GetStandardKey() string
	GetControlKey() string
	GetNarratives() []Narrative
	GetCoveredBy() common.CoveredByList
}

type Narrative interface {
	GetKey() string
	GetText() string
}

// BaseComponent is the common struct that all component schemas must have.
type Base struct {
	// SchemaVersion contains the schema version.
	SchemaVersion semver.Version `yaml:"-"`
}

// UnmarshalYAML is a overridden implementation of YAML parsing the component.yaml
// This method is similar to the one found here: http://choly.ca/post/go-json-marshalling/
// This is necessary because we want to have backwards compatibility with parsing the old types of version 2.0
// (type =float).
// To compensate for that, we have to hand roll our own UnmarshalYAML that can decide what to do for parsing
// the older version of type float and converting it into semver. In addition, we will use this logic to parse strings
// into semver.
func (c *Base) UnmarshalYAML(unmarshal func(v interface{}) error) error {
	// When we call "unmarshal" callback on an object, it will call that object's "UnmarshalYAML" if defined.
	// Since we are currently in the implementation of Component's "UnmarshalYAML", when finally we call
	// unmarshal again, if it's on type Component, we would end up in a recursive infinite loop.
	// To prevent this, we create a separate type, called Alias.
	type Alias Base
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
	*c = (Base)(aux.Alias)

	// Get the version
	c.SchemaVersion = ver
	return err
}

type componentLoadError struct {
	message string
}

// Error implements the error interface by simply returning the message as a string.
func (e componentLoadError) Error() string {
	return e.message
}
