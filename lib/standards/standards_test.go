package standards

import (
	"path/filepath"
	v1_0_0 "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
	"testing"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

type v1standardsTest struct {
	standardsFile    string
	expected         common.Standard
	expectedControls int
}

var standardsTests = []v1standardsTest{
	// Check loading a standard that has 326 controls
	{filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "standards", "NIST-800-53.yaml"), v1_0_0.Standard{Name: "NIST-800-53"}, 326},
	// Check loading a standard that has 258 controls
	{filepath.Join("..", "..", "fixtures", "opencontrol_fixtures", "standards", "PCI-DSS-MAY-2015.yaml"), v1_0_0.Standard{Name: "PCI-DSS-MAY-2015"}, 258},
}

func TestLoadStandard(t *testing.T) {
	for _, example := range standardsTests {
		actual, err := Load(example.standardsFile)
		if err != nil {
			t.Errorf("Expected nil error, Actual %s", err.Error())
		}
		// Check that the name of the standard was correctly loaded
		if actual.GetName() != example.expected.GetName() {
			t.Errorf("Expected %s, Actual: %s", example.expected.GetName(), actual.GetName())
		}
		// Get the length of the control by using the GetSortedData method
		totalControls := len(actual.GetSortedControls())
		if totalControls != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, totalControls)
		}
	}
}


type standardTestError struct {
	standardsFile string
	expectedError error
}

var standardTestErrors = []standardTestError{
	// Check the error loading a file that doesn't exist
	{"", common.ErrReadFile},
	// Check the error loading a file that has a broken schema
	{filepath.Join("..", "..", "fixtures", "standards_fixtures", "BrokenStandard", "NIST-800-53.yaml"), common.ErrStandardSchema},
}

func TestLoadStandardsErrors(t *testing.T) {
	for _, example := range standardTestErrors {
		_, actualError := Load(example.standardsFile)
		// Check that the expected error and the actual error are the same
		if example.expectedError != actualError {
			t.Errorf("Expected %s, Actual: %s", example.expectedError, actualError)
		}
	}
}