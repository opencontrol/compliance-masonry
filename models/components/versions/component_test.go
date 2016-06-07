package versions_test

import (
	"github.com/opencontrol/compliance-masonry/models/common"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"

	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/config"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/constants"
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
	// Check that a component without a key loads correctly, uses the key of its directory and loads correctly
	{
		filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2"),
		v3.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies:     regularV3Satisfies,
		},
	},
	// Check that a component with a key
	{
		filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2WithKey"),
		v3.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies:     regularV3Satisfies,
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
		filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2"),
		v2.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies:     regularV2Satisfies,
		},
	},
	// Check that a component with a key
	{
		filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2WithKey"),
		v2.Component{
			Name:          "Amazon Elastic Compute Cloud",
			Key:           "EC2",
			References:    common.GeneralReferences{{}},
			Verifications: common.VerificationReferences{{}, {}},
			Satisfies:     regularV2Satisfies,
		},
	},
}

func testSet(example base.Component, actual base.Component, t *testing.T) {
	// Check that the key was loaded
	assert.Equal(t, example.GetKey(), actual.GetKey())

	// Check that the name was loaded
	assert.Equal(t, example.GetName(), actual.GetName())

	// Check that the narrative equals
	if assert.Equal(t, len(example.GetAllSatisfies()), len(actual.GetAllSatisfies())) {
		for i, satisfies := range actual.GetAllSatisfies() {
			for idx, narrative := range satisfies.GetNarratives() {
				assert.Equal(t, (example.GetAllSatisfies())[i].GetNarratives()[idx].GetKey(), narrative.GetKey())
				assert.Equal(t, (example.GetAllSatisfies())[i].GetNarratives()[idx].GetText(), narrative.GetText())
			}
		}
	}

	// Check that the references were loaded
	assert.Equal(t, example.GetReferences().Len(), actual.GetReferences().Len())

	// Check that the satisfies data were loaded
	assert.Equal(t, len(example.GetAllSatisfies()), len(actual.GetAllSatisfies()))

	// Check that the verifications were loaded
	assert.Equal(t, example.GetVerifications().Len(), actual.GetVerifications().Len())
}

func loadValidAndTestComponent(path string, t *testing.T, example base.Component) {
	openControl := &models.OpenControl{
		Justifications: models.NewJustifications(),
		Components:     components.NewComponents(),
	}
	err := openControl.LoadComponent(path)
	if !assert.Nil(t, err) {
		t.Errorf("Expected reading component found in %s to be successful", path)
	}

	// Check the test set with the GetAndApply function
	openControl.Components.GetAndApply(example.GetKey(), func(actual base.Component) {
		testSet(example, actual, t)
	})
	// Check the test set with the simple Get function
	actualComponent := openControl.Components.Get(example.GetKey())
	testSet(example, actualComponent, t)

}

func TestLoadComponent(t *testing.T) {
	// Test Version 3.0.0
	for _, example := range componentV3Tests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
	// Test Version 2.0.0
	for _, example := range componentV2Tests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
}

var componentTestErrors = []componentTestError{
	// Check loading a component with no file
	{"", constants.ErrComponentFileDNE},
	// Check loading a component with a broken schema
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2BrokenControl"), errors.New(constants.ErrMissingVersion)},
	// Check for versions not in semver format for all versions > 2.0.0.
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2BadVersion_NotSemver"), errors.New(constants.ErrMissingVersion)},
	// Check for version that can't be parsed from string to semver
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2BadVersion_BadVersionString"), errors.New(constants.ErrMissingVersion)},
	// Check for version because it's missing
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2BadVersion_Missing"), errors.New(constants.ErrMissingVersion)},
	// Check for version that is unsupported
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2UnsupportedVersion"), config.ErrUnknownSchemaVersion},
	// Check for the case when someone says they are using a certain version but it actually is not.
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2_InvalidFieldTypeForVersion"), fmt.Errorf(constants.ErrComponentSchemaParsef, semver.MustParse("3.0.0"))},
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &models.OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		// Check that the expected error is the actual error
		assert.Equal(t, example.expectedError, actualError)
	}
}
