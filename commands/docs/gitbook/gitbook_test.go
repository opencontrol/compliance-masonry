package gitbook

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
	markdownPath      string
	expectedOutputDir string
}

var exportLinkTests = []exportLinkTest{
	// Check that text and location create the correct output
	{"test text", "location", "* [test text](location)\n"},
	// Check that an emtpy text and location create the correct output
	{"", "", "* []()\n"},
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
	// Check that Parentheses are replaced in multiple places
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
	// Check that the gitbook is correctly exported given the fixtures
	{
		filepath.Join("..", "..", "..", "fixtures", "opencontrol_fixtures"),
		filepath.Join("..", "..", "..", "fixtures", "opencontrol_fixtures", "certifications", "LATO.yaml"),
		"",
		filepath.Join("..", "..", "..", "fixtures", "exports_fixtures", "complete_export"),
	},
	// Check that the gitbook is correctly exported given the fixtures with markdowns
	{
		filepath.Join("..", "..", "..", "fixtures", "opencontrol_fixtures_with_markdown"),
		filepath.Join("..", "..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "certifications", "LATO.yaml"),
		filepath.Join("..", "..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
		filepath.Join("..", "..", "..", "fixtures", "exports_fixtures", "complete_export_with_markdown"),
	},
}

func TestBuildGitbook(t *testing.T) {
	for _, example := range buildGitbookTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config := Config{
			OpencontrolDir: example.inputDir,
			Certification:  example.certificationPath,
			ExportPath:     tempDir,
			MarkdownPath:   example.markdownPath,
		}

		config.BuildGitbook()
		// Loop through the expected output to verify it matches the actual output
		matches, _ := filepath.Glob(filepath.Join(example.expectedOutputDir, "*", "*"))
		for _, expectedfilePath := range matches {
			actualFilePath := strings.Replace(expectedfilePath, example.expectedOutputDir, tempDir, -1)
			expectedData, _ := ioutil.ReadFile(expectedfilePath)
			actualData, _ := ioutil.ReadFile(actualFilePath)
			actualDataString := strings.Replace(string(actualData), "\\", "/", -1)
			// Verify the expected text is the same as the actual text
			if string(expectedData) != actualDataString {
				t.Errorf("Expected (%s):\n`%s`\nActual:\n`%s`", expectedfilePath, string(expectedData), string(actualData))
			}
		}
	}
}
