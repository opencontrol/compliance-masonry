package models

import "testing"

type certificationTest struct {
	certificationFile string
	expected          Certification
	expectedStandards int
	expectedControls  int
}

var certificationTests = []certificationTest{
	{"./opencontrol_fixtures/certifications/LATO.yaml", Certification{Key: "LATO"}, 2, 6},
}

func TestLoadCertification(t *testing.T) {
	for _, example := range certificationTests {
		openControl := &OpenControl{}
		openControl.LoadCertification(example.certificationFile)
		actual := openControl.Certification
		if actual.Key != example.expected.Key {
			t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
		}

		if len(actual.Standards) != example.expectedStandards {
			t.Errorf("Expected %d, Actual: %d", example.expectedStandards, len(actual.Standards))
		}
		// Get the length of the control by using the GetSortedData method
		totalControls := 0
		actual.GetSortedData(func(_ string, _ string) {
			totalControls++
		})

		if totalControls != example.expectedControls {
			t.Errorf("Expected %d, Actual: %d", example.expectedControls, totalControls)
		}
	}
}
