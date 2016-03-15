package renderers

import "testing"

type exportLinkTest struct {
	text     string
	location string
	expected string
}

var exportLinkTests = []exportLinkTest{
	{"test text", "location", "* [test text](location)  \n"},
	{"", "", "* []()  \n"},
}

func TestExportLink(t *testing.T) {
	for _, example := range exportLinkTests {
		actual := exportLink(example.text, example.location)
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}
