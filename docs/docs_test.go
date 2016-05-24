package docs_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/opencontrol/compliance-masonry/docs"
	"github.com/opencontrol/compliance-masonry/docx"
	"github.com/opencontrol/compliance-masonry/gitbook"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Doc Tests", func() {
	DescribeTable("BuildTemplateTest", func(config docx.Config, expectedMessages []string) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualMessages := BuildTemplate(config)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		Entry(
			"Check the template method returns an error message when no template is defined",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   "",
				ExportPath:     "",
			},
			[]string{"Error: No Template Supplied"},
		),

		Entry(
			"Check the template method returns an error message when no template does not exist",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   "fake",
				ExportPath:     "",
			},
			[]string{"Error: Template does not exist"},
		),

		Entry(
			"Check that template is created when inputs are correct",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   filepath.Join("..", "fixtures", "template_fixtures", "test.docx"),
				ExportPath:     "",
			},
			[]string{"New Docx Created"},
		),
	)

	DescribeTable("BuildGitbookTest", func(config gitbook.Config, expectedMessages []string) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualMessages := MakeGitbook(config)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		Entry(
			"Check that the gitbook is correctly exported given the fixtures",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures"),
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			[]string{"Warning: markdown directory does not exist", "New Gitbook Documentation Created"},
		),

		Entry(
			"Check that there is an error when there is no opencontrol dir",
			gitbook.Config{
				OpencontrolDir: "",
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			[]string{"Error: `certifications` directory does exist"},
		),

		Entry(
			"Check that gitbook is created with markdowns",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LATO",
				MarkdownPath:   filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			[]string{"New Gitbook Documentation Created"},
		),

		Entry(
			"Check that thre is an error returned when the certification does not exist",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LAT",
				MarkdownPath:   filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			[]string{fmt.Sprintf("Error: `%s` does not exist\nUse one of the following:", filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown", "certifications", "LAT.yaml")), "`LATO`"},
		),

		Entry(
			"Check that error is returned when certification argument is not present",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "",
				MarkdownPath:   filepath.Join("..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns/"),
			},
			[]string{"Error: Missing Certification Argument"},
		),
	)
})
