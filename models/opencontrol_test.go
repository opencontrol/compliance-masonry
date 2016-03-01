package models

import "testing"

type keyTest struct {
	input    string
	expected string
}

var keyTests = []keyTest{
	{".", "."},
	{"system/component", "component"},
	{"system", "system"},
}

func TestGetKey(t *testing.T) {
	for _, example := range keyTests {
		actual := getKey(example.input)
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}
