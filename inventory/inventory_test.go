package inventory

import "testing"

//"github.com/opencontrol/compliance-masonry-go/models"

type getControlInfoTest struct {
	componentsDir string
	standard      string
	control       string
	controlInfo   *ControlInfo
}

var getControlInfoTests = []getControlInfoTest{
	// Check getting info for a control that is not documented
	{
		"../fixtures/opencontrol_fixtures/components",
		"NIST-800-53",
		"AU-6",
		&ControlInfo{false, ImplementationMapping{mapping: map[string]string{}}},
	},
	// Check getting info for a control that is documented
	{
		"../fixtures/opencontrol_fixtures/components",
		"NIST-800-53",
		"CM-2",
		&ControlInfo{true, ImplementationMapping{mapping: map[string]string{"EC2": "partial"}}},
	},
}

func TestGetControlInfo(t *testing.T) {
	for _, example := range getControlInfoTests {
		inventory, _ := InitInventory()
		inventory.LoadComponents(example.componentsDir)
		actualInfo := inventory.GetControlInfo(example.standard, example.control)
		// Check if the documentation exists
		if actualInfo.Exists != example.controlInfo.Exists {
			t.Error("Did not correctly return Exists flag")
		}
		// Check the length of the Implementations mapping
		if len(actualInfo.Implementations.mapping) != len(example.controlInfo.Implementations.mapping) {
			t.Error("Did not correctly return the number of Implementations")
		}
		// Check that the expected mappings are the actual mappings
		for component, expectedImplementation := range example.controlInfo.Implementations.mapping {
			if expectedImplementation != actualInfo.Implementations.mapping[component] {
				t.Error("Did not correctly return t")
			}
		}
	}
}

func TestInitInventory(t *testing.T) {
	_, err := InitInventory()
	if err != nil {
		t.Error("InitInventory failed")
	}
}

func TestInitControlInfo(t *testing.T) {
	_, err := InitControlInfo()
	if err != nil {
		t.Error("InitInventory failed")
	}
}
