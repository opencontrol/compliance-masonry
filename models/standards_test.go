package models

import (
	"path/filepath"
	"testing"
)

type standardsTest struct {
	standardsFile    string
	expected         Standard
	expectedControls int
}

type standardTestError struct {
	standardsFile string
	expectedError error
}

type controlOrderTest struct {
	standard      Standard
	expectedOrder string
}

var standardsTests = []standardsTest{
	// Check loading a standard that has 326 controls
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "standards", "NIST-800-53.yaml"), Standard{Name: "NIST-800-53"}, 326},
	// Check loading a standard that has 258 controls
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "standards", "PCI-DSS-MAY-2015.yaml"), Standard{Name: "PCI-DSS-MAY-2015"}, 258},
}

func TestLoadStandard(t *testing.T) {
	for _, example := range standardsTests {
		openControl := &OpenControl{Standards: NewStandards()}
		openControl.LoadStandard(example.standardsFile)
		actual := openControl.Standards.Get(example.expected.Name)
		// Check that the name of the standard was correctly loaded
		if actual.Name != example.expected.Name {
			t.Errorf("Expected %s, Actual: %s", example.expected.Name, actual.Name)
		}
		// Get the length of the control by using the GetSortedData method
		totalControls := 0
		actual.GetSortedData(func(_ string) {
			totalControls++
		})
		if totalControls != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, totalControls)
		}
	}
}

var controlOrderTests = []controlOrderTest{
	// Verify that numeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}},
		},
		"123",
	},
	// Verify that alphabetical controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"c": Control{}, "b": Control{}, "a": Control{}},
		},
		"abc",
	},
	// Verify that alphanumeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"1": Control{}, "b": Control{}, "2": Control{}},
		},
		"12b",
	},
	// Verify that complex alphanumeric controls are ordered correctly
	{
		Standard{
			Controls: map[string]Control{"AC-1": Control{}, "AB-2": Control{}, "1.1.1": Control{}, "2.1.1": Control{}},
		},
		"1.1.12.1.1AB-2AC-1",
	},
	// Verify Natural sort order
	{
		Standard{
			Controls: map[string]Control{"AC-1": Control{}, "AC-12": Control{}, "AC-2 (1)": Control{}, "AC-2 (11)": Control{}, "AC-2 (3)": Control{}, "AC-3 (1)": Control{}},
		},
		"AC-1AC-2 (1)AC-2 (3)AC-2 (11)AC-3 (1)AC-12",
	},
}

func TestControlOrder(t *testing.T) {
	for _, example := range controlOrderTests {
		actualOrder := ""
		example.standard.GetSortedData(func(controlKey string) {
			actualOrder += controlKey
		})
		// Check that the expected order is the actual order
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}

var standardTestErrors = []standardTestError{
	// Check the error loading a file that doesn't exist
	{"", ErrReadFile},
	// Check the error loading a file that has a broken schema
	{filepath.Join("..", "fixtures", "standards_fixtures", "BrokenStandard", "NIST-800-53.yaml"), ErrStandardSchema},
}

func TestLoadStandardsErrors(t *testing.T) {
	for _, example := range standardTestErrors {
		openControl := &OpenControl{}
		actualError := openControl.LoadStandard(example.standardsFile)
		// Check that the expected error and the actual error are the same
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}
