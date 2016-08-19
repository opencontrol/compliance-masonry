package components

import "testing"


type keyTest struct {
	input    string
	expected string
}

var keyTests = []keyTest{
	// Check that the key is extracted by using the local directory
	{".", "."},
	// Check that the key is extracted from the component directory
	{"system/component", "component"},
	// Check that the key is extracted by using the system directory
	{"system", "system"},
}

func TestGetKey(t *testing.T) {
	for _, example := range keyTests {
		actual := getKey(example.input)
		// Check that the actual key is the expected key
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}