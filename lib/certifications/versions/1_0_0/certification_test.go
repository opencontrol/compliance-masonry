package certification_test

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/lib/certifications/versions/1_0_0"
)

type standardOrderTest struct {
	certification certification.Certification
	expectedOrder string
}

var standardOrderTests = []standardOrderTest{
	{
		// Verify Natural sort order
		certification.Certification{Standards: map[string]map[string]interface{}{
			"A": map[string]interface{}{"3": nil, "2": nil, "1": nil},
			"B": map[string]interface{}{"12": nil, "2": nil, "1": nil},
			"C": map[string]interface{}{"2": nil, "11": nil, "101": nil, "1000": nil, "100": nil, "10": nil, "1": nil},
		}},
		"A1A2A3B1B2B12C1C2C10C11C100C101C1000",
	},
	{
		// Check that data is returned in order given letters and numbers
		certification.Certification{Standards: map[string]map[string]interface{}{
			"1":  map[string]interface{}{"3": nil, "2": nil, "1": nil},
			"B":  map[string]interface{}{"3": nil, "2": nil, "1": nil},
			"B2": map[string]interface{}{"3": nil, "2": nil, "1": nil},
		}},
		"111213B1B2B3B21B22B23",
	},
}

func TestStandardOrder(t *testing.T) {
	for _, example := range standardOrderTests {
		actualOrder := ""
		standardKeys := example.certification.GetSortedStandards()
		for _, standardKey := range standardKeys {
			controlKeys := example.certification.GetControlKeysFor(standardKey)
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
	cert := certification.Certification{Key: "test"}
	if cert.GetKey() != "test" {
		t.Errorf("GetKey expected test. Actual %s", cert.GetKey())
	}
}
