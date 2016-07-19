package versions_test

import (
	"errors"
	"path/filepath"
	"testing"

	"github.com/blang/semver"
	"github.com/opencontrol/compliance-masonry/config"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	v3 "github.com/opencontrol/compliance-masonry/models/components/versions/3_0_0"
	v31 "github.com/opencontrol/compliance-masonry/models/components/versions/3_1_0"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/stretchr/testify/assert"
)

type componentV3_1Test struct {
	componentDir string
	expected     v31.Component
}

type componentV3Test struct {
	componentDir string
	expected     v3.Component
}

type componentV2Test struct {
	componentDir string
	expected     v2.Component
}

type componentTestError struct {
	componentDir  string
	expectedError error
}

var v3_1Satisfies = []v31.Satisfies{
	{
		Narrative: []v31.NarrativeSection{
			v31.NarrativeSection{Key: "a", Text: "Justification in narrative form A for CM-2"},
			v31.NarrativeSection{Key: "b", Text: "Justification in narrative form B for CM-2"},
		},
		ControlOrigin:          "shared",
		ControlOrigins:         []string{"shared", "inherited"},
		ImplementationStatus:   "partial",
		ImplementationStatuses: []string{"planned", "partial"},
	},
	{
		Narrative: []v31.NarrativeSection{
			v31.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1"},
			v31.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1"},
		},
		Parameters: []v31.Section{
			v31.Section{Key: "a", Text: "Parameter A for 1.1"},
			v31.Section{Key: "b", Text: "Parameter B for 1.1"},
		},
		ControlOrigin:          "inherited",
		ControlOrigins:         []string{"inherited"},
		ImplementationStatus:   "partial",
		ImplementationStatuses: []string{"partial"},
	},
	{
		Narrative: []v31.NarrativeSection{
			v31.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1.1"},
			v31.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1.1"},
		},
		Parameters: []v31.Section{
			v31.Section{Key: "a", Text: "Parameter A for 1.1.1"},
			v31.Section{Key: "b", Text: "Parameter B for 1.1.1"},
		},
		ControlOrigin:          "inherited",
		ImplementationStatus:   "partial",
		ImplementationStatuses: []string{"partial"},
	},
	{
		Narrative: []v31.NarrativeSection{
			v31.NarrativeSection{Text: "Justification in narrative form for 2.1"},
		},
		ControlOrigin:          "inherited",
		ImplementationStatus:   "partial",
		ImplementationStatuses: []string{"partial"},
	},
}

var componentV3_1Tests = []componentV3_1Test{
	// Check that a component with a key loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_1_0", "EC2"), v31.Component{
		Name:            "Amazon Elastic Compute Cloud",
		Key:             "EC2",
		References:      common.GeneralReferences{{}},
		Verifications:   common.VerificationReferences{{}, {}},
		Satisfies:       v3_1Satisfies,
		SchemaVersion:   semver.MustParse("3.1.0"),
		ResponsibleRole: "AWS Staff",
	}},
}

var v3Satisfies = []v3.Satisfies{
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for CM-2"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for CM-2"},
		},
		ControlOrigin:        "shared",
		ImplementationStatus: "partial",
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1"},
		},
		Parameters: []v3.Section{
			v3.Section{Key: "a", Text: "Parameter A for 1.1"},
			v3.Section{Key: "b", Text: "Parameter B for 1.1"},
		},
		ControlOrigin:        "inherited",
		ImplementationStatus: "partial",
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Key: "a", Text: "Justification in narrative form A for 1.1.1"},
			v3.NarrativeSection{Key: "b", Text: "Justification in narrative form B for 1.1.1"},
		},
		ControlOrigin:        "inherited",
		ImplementationStatus: "partial",
		Parameters: []v3.Section{
			v3.Section{Key: "a", Text: "Parameter A for 1.1.1"},
			v3.Section{Key: "b", Text: "Parameter B for 1.1.1"},
		},
	},
	{
		Narrative: []v3.NarrativeSection{
			v3.NarrativeSection{Text: "Justification in narrative form for 2.1"},
		},
		ControlOrigin:        "inherited",
		ImplementationStatus: "partial",
	},
}

var componentV3Tests = []componentV3Test{
	// Check that a component with a key loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2"), v3.Component{
		Name:            "Amazon Elastic Compute Cloud",
		Key:             "EC2",
		References:      common.GeneralReferences{{}},
		Verifications:   common.VerificationReferences{{}, {}},
		Satisfies:       v3Satisfies,
		SchemaVersion:   semver.MustParse("3.0.0"),
		ResponsibleRole: "AWS Staff",
	}},
	// Check that a component with no key, uses the key of its directory and loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v3_0_0", "EC2WithKey"), v3.Component{
		Name:            "Amazon Elastic Compute Cloud",
		Key:             "EC2",
		References:      common.GeneralReferences{{}},
		Verifications:   common.VerificationReferences{{}, {}},
		Satisfies:       v3Satisfies,
		SchemaVersion:   semver.MustParse("3.0.0"),
		ResponsibleRole: "AWS Staff",
	}},
}

