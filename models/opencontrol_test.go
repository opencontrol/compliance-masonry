package models

import (
	"testing"
)

type keyTest struct {
	input    string
	expected string
}

var keyTests = []keyTest{
	{".", "."},
	{"system/component", "component"},
	{"system", "system"},
}

func TestGetKey(t *testing.T) {
	for _, example := range keyTests {
		actual := getKey(example.input)
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}

type loadDataTest struct {
	openControlDir           string
	certificationPath        string
	expectedStandardsNum     int
	expectedJustificationNum int
	expectedComponents       int
	expectedCertKey          string
}

var loadDataTests = []loadDataTest{
	{"../fixtures/opencontrol_fixtures/", "../fixtures/opencontrol_fixtures/certifications/LATO.yaml", 2, 2, 1, "LATO"},
}

func TestLoadData(t *testing.T) {
	for _, example := range loadDataTests {
		actual := LoadData(example.openControlDir, example.certificationPath)
		actualComponentNum := len(actual.Components.GetAll())
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		actualStandardsNum := len(actual.Standards.GetAll())
		if actualStandardsNum != example.expectedStandardsNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		actualJustificationNum := len(actual.Justifications.mapping)
		if actualJustificationNum != example.expectedJustificationNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		if actual.Certification.Key != example.expectedCertKey {
			t.Errorf("Expected: `%s`, Actual: `%s`", actual.Certification.Key, example.expectedCertKey)
		}
	}
}

type loadComponentsTest struct {
	dir                string
	expectedComponents int
}

var loadComponentsTests = []loadComponentsTest{
	{"../fixtures/opencontrol_fixtures/components", 1},
}

func TestLoadComponents(t *testing.T) {
	for _, example := range loadComponentsTests {
		openControl := NewOpenControl()
		openControl.LoadComponents(example.dir)
		actualComponentNum := len(openControl.Components.GetAll())
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
	}
}

type loadStandardsTest struct {
	dir               string
	expectedStandards int
}

var loadStandardsTests = []loadStandardsTest{
	{"../fixtures/opencontrol_fixtures/standards", 2},
}

func TestLoadStandards(t *testing.T) {
	for _, example := range loadStandardsTests {
		openControl := NewOpenControl()
		openControl.LoadStandards(example.dir)
		actualStandards := len(openControl.Standards.GetAll())
		if actualStandards != example.expectedStandards {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedStandards, actualStandards)
		}
	}
}
