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

type controlOrderTest struct {
	standard      Standard
	expectedOrder string
}

var controlOrderTests = []controlOrderTest{
	{
		Standard{
			Controls: map[string]Control{"3": Control{}, "2": Control{}, "1": Control{}},
		},
		"123",
	},
	{
		Standard{
			Controls: map[string]Control{"c": Control{}, "b": Control{}, "a": Control{}},
		},
		"abc",
	},
	{
		Standard{
			Controls: map[string]Control{"1": Control{}, "b": Control{}, "2": Control{}},
		},
		"12b",
	},
	{
		Standard{
			Controls: map[string]Control{"AC-1": Control{}, "AB-2": Control{}, "1.1.1": Control{}, "2.1.1": Control{}},
		},
		"1.1.12.1.1AB-2AC-1",
	},
}

func TestControlOrder(t *testing.T) {
	for _, example := range controlOrderTests {
		actualOrder := ""
		example.standard.GetSortedData(func(controlKey string) {
			actualOrder += controlKey
		})
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}
