package models

import (
	"path/filepath"
	"testing"
)

type certificationTest struct {
	certificationFile string
	expected          Certification
	expectedStandards int
	expectedControls  int
}

type certificationTestError struct {
	certificationFile string
	expectedError     error
}

type standardOrderTest struct {
	certification Certification
	expectedOrder string
}

var certificationTests = []certificationTest{
	// Test loading a certification file that has the LATO key, 2 standards, and 6 controls.
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"), Certification{Key: "LATO"}, 2, 6},
}

func TestLoadCertification(t *testing.T) {
	for _, example := range certificationTests {
		openControl := &OpenControl{}
		openControl.LoadCertification(example.certificationFile)
		actual := openControl.Certification
		// Check if loaded certification has the expected key
		if actual.Key != example.expected.Key {
			t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
		}
		// Check if loaded certification has the expected number of standards
		if len(actual.Standards) != example.expectedStandards {
			t.Errorf("Expected %d, Actual: %d", example.expectedStandards, len(actual.Standards))
		}
		// Get the length of the control by using the GetSortedData method
		totalControls := 0
		actual.GetSortedData(func(_ string, _ string) {
			totalControls++
		})
		// Check if loaded certification has the expected number of controls
		if totalControls != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, totalControls)
		}
	}
}

var certificationTestErrors = []certificationTestError{
	// Test a file that can't be read
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "certifications"), ErrReadFile},
	// Test a file that has a broken schema
	{filepath.Join("..", "fixtures", "opencontrol_fixtures", "components", "EC2", "artifact-ec2-1.png"), ErrCertificationSchema},
}

func TestLoadCertificationErrors(t *testing.T) {
	for _, example := range certificationTestErrors {
		openControl := &OpenControl{}
		actualError := openControl.LoadCertification(example.certificationFile)
		// Check that the expected error is the actual error returned
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}

var standardOrderTests = []standardOrderTest{
	{
		// Verify Natural sort order
		Certification{Standards: map[string]Standard{
			"A": Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
			"B": Standard{Controls: map[string]Control{"12": Control{}, "2": Control{}, "1": Control{}}},
			"C": Standard{Controls: map[string]Control{"2": Control{}, "11": Control{}, "101": Control{}, "1000": Control{}, "100": Control{}, "10": Control{}, "1": Control{}}},
		}},
		"A1A2A3B1B2B12C1C2C10C11C100C101C1000",
	},
	{
		// Check that data is returned in order given letters and numbers
		Certification{Standards: map[string]Standard{
			"1":  Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
			"B":  Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
			"B2": Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
		}},
		"111213B1B2B3B21B22B23",
	},
}

func TestStandardOrder(t *testing.T) {
	for _, example := range standardOrderTests {
		actualOrder := ""
		example.certification.GetSortedData(func(standardKey string, controlKey string) {
			actualOrder += standardKey + controlKey
		})
		// Verify that the actual order is the expected order
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}
