package models

import (
	"log"
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
	{"../fixtures/component_fixtures/EC2", Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    &GeneralReferences{{}},
		Verifications: &VerificationReferences{{}, {}},
		Satisfies:     &SatisfiesList{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}},
	{"../fixtures/component_fixtures/EC2WithKey", Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    &GeneralReferences{{}},
		Verifications: &VerificationReferences{{}, {}},
		Satisfies:     &SatisfiesList{{}, {}, {}, {}},
		SchemaVersion: 2.0,
	}},
}

func testSet(example componentTest, actual *Component, t *testing.T) {
	if example.expected.Key != actual.Key {
		t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
	}
	if example.expected.Name != actual.Name {
		t.Errorf("Expected %s, Actual: %s", example.expected.Name, actual.Name)
	}

	if example.expected.SchemaVersion != actual.SchemaVersion {
		t.Errorf("Expected %f, Actual: %f", example.expected.SchemaVersion, actual.SchemaVersion)
	}

	if example.expected.References.Len() != actual.References.Len() {
		t.Errorf("Expected %d, Actual: %d", example.expected.References.Len(), actual.References.Len())
	}

	if example.expected.Satisfies.Len() != actual.Satisfies.Len() {
		t.Errorf("Expected %d, Actual: %d", example.expected.Satisfies.Len(), actual.Satisfies.Len())
	}

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
		// Test Get and Apply
		openControl.Components.GetAndApply(example.expected.Key, func(actual *Component) {
			testSet(example, actual, t)
		})
		// Test Get
		actualComponent := openControl.Components.Get(example.expected.Key)
		testSet(example, actualComponent, t)
	}
}

var componentTestErrors = []componentTestError{
	{"", ErrComponentFileDNE},
	{"../fixtures/component_fixtures/EC2Broken", ErrControlSchema},
}

func TestLoadComponentErrors(t *testing.T) {
	for _, example := range componentTestErrors {
		openControl := &OpenControl{}
		actualError := openControl.LoadComponent(example.componentDir)
		log.Println(actualError)
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}
