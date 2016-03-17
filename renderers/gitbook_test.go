package renderers

import "testing"

type exportLinkTest struct {
	text     string
	location string
	expected string
}

type replaceParenthesesTest struct {
	text     string
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

var replaceParenthesesTests = []replaceParenthesesTest{
	{"NIST-800-53-(1).md", "NIST-800-53-1.md"},
	{"NIST-800-53-(1.md", "NIST-800-53-1.md"},
	{"NIST-()800()-53-1.md", "NIST-800-53-1.md"},
	{"NIST-()800-)53-(1).md", "NIST-800-53-1.md"},
}

func TestReplaceParentheses(t *testing.T) {
	for _, example := range replaceParenthesesTests {
		actual := replaceParentheses(example.text)
		if actual != example.expected {
			t.Errorf("Expected: `%s`, Actual: `%s`", example.expected, actual)
		}
	}
}
