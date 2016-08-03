package lib

import (
	"path/filepath"
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

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

type loadStandardsTestError struct {
	dir                string
	expectedError      error
}

type loadComponentsTest struct {
	dir                string
	expectedComponents int
}

type loadComponentsTestError struct {
	dir                string
	expectedError      error
}

var loadDataTests = []loadDataTest{
	// Load a fixtures that has 2 component, 2 standards, and a certification called LATO
	{filepath.Join("..", "fixtures", "opencontrol_fixtures"), filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"), 2, 2, 1, "LATO"},
}

func TestLoadData(t *testing.T) {
	for _, example := range loadDataTests {
		actual, _ := LoadData(example.openControlDir, example.certificationPath)
		actualComponentNum := len(actual.GetAllComponents())
		// Check the number of components
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Check the number of standards
		actualStandardsNum := len(actual.GetStandards())
		if actualStandardsNum != example.expectedStandardsNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		/*
		// Check the number of justifications
		actualJustificationNum := len(actual.Justifications.mapping)
		if actualJustificationNum != example.expectedJustificationNum {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		*/
		// Check the certification key
		if actual.GetCertification().GetKey() != example.expectedCertKey {
			t.Errorf("Expected: `%s`, Actual: `%s`", actual.GetCertification().GetKey(), example.expectedCertKey)
		}
	}
}

var loadComponentsTests = []loadComponentsTest{
	// Check loading set components that only has one component
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "components"), 1},
}

func TestLoadComponents(t *testing.T) {
	for _, example := range loadComponentsTests {
		ws := NewWorkspace()
		ws.LoadComponents(example.dir)
		actualComponentNum := len(ws.GetAllComponents())
		// Check the number of components
		if actualComponentNum != example.expectedComponents {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Test Single Component Get
		component := ws.GetComponent("EC2")
		if component.GetKey() != "EC2" {
			t.Errorf("Expected component key to equal EC2. Actual %s", component.GetKey())
		}
		if component.GetName() != "Amazon Elastic Compute Cloud" {
			t.Errorf("Exepcted component name to equal Amazon Elastic Compute Cloud. Actual %s", component.GetName())
		}
	}
}

var loadComponentsTestErrors = []loadComponentsTestError{
	{
		filepath.Join("..", "fixtures", "opencontrol_fixtures", "missing"),
	 	errors.New("Error: Unable to read the directory "+filepath.Join("..", "fixtures", "opencontrol_fixtures", "missing")),
	},
}

func TestLoadComponentErrors (t *testing.T) {
	for _, example := range loadComponentsTestErrors {
		ws := NewWorkspace()
		actualErrors := ws.LoadComponents(example.dir)
		// Check that the actual error is the expected error
		if !assert.Equal(t, example.expectedError, actualErrors[0]) {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualErrors[0])
		}
	}
}

var loadStandardsTests = []loadStandardsTest{
	// Load a series of standards file that have 2 standards
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "standards"), 2},
}

func TestLoadStandards(t *testing.T) {
	for _, example := range loadStandardsTests {
		ws := NewWorkspace()
		ws.LoadStandards(example.dir)
		actualStandards := len(ws.GetStandards())
		// Check that the actual number of standards is the expected number of standards
		if actualStandards != example.expectedStandards {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedStandards, actualStandards)
		}
		// Test Single Get Standards
		standard := ws.GetStandard("NIST-800-53")
		if standard.GetName() != "NIST-800-53" {
			t.Errorf("Expected standard name NIST-800-53. Actual %s", standard.GetName())
		}
	}
}

var loadStandardsTestErrors = []loadStandardsTestError{
	{
		filepath.Join("..", "fixtures", "opencontrol_fixtures", "missing"),
	 	errors.New("Error: Unable to read the directory "+filepath.Join("..", "fixtures", "opencontrol_fixtures", "missing")),
	},
}

func TestLoadStandardErrors (t *testing.T) {
	for _, example := range loadStandardsTestErrors {
		ws := NewWorkspace()
		actualErrors := ws.LoadStandards(example.dir)
		// Check that the actual error is the expected error
		if !assert.Equal(t, example.expectedError, actualErrors[0]) {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualErrors[0])
		}
	}
}
