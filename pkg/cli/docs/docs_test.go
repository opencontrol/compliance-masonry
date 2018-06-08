package docs_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	. "github.com/opencontrol/compliance-masonry/pkg/cli/docs"
	"github.com/opencontrol/compliance-masonry/pkg/cli/docs/gitbook"
	"github.com/opencontrol/compliance-masonry/pkg/tests"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Doc Tests", func() {
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
				OpencontrolDir: filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
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
				OpencontrolDir: filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LATO",
				MarkdownPath:   filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			"",
			nil,
		),

		Entry(
			"Check that thre is an error returned when the certification does not exist",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "LAT",
				MarkdownPath:   filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns"),
			},
			"",
			[]error{fmt.Errorf("Error: `%s` does not exist\nUse one of the following:\nLATO", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown", "certifications", "LAT.yaml"))},
		),

		Entry(
			"Check that error is returned when certification argument is not present",
			gitbook.Config{
				OpencontrolDir: filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown"),
				Certification:  "",
				MarkdownPath:   filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns/"),
			},
			"",
			[]error{errors.New("Error: Missing Certification Argument")},
		),
	)
	Describe("Base Docs Commands", func() {
		Describe("When the CLI is run with the docs command", func() {
			It("should list the available doc commands", func() {
				output := masonry_test.Masonry("docs", "")
				Eventually(output.Out.Contents).Should(ContainSubstring("Create compliance documentation in Gitbook format"))
			})
		})
	})

	Describe("Gitbook Docs Commands", func() {

		var exportTempDir string
		BeforeEach(func() {
			exportTempDir, _ = ioutil.TempDir("", "exports")
		})

		Describe("Gitbook Commands", func() {
			Describe("When the CLI is run with the `docs gitbook` command", func() {
				It("should let the user know that they have not described a certification and show how to use the command", func() {
					output := masonry_test.Masonry("docs", "gitbook")
					Eventually(output.Err.Contents).Should(ContainSubstring("certification type not specified\n"))
				})
			})

			Describe("When the CLI is run with the `docs gitbook` command without opencontrols dir", func() {
				It("should let the user know that there is no opencontrols/certifications directory", func() {
					output := masonry_test.Masonry("docs", "gitbook", "LATO")
					Eventually(output.Err.Contents).Should(ContainSubstring("Error: `" + filepath.Join("opencontrols", "certifications") + "` directory does exist\n"))
				})
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification and no markdown", func() {
			It("should create the documentation but warn users that there is no markdown dir", func() {
				output := masonry_test.Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures"),
					"-m", "sdfds").Wait(1 * time.Second)
				Eventually(output.Out.Contents).Should(ContainSubstring("Warning: markdown directory does not exist\n"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created\n"))
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification", func() {
			It("should create the documentation without warning the user", func() {
				exportTempDir, _ := ioutil.TempDir("", "exports")
				output := masonry_test.Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown"),
					"-m", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures_with_markdown", "markdowns")).Wait(1 * time.Second)
				Eventually(output.Out.Contents).ShouldNot(ContainSubstring("Warning: markdown directory does not exist\n"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created\n"))
			})
		})
		AfterEach(func() {
			_ = os.RemoveAll(exportTempDir)
		})
	})
})
