package main_test

import (
	"bufio"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Masonry CLI", func() {
	Describe("When the CLI is run with no commands", func() {
		It("should list the available commands", func() {
			output := Masonry("", "")
			Eventually(output.Out.Contents).Should(ContainSubstring("Install compliance dependencies"))
			Eventually(output.Out.Contents).Should(ContainSubstring("Create Documentation"))
		})
	})

	Describe("Base Docs Commands", func() {
		Describe("When the CLI is run with the docs command", func() {
			It("should list the available doc commands", func() {
				output := Masonry("docs", "")
				Eventually(output.Out.Contents).Should(ContainSubstring("Create Gitbook Documentation"))
				Eventually(output.Out.Contents).Should(ContainSubstring("Create Docx Documentation using a Template"))
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
					output := Masonry("docs", "gitbook")
					Eventually(output.Out.Contents).Should(ContainSubstring("Error: Missing Certification Argument"))
				})
			})

			Describe("When the CLI is run with the `docs gitbook` command without opencontrols dir", func() {
				It("should let the user know that there is no opencontrols/certifications directory", func() {
					output := Masonry("docs", "gitbook", "LATO")
					Eventually(output.Out.Contents).Should(ContainSubstring("Error: `opencontrols/certifications` directory does exist"))
				})
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification and no markdown", func() {
			It("should create the documentation but warn users that there is no markdown dir", func() {
				output := Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("fixtures", "opencontrol_fixtures"),
					"-m", "sdfds").Wait(1 * time.Second)
				Eventually(output.Out.Contents).Should(ContainSubstring("Warning: markdown directory does not exist"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created"))
			})
		})

		Describe("When the CLI is run with the `docs gitbook` command with a certification", func() {
			It("should create the documentation without warning the user", func() {
				exportTempDir, _ := ioutil.TempDir("", "exports")
				output := Masonry(
					"docs", "gitbook", "LATO",
					"-e", exportTempDir,
					"-o", filepath.Join("fixtures", "opencontrol_fixtures_with_markdown"),
					"-m", filepath.Join("fixtures", "opencontrol_fixtures_with_markdown", "markdowns")).Wait(1 * time.Second)
				Eventually(output.Out.Contents).ShouldNot(ContainSubstring("Warning: markdown directory does not exist"))
				Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created"))
			})
		})
		AfterEach(func() {
			os.RemoveAll(exportTempDir)
		})
	})

	Describe("Template Engine Commands", func() {

		var exportTempDir string
		BeforeEach(func() {
			exportTempDir, _ = ioutil.TempDir("", "exports")
		})

		Describe("When the docs docx command is run", func() {
			It("should warn the user that no template has been supplied", func() {
				output := Masonry("docs", "docx")
				Eventually(output.Out.Contents).Should(ContainSubstring("Error: No Template Supplied"))
			})
		})

		Describe("When the docs docx command is run with a none existent template", func() {
			It("should warn the user that no template does not exist", func() {
				output := Masonry("docs", "docx", "-t", "test")
				Eventually(output.Out.Contents).Should(ContainSubstring("Error: Template does not exist"))
			})
		})

		Describe("When the docs docx command is run with a corrupted template", func() {
			It("should return an error", func() {
				output := Masonry(
					"docs", "docx",
					"-o", filepath.Join("fixtures", "opencontrol_fixtures"),
					"-t", filepath.Join("fixtures", "template_fixtures", "test_corrupted.docx"),
					"-e", filepath.Join(exportTempDir, "export.docx"),
				)
				Eventually(output.Out.Contents).Should(ContainSubstring("Cannot Open File"))
			})
		})

		Describe("When the docs docx command is run with an existing template and certification", func() {
			It("should run the script", func() {
				output := Masonry(
					"docs", "docx",
					"-o", filepath.Join("fixtures", "opencontrol_fixtures"),
					"-t", filepath.Join("fixtures", "template_fixtures", "test.docx"),
					"-e", filepath.Join(exportTempDir, "export.docx"),
				)
				Eventually(output.Out.Contents).Should(ContainSubstring("New Docx Created"))
			})
		})

		AfterEach(func() {
			os.RemoveAll(exportTempDir)
		})

	})
})

func Masonry(args ...string) *Session {
	path, err := Build("github.com/opencontrol/compliance-masonry-go")
	Expect(err).NotTo(HaveOccurred())
	cmd := exec.Command(path, args...)
	stdin, err := cmd.StdinPipe()
	Expect(err).ToNot(HaveOccurred())
	buffer := bufio.NewWriter(stdin)
	buffer.WriteString(strings.Join(args, " "))
	buffer.Flush()
	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
