package certifications

import (
	"testing"
	"path/filepath"
	v1_0_0 "github.com/opencontrol/compliance-masonry/lib/certifications/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

type v1certificationTest struct {
	certificationFile string
	expected          v1_0_0.Certification
	expectedStandards int
	expectedControls  int
}

type certificationTestError struct {
	certificationFile string
	expectedError     error
}

var v1certificationTests = []v1certificationTest{
	// Test loading a certification file that has the LATO key, 2 standards, and 6 controls.
	{filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"), v1_0_0.Certification{Key: "LATO"}, 2, 6},
}

func TestLoadCertification(t *testing.T) {
	for _, example := range v1certificationTests {

		actual, _ := Load(example.certificationFile)
		// Check if loaded certification has the expected key
		if actual.GetKey() != example.expected.GetKey() {
			t.Errorf("Expected %s, Actual: %s", example.expected.GetKey(), actual.GetKey())
		}
		// Check if loaded certification has the expected number of standards
		if len(actual.GetStandards()) != example.expectedStandards {
			t.Errorf("Expected %d, Actual: %d", example.expectedStandards, len(actual.GetStandards()))
		}
		// Get the length of the control by using the GetSortedData method
		totalControls := 0
		standardKeys := actual.GetSortedStandards()
		for _, standardKey := range standardKeys {
			totalControls += len(actual.GetStandards()[standardKey].GetSortedControls())
		}
		// Check if loaded certification has the expected number of controls
		if totalControls != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, totalControls)
		}
	}
}

var certificationTestErrors = []certificationTestError{
	// Test a file that can't be read
	{filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "certifications"), common.ErrReadFile},
	// Test a file that has a broken schema
	{filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "components", "EC2", "artifact-ec2-1.png"), common.ErrCertificationSchema},
}

func TestLoadCertificationErrors(t *testing.T) {
	for _, example := range certificationTestErrors {
		_, actualError := Load(example.certificationFile)
		// Check that the expected error is the actual error returned
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}