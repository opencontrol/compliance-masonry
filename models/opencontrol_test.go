package models

import (
	"path/filepath"
	"testing"
)

type keyTest struct {
	input    string
	expected string
}

type loadDataTest struct {
	openControlDir           string
	certificationPath        string
	expectedStandardsNum     int
	expectedJustificationNum int
	expectedComponents       int
	expectedCertKey          string
}

type loadStandardsTest struct {
	dir               string
	expectedStandards int
}

type loadComponentsTest struct {
	dir                string
	expectedComponents int
}

var keyTests = []keyTest{
	// Check that the key is extracted by using the local directory
	{".", "."},
	// Check that the key is extracted from the component directory
	{"system/component", "component"},
	// Check that the key is extracted by using the system directory
	{"system", "system"},
}

func TestGetKey(t *testing.T) {
	for _, example := range keyTests {
		actual := getKey(example.input)
		// Check that the actual key is the expected key
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}

var loadDataTests = []loadDataTest{
	// Load a fixtures that has 2 component, 2 standards, and a certification called LATO
	{filepath.Join("..", "fixtures", "opencontrol_fixtures"), filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"), 2, 2, 1, "LATO"},
}

func TestLoadData(t *testing.T) {
	for _, example := range loadDataTests {
		actual := LoadData(example.openControlDir, example.certificationPath)
		actualComponentNum := len(actual.Components.GetAll())
		// Check the number of components
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Check the number of standards
		actualStandardsNum := len(actual.Standards.GetAll())
		if actualStandardsNum != example.expectedStandardsNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Check the number of justifications
		actualJustificationNum := len(actual.Justifications.mapping)
		if actualJustificationNum != example.expectedJustificationNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Check the certification key
		if actual.Certification.Key != example.expectedCertKey {
			t.Errorf("Expected: `%s`, Actual: `%s`", actual.Certification.Key, example.expectedCertKey)
		}
	}
}

var loadComponentsTests = []loadComponentsTest{
	// Check loading set components that only has one component
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "components"), 1},
}

func TestLoadComponents(t *testing.T) {
	for _, example := range loadComponentsTests {
		openControl := NewOpenControl()
		openControl.LoadComponents(example.dir)
		actualComponentNum := len(openControl.Components.GetAll())
		// Check the number of components
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
	}
}

var loadStandardsTests = []loadStandardsTest{
	// Load a series of standards file that have 2 standards
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "standards"), 2},
}

func TestLoadStandards(t *testing.T) {
	for _, example := range loadStandardsTests {
		openControl := NewOpenControl()
		openControl.LoadStandards(example.dir)
		actualStandards := len(openControl.Standards.GetAll())
		// Check that the actual number of standards is the expected number of standards
		if actualStandards != example.expectedStandards {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedStandards, actualStandards)
		}
	}
}
