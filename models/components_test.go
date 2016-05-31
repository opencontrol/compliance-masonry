package models

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/version"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

type componentTest struct {
	componentDir string
	expected     Component
}

type componentTestError struct {
	componentDir  string
	expectedError error
}

var componentTests = []componentTest{
	// Check that a component without a key loads correctly, uses the key of its directory and loads correctly
	{
		filepath.Join("..", "fixtures", "component_fixtures", "EC2"),
		Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    &GeneralReferences{{}},
			Verifications: &VerificationReferences{{}, {}},
			Satisfies: &SatisfiesList{
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for CM-2"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for CM-2"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1.1"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1.1"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Text: "Justification in narrative form for 2.1"},
					},
				},
			},
			SchemaVersion: semver.MustParse("3.0.0"),
		},
	},
	// Check that a component with a key
	{
		filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKey"),
		Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    &GeneralReferences{{}},
			Verifications: &VerificationReferences{{}, {}},
			Satisfies: &SatisfiesList{
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for CM-2"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for CM-2"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1.1"},
						NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1.1"},
					},
				},
				{
					Narrative: []NarrativeSection{
						NarrativeSection{Text: "Justification in narrative form for 2.1"},
					},
				},
			},
			SchemaVersion: semver.MustParse("3.0.0"),
		},
	},
}

func testSet(example componentTest, actual *Component, t *testing.T) {
	// Check that the key was loaded
	assert.Equal(t, example.expected.Key, actual.Key)

	// Check that the name was loaded
	assert.Equal(t, example.expected.Name, actual.Name)

	// Check that the schema version was loaded
	assert.Equal(t, example.expected.SchemaVersion, actual.SchemaVersion)

	// Check that the narrative equals
	if assert.Equal(t, example.expected.Satisfies.Len(), actual.Satisfies.Len()) {
		for idx, _ := range *actual.Satisfies {
			assert.Equal(t, (*example.expected.Satisfies)[idx].Narrative, (*actual.Satisfies)[idx].Narrative)
		}
	}

	// Check that the references were loaded
	assert.Equal(t, example.expected.References.Len(), actual.References.Len())

	// Check that the satisfies data were loaded
	assert.Equal(t, example.expected.Satisfies.Len(), actual.Satisfies.Len())

	// Check that the verifications were loaded
	assert.Equal(t, example.expected.Verifications.Len(), actual.Verifications.Len())
}

func TestLoadComponent(t *testing.T) {
	for _, example := range componentTests {
		openControl := &OpenControl{
			Justifications: NewJustifications(),
			Components:     NewComponents(),
		}
		err := openControl.LoadComponent(example.componentDir)
		if !assert.Nil(t, err) {
			t.Errorf("Expected reading component found in %s to be successful", example.componentDir)
			continue
		}
		// Check the test set with the GetAndApply function
		openControl.Components.GetAndApply(example.expected.Key, func(actual *Component) {
			testSet(example, actual, t)
		})
		// Check the test set with the simple Get function
		actualComponent := openControl.Components.Get(example.expected.Key)
		testSet(example, actualComponent, t)
	}
}

var componentTestErrors = []componentTestError{
	// Check loading a component with no file
	{"", ErrComponentFileDNE},
	// Check loading a component with a broken schema
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2BrokenControl"), ErrControlSchema},
	//Check loading an older schema without a key.
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion2_0"), version.NewIncompatibleVersionError(filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion2_0", "component.yaml"), "component", semver.MustParse("2.0.0"), constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion)},
	//Check loading an older schema with a key.
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKeyBadVersion2_0"), version.NewIncompatibleVersionError(filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKeyBadVersion2_0", "component.yaml"), "component", semver.MustParse("2.0.0"), constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion)},
	// Check for versions not in semver format.
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion_NotSemver"), componentLoadError{fmt.Sprintf(constants.ErrVersionNotInSemverFormatf, 3)}},
	// Check for version that can't be parsed from string to semver
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion_BadVersionString"), componentLoadError{fmt.Sprintf(constants.ErrMissingVersion)}},
	// Check for version because it's missing
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion_Missing"), componentLoadError{fmt.Sprintf(constants.ErrMissingVersion)}},
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		// Check that the expected error is the actual error
		assert.Equal(t, example.expectedError, actualError)
	}
}

func TestComponentLoadErrorString(t *testing.T) {
	err := componentLoadError{"test message"}
	assert.Equal(t, "test message", err.Error())
}
