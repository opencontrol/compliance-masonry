package models

import (
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
	// Check that a component with a key loads correctly
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2"), Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    &GeneralReferences{{}},
		Verifications: &VerificationReferences{{}, {}},
		Satisfies:     &SatisfiesList{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}},
	// Check that a component with no key, uses the key of its directory and loads correctly
	{filepath.Join("..", "fixtures", "component_fixtures", "EC2WithKey"), Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    &GeneralReferences{{}},
		Verifications: &VerificationReferences{{}, {}},
		Satisfies:     &SatisfiesList{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}},
}

func testSet(example componentTest, actual *Component, t *testing.T) {
	// Check that the key was loaded
	if example.expected.Key != actual.Key {
		t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
	}
	// Check that the name was loaded
	if example.expected.Name != actual.Name {
		t.Errorf("Expected %s, Actual: %s", example.expected.Name, actual.Name)
	}
	// Check that the schema version was loaded
	if example.expected.SchemaVersion != actual.SchemaVersion {
		t.Errorf("Expected %f, Actual: %f", example.expected.SchemaVersion, actual.SchemaVersion)
	}
	// Check that the references were loaded
	if example.expected.References.Len() != actual.References.Len() {
		t.Errorf("Expected %d, Actual: %d", example.expected.References.Len(), actual.References.Len())
	}
	// Check that the satisfies data were loaded
	if example.expected.Satisfies.Len() != actual.Satisfies.Len() {
		t.Errorf("Expected %d, Actual: %d", example.expected.Satisfies.Len(), actual.Satisfies.Len())
	}
	// Check that the verifications were loaded
	if example.expected.Verifications.Len() != actual.Verifications.Len() {
		t.Errorf("Expected %d, Actual: %d", example.expected.Verifications.Len(), actual.Verifications.Len())
	}
}

func TestLoadComponent(t *testing.T) {
	for _, example := range componentTests {
		openControl := &OpenControl{
			Justifications: NewJustifications(),
			Components:     NewComponents(),
		}
		openControl.LoadComponent(example.componentDir)
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
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		// Check that the expected error is the actual error
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}
