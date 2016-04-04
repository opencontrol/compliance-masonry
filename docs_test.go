package main_test

import (
	"io/ioutil"
	"os"

	. "github.com/opencontrol/compliance-masonry-go"
	"github.com/opencontrol/compliance-masonry-go/docx"
	"github.com/opencontrol/compliance-masonry-go/gitbook"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Doc Tests", func() {
	table.DescribeTable("BuildTemplateTest", func(config docx.Config, expectedMessages []string) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualMessages := BuildTemplate(&config)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		table.Entry(
			"Check the template method returns an error message when no template is defined",
			docx.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures/",
				TemplatePath:   "",
				ExportPath:     "",
			},
			[]string{"Error: No Template Supplied"},
		),

		table.Entry(
			"Check the template method returns an error message when no template does not exist",
			docx.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures/",
				TemplatePath:   "fake",
				ExportPath:     "",
			},
			[]string{"Error: Template does not exist"},
		),

		table.Entry(
			"Check that template is created when inputs are correct",
			docx.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures/",
				TemplatePath:   "fixtures/template_fixtures/test.docx",
				ExportPath:     "",
			},
			[]string{"New Docx Created"},
		),
	)

	table.DescribeTable("BuildGitbookTest", func(config gitbook.Config, expectedMessages []string) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualMessages := MakeGitbook(&config)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		table.Entry(
			"Check that the gitbook is correctly exported given the fixtures",
			gitbook.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures/",
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			[]string{"Warning: markdown directory does not exist", "New Gitbook Documentation Created"},
		),

		table.Entry(
			"Check that there is an error when there is no opencontrol dir",
			gitbook.Config{
				OpencontrolDir: "",
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			[]string{"Error: `opencontrols/certifications` directory does exist"},
		),

		table.Entry(
			"Check that gitbook is created with markdowns",
			gitbook.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
				Certification:  "LATO",
				MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
			},
			[]string{"New Gitbook Documentation Created"},
		),

		table.Entry(
			"Check that thre is an error returned when the certification does not exist",
			gitbook.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
				Certification:  "LAT",
				MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
			},
			[]string{"Error: `fixtures/opencontrol_fixtures_with_markdown/certifications/LAT.yaml` does not exist\nUse one of the following:", "`LATO`"},
		),

		table.Entry(
			"Check that error is returned when certification argument is not present",
			gitbook.Config{
				OpencontrolDir: "fixtures/opencontrol_fixtures_with_markdown/",
				Certification:  "",
				MarkdownPath:   "fixtures/opencontrol_fixtures_with_markdown/markdowns/",
			},
			[]string{"Error: Missing Certification Argument"},
		),
	)
})
