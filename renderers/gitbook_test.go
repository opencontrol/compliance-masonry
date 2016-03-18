package renderers

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type exportLinkTest struct {
	text     string
	location string
	expected string
}

type replaceParenthesesTest struct {
	text     string
	expected string
}

type buildGitbookTest struct {
	inputDir          string
	certificationPath string
	expectedOutputDir string
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

var buildGitbookTests = []buildGitbookTest{
	{
		"../fixtures/opencontrol_fixtures/",
		"../fixtures/opencontrol_fixtures/certifications/LATO.yaml",
		"../fixtures/exports_fixtures/complete_export",
	},
}

func TestBuildGitbook(t *testing.T) {
	for _, example := range buildGitbookTests {
		tempDir, _ := ioutil.TempDir("", "example")

		defer os.RemoveAll(tempDir)
		BuildGitbook(example.inputDir, example.certificationPath, tempDir)

		matches, _ := filepath.Glob(filepath.Join(example.expectedOutputDir, "*", "*"))
		for _, expectedfilePath := range matches {
			actualFilePath := strings.Replace(expectedfilePath, example.expectedOutputDir, tempDir, -1)
			expectedData, _ := ioutil.ReadFile(expectedfilePath)
			actualData, _ := ioutil.ReadFile(actualFilePath)
			if string(expectedData) != string(actualData) {
				t.Errorf("Expected: `%s`,\n Actual: `%s`", string(expectedData), string(actualData))
			}
		}
	}
}
