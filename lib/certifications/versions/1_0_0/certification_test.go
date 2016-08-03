package certification

import (
	"testing"
	v1standards "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
)

type standardOrderTest struct {
	certification Certification
	expectedOrder string
}

var standardOrderTests = []standardOrderTest{
	{
		// Verify Natural sort order
		Certification{Standards: map[string]v1standards.Standard{
			"A": {Controls: map[string]v1standards.Control{"3": {}, "2": {}, "1": {}}},
			"B": {Controls: map[string]v1standards.Control{"12": {}, "2": {}, "1": {}}},
			"C": {Controls: map[string]v1standards.Control{"2": {}, "11": {}, "101": {}, "1000": {}, "100": {}, "10": {}, "1": {}}},
		}},
		"A1A2A3B1B2B12C1C2C10C11C100C101C1000",
	},
	{
		// Check that data is returned in order given letters and numbers
		Certification{Standards: map[string]v1standards.Standard{
			"1":  {Controls: map[string]v1standards.Control{"3": {}, "2": {}, "1": {}}},
			"B":  {Controls: map[string]v1standards.Control{"3": {}, "2": {}, "1": {}}},
			"B2": {Controls: map[string]v1standards.Control{"3": {}, "2": {}, "1": {}}},
		}},
		"111213B1B2B3B21B22B23",
	},
}

func TestStandardOrder(t *testing.T) {
	for _, example := range standardOrderTests {
		actualOrder := ""
		standardKeys := example.certification.GetSortedStandards()
		for _, standardKey := range standardKeys {
			controlKeys := example.certification.GetStandards()[standardKey].GetSortedControls()
			for _, controlKey := range controlKeys {
				actualOrder += standardKey + controlKey
			}
		}
		// Verify that the actual order is the expected order
		if actualOrder != example.expectedOrder {
			t.Errorf("Expected %s, Actual: %s", example.expectedOrder, actualOrder)
		}
	}
}

func TestGetKey(t *testing.T) {
	cert := Certification{Key: "test"}
	if cert.GetKey() != "test" {
		t.Errorf("GetKey expected test. Actual %s", cert.GetKey())
	}
}