var v2Satisfies = []v2.Satisfies{
	{
		Narrative:            "Justification in narrative form",
		ImplementationStatus: "partial",
	},
	{
		Narrative:            "Justification in narrative form",
		ImplementationStatus: "partial",
	},
	{
		Narrative:            "Justification in narrative form",
		ImplementationStatus: "partial",
	},
	{
		Narrative:            "Justification in narrative form",
		ImplementationStatus: "partial",
	},
}

var componentV2Tests = []componentV2Test{
	// Check that a component with a key loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2"), v2.Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     v2Satisfies,
		SchemaVersion: semver.MustParse("2.0.0"),
	}},
	// Check that a component with no key, uses the key of its directory and loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2WithKey"), v2.Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     v2Satisfies,
		SchemaVersion: semver.MustParse("2.0.0"),
	}},
}

func testSet(example base.Component, actual base.Component, t *testing.T) {
	// Check that the key was loaded
	if example.GetKey() != actual.GetKey() {
		t.Errorf("Expected %s, Actual: %s", example.GetKey(), actual.GetKey())
	}
	// Check that the name was loaded
	if example.GetName() != actual.GetName() {
		t.Errorf("Expected %s, Actual: %s", example.GetName(), actual.GetName())
	}
	// Check that the schema version was loaded
	if example.GetVersion().NE(actual.GetVersion()) {
		t.Errorf("Expected %v, Actual: %v", example.GetVersion(), actual.GetVersion())
	}
	// Check that the references were loaded
	if example.GetReferences().Len() != actual.GetReferences().Len() {
		t.Errorf("Expected %d, Actual: %d", example.GetReferences().Len(), actual.GetReferences().Len())
	}
	// Check that the satisfies data was loaded
	if len(example.GetAllSatisfies()) != len(actual.GetAllSatisfies()) {
		t.Errorf("Expected %d, Actual: %d", len(example.GetAllSatisfies()), len(actual.GetAllSatisfies()))
	}
	// Check Narratives and Parameters.
	for idx, _ := range actual.GetAllSatisfies() {
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetNarratives(), (actual.GetAllSatisfies())[idx].GetNarratives())
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetParameters(), (actual.GetAllSatisfies())[idx].GetParameters())
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetControlOrigin(), (actual.GetAllSatisfies())[idx].GetControlOrigin())
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetControlOrigins(), (actual.GetAllSatisfies())[idx].GetControlOrigins())
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetImplementationStatus(), (actual.GetAllSatisfies())[idx].GetImplementationStatus())
		assert.Equal(t, (example.GetAllSatisfies())[idx].GetImplementationStatuses(), (actual.GetAllSatisfies())[idx].GetImplementationStatuses())
	}
	// Check the responsible role.
	assert.Equal(t, example.GetResponsibleRole(), actual.GetResponsibleRole())
	// Check that the verifications were loaded
	if example.GetVerifications().Len() != actual.GetVerifications().Len() {
		t.Errorf("Expected %d, Actual: %d", example.GetVerifications().Len(), actual.GetVerifications())
	}
}

func loadValidAndTestComponent(path string, t *testing.T, example base.Component) {
	openControl := models.OpenControl{
		Justifications: models.NewJustifications(),
		Components:     components.NewComponents(),
	}
	err := openControl.LoadComponent(path)
	if !assert.Nil(t, err) {
		t.Fatalf("Expected reading component found in %s to be successful", path)
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
	// v3_1_0 tests
	for _, example := range componentV3_1Tests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
	// V3 tests
	for _, example := range componentV3Tests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
	// V2 tests
	for _, example := range componentV2Tests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
}

var componentTestErrors = []componentTestError{
	// Check loading a component with no file
	{"",
		errors.New(constants.ErrComponentFileDNE)},

	// Check loading a component with a broken schema
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2BrokenControl"),
		errors.New("Unable to parse component " + filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2BrokenControl", "component.yaml") + ". Error: yaml: line 16: did not find expected key")},

	// Check for version that is unsupported
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2UnsupportedVersion"),
		config.ErrUnknownSchemaVersion},

	// Check for the case when someone says they are using a certain version (2.0) but it actually is not
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2_InvalidFieldTypeForVersion2_0"),
		errors.New("Unable to parse component. Please check component.yaml schema for version 2.0.0" +
			"\n\tFile: " +
			filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2_InvalidFieldTypeForVersion2_0", "component.yaml") +
			"\n\tParse error: yaml: unmarshal errors:" +
			"\n  line 9: cannot unmarshal !!str `wrong` into common.CoveredByList")},

	// Check for the case when non-2.0 version is not in semver format.
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "common", "EC2VersionNotSemver"),
		base.NewBaseComponentParseError("Version 1 is not in semver format")},
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &models.OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		// Check that the expected error is the actual error
		if !assert.Equal(t, example.expectedError, actualError) {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}
