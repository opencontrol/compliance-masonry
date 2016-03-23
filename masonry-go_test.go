package main

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
			Eventually(output.Out.Contents).Should(ContainSubstring("init, i	Initialize Open Control documentation repository"))
			Eventually(output.Out.Contents).Should(ContainSubstring("get, g	Install compliance dependencies"))
			Eventually(output.Out.Contents).Should(ContainSubstring("docs, d	Create Documentation"))
			Eventually(output.Out.Contents).Should(ContainSubstring("help, h	Shows a list of commands or help for one command"))
		})
	})

	Describe("Docs Commands No Output", func() {
		Describe("When the CLI is run with the docs command", func() {
			It("should list the available doc commands", func() {
				output := Masonry("docs", "")
				Eventually(output.Out.Contents).Should(ContainSubstring("gitbook, g	Create Gitbook Documentation"))
				Eventually(output.Out.Contents).Should(ContainSubstring("help, h	Shows a list of commands or help for one command"))
			})
		})

		Describe("Gitbook Commands", func() {
			Describe("When the CLI is run with the `docs gitbook` command", func() {
				It("should let the user know that they have not described a certification and show how to use the command", func() {
					output := Masonry("docs", "gitbook")
					Eventually(output.Out.Contents).Should(ContainSubstring("Error: New Missing Certification Argument"))
					Eventually(output.Out.Contents).Should(ContainSubstring("Usage: masonry-go docs gitbook FedRAMP-low"))
				})
			})

			Describe("When the CLI is run with the `docs gitbook` command without opencontrols dir", func() {
				It("should let the user know that there is no opencontrols/certifications directory", func() {
					output := Masonry("docs", "gitbook", "LATO")
					Eventually(output.Out.Contents).Should(ContainSubstring("Error: `opencontrols/certifications` directory does exist"))
				})
			})
		})

		Describe("Docs Commands Output", func() {

			var exportTempDir string
			BeforeEach(func() {
				exportTempDir, _ = ioutil.TempDir("", "exports")

			})

			Describe("When the CLI is run with the `docs gitbook` command with a certification", func() {
				It("should create the documentation but warn users that there is no markdown dir", func() {
					output := Masonry(
						"docs", "gitbook", "LATO",
						"-e", exportTempDir,
						"-o", "./fixtures/opencontrol_fixtures/",
						"-m", "sdfds").Wait(1 * time.Second)
					Eventually(output.Out.Contents).Should(ContainSubstring("Warning: markdown directory does not exist"))
					Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created"))
					CompareDirs("fixtures/exports_fixtures/complete_export", exportTempDir)
				})
			})

			Describe("When the CLI is run with the `docs gitbook` command with a certification", func() {
				It("should create the documentation without warning the user", func() {
					exportTempDir, _ := ioutil.TempDir("", "exports")
					output := Masonry(
						"docs", "gitbook", "LATO",
						"-e", exportTempDir,
						"-o", "./fixtures/opencontrol_fixtures_with_markdown/",
						"-m", "./fixtures/opencontrol_fixtures_with_markdown/markdowns").Wait(1 * time.Second)
					Eventually(output.Out.Contents).ShouldNot(ContainSubstring("Warning: markdown directory does not exist"))
					Eventually(output.Out.Contents).Should(ContainSubstring("New Gitbook Documentation Created"))
					CompareDirs("fixtures/exports_fixtures/complete_export_with_markdown", exportTempDir)
				})
			})

			AfterEach(func() {
				os.RemoveAll(exportTempDir)
			})

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

func CompareDirs(expectedDir string, actualDir string) {
	matches, _ := filepath.Glob(filepath.Join(expectedDir, "*"))
	for _, expectedfilePath := range matches {
		actualFilePath := strings.Replace(expectedfilePath, expectedDir, actualDir, -1)
		Expect(actualFilePath).ToNot(Equal(expectedfilePath))
		expectedData, _ := ioutil.ReadFile(expectedfilePath)
		actualData, _ := ioutil.ReadFile(actualFilePath)
		Expect(string(actualData)).To(Equal(string(expectedData)))
	}
}
