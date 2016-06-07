package components_test

import (
	"github.com/opencontrol/compliance-masonry/models/common"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/models/components"
	"errors"
)

type componentV3Test struct {
	componentDir string
	expected     v3.Component
}

type componentTestError struct {
	componentDir  string
	expectedError error
}

var regularV3Satisfies []v3.Satisfies = []v3.Satisfies{
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for CM-2"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for CM-2"},
		},
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1"},
		},
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1.1"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1.1"},
		},
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Text: "Justification in narrative form for 2.1"},
		},
	},
}

var componentV3Tests = []componentV3Test{
	/*
		Version 3.0.0
	 */
	// Check that a component without a key loads correctly, uses the key of its directory and loads correctly
	{
		filepath.Join("..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2"),
		v3.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies: regularV3Satisfies,
		},
	},
	// Check that a component with a key
	{
		filepath.Join("..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2WithKey"),
		v3.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies: regularV3Satisfies,
		},
	},
}



type componentV2Test struct {
	componentDir string
	expected     v2.Component
}

var regularV2Satisfies []v2.Satisfies = []v2.Satisfies{
	{
		Narrative: "CM-2 Narrative",
	},
	{
		Narrative: "1.1 Narrative",
	},
	{
		Narrative: "1.1.1 Narrative",
	},
	{
		Narrative: "2.1 Narrative",
	},
}

var componentV2Tests = []componentV2Test{
	// Check that a component without a key loads correctly, uses the key of its directory and loads correctly
	{
		filepath.Join("..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2"),
		v2.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies: regularV2Satisfies,
			//SchemaVersion: semver.MustParse("3.0.0"),
		},
	},
	// Check that a component with a key
	{
		filepath.Join("..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2WithKey"),
		v2.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies: regularV2Satisfies,
			//SchemaVersion: semver.MustParse("3.0.0"),
		},
	},
}

func testSet(example componentV3Test, actual base.Component, t *testing.T) {
	// Check that the key was loaded
	assert.Equal(t, example.expected.Key, actual.GetKey())

	// Check that the name was loaded
	assert.Equal(t, example.expected.Name, actual.GetName())

	// Check that the schema version was loaded
	//assert.Equal(t, example.expected.SchemaVersion, actual.SchemaVersion)

	// Check that the narrative equals
	if assert.Equal(t, len(example.expected.Satisfies), len(actual.GetAllSatisfies())) {
		for i, satisfies := range actual.GetAllSatisfies() {
			for idx, narrative := range satisfies.GetNarratives() {
				assert.Equal(t, (example.expected.Satisfies)[i].Narrative[idx], narrative)
			}
		}
	}

	// Check that the references were loaded
	assert.Equal(t, example.expected.References.Len(), actual.GetReferences().Len())

	// Check that the satisfies data were loaded
	assert.Equal(t, len(example.expected.Satisfies), len(actual.GetAllSatisfies()))

	// Check that the verifications were loaded
	assert.Equal(t, example.expected.GetVerifications().Len(), actual.GetVerifications().Len())
}

func loadValidComponent(path string, t *testing.T) *models.OpenControl {
	openControl := &models.OpenControl{
		Justifications: models.NewJustifications(),
		Components:     components.NewComponents(),
	}
	err := openControl.LoadComponent(path)
	if !assert.Nil(t, err) {
		t.Errorf("Expected reading component found in %s to be successful", path)
	}
	return openControl
}

func TestLoadComponent(t *testing.T) {
	for _, example := range componentV3Tests {

		openControl := loadValidComponent(example.componentDir, t)
		// Check the test set with the GetAndApply function
		openControl.Components.GetAndApply(example.expected.Key, func(actual base.Component) {
			testSet(example, actual, t)
		})
		// Check the test set with the simple Get function
		actualComponent := openControl.Components.Get(example.expected.Key)
		testSet(example, actualComponent, t)
	}
}

var componentTestErrors = []componentTestError{
	// Check loading a component with no file
	{"", constants.ErrComponentFileDNE},
	// Check loading a component with a broken schema
	{filepath.Join("..", "..", "fixtures", "component_fixtures", "common", "EC2BrokenControl"), errors.New(constants.ErrMissingVersion)},
	//Check loading an older schema without a key.
	//{filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion2_0"), version.NewIncompatibleVersionError(version.NewRequirements(filepath.Join("..", "fixtures", "component_fixtures", "EC2BadVersion2_0", "component.yaml"), "component", semver.MustParse("2.0.0"), constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion))},
	//Check loading an older schema with a key.
	//{filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKeyBadVersion2_0"), version.NewIncompatibleVersionError(version.NewRequirements(filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKeyBadVersion2_0", "component.yaml"), "component", semver.MustParse("2.0.0"), constants.MinComponentYAMLVersion, constants.MaxComponentYAMLVersion))},
	// Check for versions not in semver format for all versions > 2.0.0.
	{filepath.Join("..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2BadVersion_NotSemver"), errors.New(constants.ErrMissingVersion)},
	// Check for version that can't be parsed from string to semver
	{filepath.Join("..", "..", "fixtures", "component_fixtures", "common", "EC2BadVersion_BadVersionString"), errors.New(constants.ErrMissingVersion)},
	// Check for version because it's missing
	{filepath.Join("..", "..", "fixtures", "component_fixtures", "common", "EC2BadVersion_Missing"), errors.New(constants.ErrMissingVersion)},
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &models.OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		// Check that the expected error is the actual error
		assert.Equal(t, example.expectedError, actualError)
	}
}


