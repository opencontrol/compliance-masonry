package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/opencontrol/compliance-masonry-go/docx"
	"github.com/opencontrol/compliance-masonry-go/gitbook"
	"github.com/stretchr/testify/assert"
)

type buildGitbookTest struct {
	config           gitbook.Config
	expectedMessages []string
}

type buildTemplateTest struct {
	config           docx.Config
	expectedMessages []string
}

var buildTemplateTests = []buildTemplateTest{
	//Check the template method returns an error message when no template is defined
	{
		docx.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures/",
			TemplatePath:   "",
			ExportPath:     "",
		},
		[]string{"Error: No Template Supplied"},
	},
	//Check the template method returns an error message when no template does not exist
	{
		docx.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures/",
			TemplatePath:   "fake",
			ExportPath:     "",
		},
		[]string{"Error: Template does not exist"},
	},
	// Check that template is created when inputs are correct
	{
		docx.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures/",
			TemplatePath:   "fixtures/template_fixtures/test.docx",
			ExportPath:     "",
		},
		[]string{"Template Created"},
	},
}

func TestBuildTemplate(t *testing.T) {
	for _, example := range buildTemplateTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		example.config.ExportPath = tempDir
		actualMessages := buildTemplate(&example.config)
		assert.Equal(t, example.expectedMessages, actualMessages)
	}
}

var buildGitbookTests = []buildGitbookTest{
	//Check that the gitbook is correctly exported given the fixtures
	{
		gitbook.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures/",
			Certification:  "LATO",
			MarkdownPath:   "",
		},
		[]string{"Warning: markdown directory does not exist", "New Gitbook Documentation Created"},
	},
	{
		gitbook.Config{
			OpencontrolDir: "",
			Certification:  "LATO",
			MarkdownPath:   "",
		},
		[]string{"Error: `opencontrols/certifications` directory does exist"},
	},
	{
		gitbook.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			Certification:  "LATO",
			MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"New Gitbook Documentation Created"},
	},
	{
		gitbook.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			Certification:  "LAT",
			MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"Error: `fixtures/opencontrol_fixtures_with_markdown/certifications/LAT.yaml` does not exist\nUse one of the following:", "`LATO`"},
	},
	{
		gitbook.Config{
			OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
			Certification:  "",
			MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
		},
		[]string{"Error: Missing Certification Argument"},
	},
}

func TestMakeGitbook(t *testing.T) {
	for _, example := range buildGitbookTests {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		example.config.ExportPath = tempDir
		actualMessages := makeGitbook(&example.config)
		assert.Equal(t, example.expectedMessages, actualMessages)
	}
}
