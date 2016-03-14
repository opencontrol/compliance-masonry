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

type standardOrderTest struct {
	certification Certification
	expectedOrder string
}

var standardOrderTests = []standardOrderTest{
	{
		Certification{Standards: map[string]Standard{
			"A": Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
			"B": Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
			"C": Standard{Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}}},
		}},
		"A1A2A3B1B2B3C1C2C3",
	},
	{
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
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}
