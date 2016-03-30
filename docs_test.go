package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type buildGitbookTest struct {
	config           gitbookConfig
	expectedMessages []string
}

type buildTemplateTest struct {
	config           templateConfig
	expectedMessages []string
}

var buildTemplateTests = []buildTemplateTest{
	//Check the template method returns an error message when no template is defined
	{
		templateConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "LATO",
			templatePath:   "",
			exportPath:     "",
		},
		[]string{"Error: No Template Supplied"},
	},
	//Check the template method returns an error message when no template does not exist
	{
		templateConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "LATO",
			templatePath:   "fake",
			exportPath:     "",
		},
		[]string{"Error: Template does not exist"},
	},
	// Check the template method returns an error messages when certification doesn't exist
	{
		templateConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "",
			templatePath:   "fixtures/template_fixtures/test.docx",
			exportPath:     "",
		},
		[]string{"Error: Missing Certification Argument"},
	},
	// Check that template is created when inputs are correct
	{
		templateConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "LATO",
			templatePath:   "fixtures/template_fixtures/test.docx",
			exportPath:     "",
		},
		[]string{"Template Created"},
	},
}

func TestBuildTemplate(t *testing.T) {
	for _, example := range buildTemplateTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		example.config.exportPath = tempDir
		actualMessages := example.config.buildTemplate()
		assert.Equal(t, example.expectedMessages, actualMessages)
	}
}

var buildGitbookTests = []buildGitbookTest{
	//Check that the gitbook is correctly exported given the fixtures
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures/",
			certification:  "LATO",
			markdownPath:   "",
		},
		[]string{"Warning: markdown directory does not exist", "New Gitbook Documentation Created"},
	},
	{
		gitbookConfig{
			opencontrolDir: "",
			certification:  "LATO",
			markdownPath:   "",
		},
		[]string{"Error: `opencontrols/certifications` directory does exist"},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "LATO",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"New Gitbook Documentation Created"},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "LAT",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"Error: `fixtures/opencontrol_fixtures_with_markdown/certifications/LAT.yaml` does not exist\nUse one of the following:", "`LATO`"},
	},
	{
		gitbookConfig{
			opencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			certification:  "",
			markdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"Error: Missing Certification Argument"},
	},
}

func TestMakeGitbook(t *testing.T) {
	for _, example := range buildGitbookTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		example.config.exportPath = tempDir
		actualMessages := example.config.makeGitbook()
		assert.Equal(t, example.expectedMessages, actualMessages)
	}
}
