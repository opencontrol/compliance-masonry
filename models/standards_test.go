package models

import "testing"

type standardsTest struct {
	standardsFile    string
	expected         Standard
	expectedControls int
}

var standardsTests = []standardsTest{
	{"./opencontrol_fixtures/standards/NIST-800-53.yaml", Standard{Key: "NIST-800-53"}, 326},
	{"./opencontrol_fixtures/standards/PCI-DSS-MAY-2015.yaml", Standard{Key: "PCI-DSS-MAY-2015"}, 258},
}

func TestLoadStandard(t *testing.T) {
	for _, example := range standardsTests {
		openControl := &OpenControl{Standards: NewStandards()}
		openControl.LoadStandard(example.standardsFile)
		actual := openControl.Standards.Get(example.expected.Key)
		if actual.Key != example.expected.Key {
			t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
		}
		if len(actual.Controls) != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, len(actual.Controls))
		}
	}
}
