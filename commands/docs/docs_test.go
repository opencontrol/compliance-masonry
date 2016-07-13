package docs_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/opencontrol/compliance-masonry/commands/docs"
	"github.com/opencontrol/compliance-masonry/commands/docs/docx"
	"github.com/opencontrol/compliance-masonry/commands/docs/gitbook"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Doc Tests", func() {
	DescribeTable("BuildTemplateTest", func(config docx.Config, expectedMessages error) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualMessages := BuildTemplate(config)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		Entry(
			"Check the template method returns an error message when no template is defined",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   "",
				ExportPath:     "",
			},
			errors.New("Error: No Template Supplied"),
		),

		Entry(
			"Check the template method returns an error message when no template does not exist",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   "fake",
				ExportPath:     "",
			},
			errors.New("Error: Template does not exist"),
		),

		Entry(
			"Check that template is created when inputs are correct",
			docx.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures"),
				TemplatePath:   filepath.Join("..", "..", "fixtures", "template_fixtures", "test.docx"),
				ExportPath:     "",
			},
			nil,
		),
	)

	DescribeTable("BuildGitbookTest", func(config gitbook.Config, expectedWarning string, expectedMessages []error) {
		tempDir, _ := ioutil.TempDir("", "example")
		defer os.RemoveAll(tempDir)
		config.ExportPath = tempDir
		actualWarning, actualMessages := MakeGitbook(config)
		assert.Equal(GinkgoT(), expectedWarning, actualWarning)
		assert.Equal(GinkgoT(), expectedMessages, actualMessages)
	},
		Entry(
			"Check that the gitbook is correctly exported given the fixtures",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures"),
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			"Warning: markdown directory does not exist",
			nil,
		),

		Entry(
			"Check that there is an error when there is no opencontrol dir",
			gitbook.Config{
				OpencontrolDir: "",
				Certification:  "LATO",
				MarkdownPath:   "",
			},
			"",
			[]error{errors.New("Error: `certifications` directory does exist")},
		),

		Entry(
			"Check that gitbook is created with markdowns",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LATO",
				MarkdownPath:   filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			"",
			nil,
		),

		Entry(
			"Check that thre is an error returned when the certification does not exist",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LAT",
				MarkdownPath:   filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			"",
			[]error{fmt.Errorf("Error: `%s` does not exist\nUse one of the following:\nLATO", filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "certifications", "LAT.yaml"))},
		),

		Entry(
			"Check that error is returned when certification argument is not present",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "",
				MarkdownPath:   filepath.Join("..", "..", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns/"),
			},
			"",
			[]error{errors.New("Error: Missing Certification Argument")},
		),
	)
})
