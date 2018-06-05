package lib

import (
	"errors"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/stretchr/testify/assert"
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
	expectedStandards        []string
	expectedJustificationNum int
	expectedComponents       int
	expectedCertKey          string
}

type loadStandardsTest struct {
	dir               string
	expectedStandards []string
}

type loadStandardsTestError struct {
	dir           string
	expectedError error
}

type loadComponentsTest struct {
	dir                string
	expectedComponents int
}

type loadComponentsTestError struct {
	dir           string
	expectedError error
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
	{
		filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		[]string{"NIST-800-53", "PCI-DSS-MAY-2015"}, 2, 1, "LATO"},
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
		actualStandardsNum := findNumberOfStandards(actual, example.expectedStandards)
		if actualStandardsNum != len(example.expectedStandards) {
			t.Errorf("Expected: `%d`, Actual: `%d`", example.expectedComponents, actualComponentNum)
		}
		// Check the certification key
		if actual.GetCertification().GetKey() != example.expectedCertKey {
			t.Errorf("Expected: `%s`, Actual: `%s`", actual.GetCertification().GetKey(), example.expectedCertKey)
		}
	}
}

var loadComponentsTests = []loadComponentsTest{
	// Check loading set components that only has one component
	{filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "components"), 1},
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
	}
}

var loadComponentsTestErrors = []loadComponentsTestError{
	{
		filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "missing"),
		errors.New("Error: Unable to read the directory " + filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "missing")),
	},
}

func TestLoadComponentErrors(t *testing.T) {
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
	{
		filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "standards"),
		[]string{"NIST-800-53", "PCI-DSS-MAY-2015"},
	},
}

func TestLoadStandards(t *testing.T) {
	for _, example := range loadStandardsTests {
		ws := NewWorkspace()
		ws.LoadStandards(example.dir)
		actualStandardsCount := findNumberOfStandards(ws, example.expectedStandards)
		// Check that the actual number of standards is the expected number of standards
		if actualStandardsCount != len(example.expectedStandards) {
			t.Errorf("Expected: `%d`, Actual: `%d`", len(example.expectedStandards), actualStandardsCount)
		}
	}
}

var loadStandardsTestErrors = []loadStandardsTestError{
	{
		filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "missing"),
		errors.New("Error: Unable to read the directory " + filepath.Join("..", "..", "test", "fixtures", "opencontrol_fixtures", "missing")),
	},
}

func TestLoadStandardErrors(t *testing.T) {
	for _, example := range loadStandardsTestErrors {
		ws := NewWorkspace()
		actualErrors := ws.LoadStandards(example.dir)
		// Check that the actual error is the expected error
		if !assert.Equal(t, example.expectedError, actualErrors[0]) {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualErrors[0])
		}
	}
}

func findNumberOfStandards(actual common.Workspace, standardKeys []string) int {
	actualStandardsNum := 0
	for _, standardKey := range standardKeys {
		if _, found := actual.GetStandard(standardKey); found {
			actualStandardsNum++
		}
	}
	return actualStandardsNum
}
