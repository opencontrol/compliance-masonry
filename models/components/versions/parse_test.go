package versions_test

import (
	"path/filepath"
	"testing"
	v2 "github.com/opencontrol/compliance-masonry/models/components/versions/2_0_0"
	"github.com/opencontrol/compliance-masonry/models/common"
	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/models/components"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

type componentV2Test struct {
	componentDir string
	expected     v2.Component
}

type componentTestError struct {
	componentDir  string
	expectedError error
}

var componentTests = []componentV2Test{
	// Check that a component with a key loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2"), v2.Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     []v2.Satisfies{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}},
	// Check that a component with no key, uses the key of its directory and loads correctly
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "v2_0_0", "EC2WithKey"), v2.Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    common.GeneralReferences{{}},
		Verifications: common.VerificationReferences{{}, {}},
		Satisfies:     []v2.Satisfies{{}, {}, {}, {}},
		SchemaVersion: 2.0,
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
	if example.GetVersion() != actual.GetVersion() {
		t.Errorf("Expected %f, Actual: %f", example.GetVersion(), actual.GetVersion())
	}
	// Check that the references were loaded
	if example.GetReferences().Len() != actual.GetReferences().Len() {
		t.Errorf("Expected %d, Actual: %d", example.GetReferences().Len(), actual.GetReferences().Len())
	}
	// Check that the satisfies data were loaded
	if len(example.GetAllSatisfies()) != len(actual.GetAllSatisfies()) {
		t.Errorf("Expected %d, Actual: %d", len(example.GetAllSatisfies()), len(actual.GetAllSatisfies()))
	}
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
	for _, example := range componentTests {
		loadValidAndTestComponent(example.componentDir, t, &example.expected)
	}
}

var componentTestErrors = []componentTestError{
	// Check loading a component with no file
	{"", constants.ErrComponentFileDNE},
	// Check loading a component with a broken schema
	{filepath.Join("..", "..", "..", "fixtures", "component_fixtures", "EC2BrokenControl"), constants.ErrMissingVersion},
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
