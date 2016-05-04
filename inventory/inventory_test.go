package inventory

import (
	"testing"

	"github.com/opencontrol/compliance-masonry/config"
	"github.com/opencontrol/compliance-masonry/models"
)

//"github.com/opencontrol/compliance-masonry-go/models"

type getControlInfoTest struct {
	componentsDir string
	standard      string
	control       string
	controlInfo   *ControlInfo
}

type inventoryTest struct {
	configBytes        []byte
	expectedError      error
	expectedComponents int
	expectedRequiredComponents int
}

type loadLocalComponentsTest struct {
	components         []string
	expectedComponents int
}

type loadRequiredComponentsTest struct {
	controls						[]string
	expectedComponents		int
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
		inventory, _ := InitInventory([]byte(`schema_version: "1.0.0"`))
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

var getLocalComponentsTests = []inventoryTest{
	{
		// Check a schema with 0 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
required_controls:
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 0, 0,
	},
	{
		// Check a schema with 3 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ./component-1
  - ./component-2
  - ./component-3
required_controls:
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 3, 0,
	},
}

func TestGetLocalComponents(t *testing.T) {
	for _, example := range getLocalComponentsTests {
		actualComponents, err := GetLocalComponents(example.configBytes)
		if err != nil {
			t.Error(err)
		}
		if len(actualComponents) != example.expectedComponents {
			t.Error("The number of actual components and expected components do not match")
		}
	}
}

var loadLocalComponentsTests = []loadLocalComponentsTest{
	// Test loading one component
	{[]string{"../fixtures/component_fixtures/EC2"}, 1},
}

func TestLoadLocalComponents(t *testing.T) {
	for _, example := range loadLocalComponentsTests {
		inventory, _ := InitInventory([]byte(`schema_version: "1.0.0"`))
		err := inventory.LoadLocalComponents(example.components)
		if err != nil {
			t.Error(err)
		}
		if len(inventory.Components.GetAll()) != example.expectedComponents {
			t.Error("The number of actual components and expected components do not match")
		}
	}
}

//TODO: start new tests here
var getRequiredComponentsTests = []inventoryTest{
	{
		// Check a schema with 0 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
required_components:
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 0, 0,
	},
	{
		// Check a schema with 3 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ./component-1
  - ./component-2
  - ./component-3
required_components:
  - component-1
  - component-2
  - component-3
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 0, 3,
	},
}

func TestGetRequiredComponents(t *testing.T) {
	for _, example := range getRequiredComponentsTests {
		actualComponents, err := GetRequiredComponents(example.configBytes)
		if err != nil {
			t.Error(err)
		}
		if len(actualComponents) != example.expectedRequiredComponents {
			t.Error("The number of actual required components and expected required components do not match")
		}
	}
}
//TODO: end new tests here

var initInventoryTests = []inventoryTest{
	{
		// Check a schema with 0 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
required_controls:
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 0, 0,
	},
	{
		// Check a schema with 1 components
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ../fixtures/component_fixtures/EC2
required_controls:
certifications:
  - ./cert-1.yaml
standards:
  - ./standard-1.yaml
dependencies:
  certifications:
    - url: github.com/18F/LATO
      revision: master
  systems:
    - url: github.com/18F/cg-complinace
      revision: master
  standards:
    - url: github.com/18F/NIST-800-53
      revision: master
`), nil, 1, 0,
	},
	{
		// Check example with broken opencontrol.yaml
		[]byte(`
schema_versio: "1.0.0"
`), config.ErrCantParseSemver, 1, 0,
	},
	{
		// With broken component
		[]byte(`
schema_version: "1.0.0"
name: test
metadata:
  description: "A system to test parsing"
  maintainers:
    - test@test.com
components:
  - ../fixtures/component_fixtures/EC2BrokenControl
`), models.ErrControlSchema, 1, 0,
	},
}

func TestInitInventory(t *testing.T) {
	for _, example := range initInventoryTests {
		inventory, err := InitInventory(example.configBytes)
		if err != example.expectedError {
			t.Error("Expected Error", example.expectedError, "But returned", err)
		}
		if err == nil {
			actualComponentNum := len(inventory.Components.GetAll())
			if actualComponentNum != example.expectedComponents {
				t.Error("The number of actual components and expected components do not match")
			}
		}
	}
}

func TestInitControlInfo(t *testing.T) {
	_, err := InitControlInfo()
	if err != nil {
		t.Error("InitInventory failed")
	}
}
