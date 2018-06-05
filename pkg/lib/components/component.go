package components

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"path/filepath"
)

// NewComponentParseError is a constructor for creating errors of type BaseComponentParseError
func NewComponentParseError(message string) BaseComponentParseError {
	return BaseComponentParseError{message}
}

// Load will read the file at the given path and attempt to return a component object.
func Load(path string) (common.Component, error) {
	// Get file system assistance.
	fs := fs.OSUtil{}
	// Read the component file.
	fileName := filepath.Join(path, "component.yaml")
	componentData, err := fs.OpenAndReadFile(fileName)
	if err != nil {
		return nil, errors.New(constants.ErrComponentFileDNE)
	}
	// Parse the component.
	var component common.Component
	component, err = parseComponent(componentData, fileName)
	if err != nil {
		return nil, err
	}
	// Ensure we have a key for the component.
	if component.GetKey() == "" {
		component.SetKey(getKey(path))
	}
	return component, nil
}

// BaseComponentParseError is the type of error that will be returned if the parsing failed for ONLY the `Base` struct.
type BaseComponentParseError struct {
	message string
}

func (b BaseComponentParseError) Error() string {
	return b.message
}

// Base is the bare minimum that every component YAML will have and is used to find the schema version.
// Complete implementations of component do not need to embed this struct or put it as a field in the component.
// When this struct is used in the ParseComponent function, it will transfer the version from this struct to the
// final component struct via SetVersion.
type Base struct {
	SchemaVersion semver.Version `yaml:"-" json:"-"`
}

// UnmarshalYAML is a overridden implementation of YAML parsing the component.yaml
// This method is similar to the one found here: http://choly.ca/post/go-json-marshalling/
// This is necessary because we want to have backwards compatibility with parsing the old types of version 2.0
// (type =float).
// To compensate for that, we have to hand roll our own UnmarshalYAML that can decide what to do for parsing
// the older version of type float and converting it into semver. In addition, we will use this logic to parse strings
// into semver.
func (b *Base) UnmarshalYAML(unmarshal func(v interface{}) error) error {
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
		Alias: (Alias)(*b),
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
			return BaseComponentParseError{fmt.Sprintf("Version %v is not in semver format", v)}

		}
	// The interface type will default to string if not numeric which is what all semver types will be initially.
	case string:
		ver, versionErr = semver.Parse(v)
		if versionErr != nil {
			return BaseComponentParseError{constants.ErrMissingVersion}
		}
	// In the case, it's just missing completely.
	default:
		return BaseComponentParseError{constants.ErrMissingVersion}
	}
	// Copy everything from the Alias back to the original component.
	*b = (Base)(aux.Alias)

	// Get the version
	b.SchemaVersion = ver
	return err
}

// getKey extracts a component key from the filepath
func getKey(filePath string) string {
	_, key := filepath.Split(filePath)
	return key
}
