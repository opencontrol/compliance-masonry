package certifications_test

import (
	"github.com/opencontrol/compliance-masonry/lib/certifications"
	v1_0_0 "github.com/opencontrol/compliance-masonry/lib/certifications/versions/1_0_0"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

type v1certificationTest struct {
	certificationFile string
	expected          v1_0_0.Certification
	expectedControls  map[string][]string
}

type certificationTestError struct {
	certificationFile string
	expectedError     error
}

var v1certificationTests = []v1certificationTest{
	// Test loading a certification file that has the LATO key, 2 standards, and 6 controls.
	{
		filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		v1_0_0.Certification{Key: "LATO"},
		map[string][]string{
			"NIST-800-53":      []string{"AC-2", "AC-6", "CM-2"},
			"PCI-DSS-MAY-2015": []string{"1.1", "1.1.1", "2.1"},
		},
	},
}

func TestLoadCertification(t *testing.T) {
	for _, example := range v1certificationTests {

		actual, err := certifications.Load(example.certificationFile)
		assert.Nil(t, err)
		// Check if loaded certification has the expected key
		if actual.GetKey() != example.expected.GetKey() {
			t.Errorf("Expected %s, Actual: %s", example.expected.GetKey(), actual.GetKey())
		}
		for expectedStandardKey, expectedControls := range example.expectedControls {
			assert.Equal(t, len(actual.GetControlKeysFor(expectedStandardKey)), len(expectedControls))
			assert.Equal(t, actual.GetControlKeysFor(expectedStandardKey), expectedControls)
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
		_, actualError := certifications.Load(example.certificationFile)
		// Check that the expected error is the actual error returned
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}
