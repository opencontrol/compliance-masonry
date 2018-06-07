package diff_test

import (
	"bufio"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Diff Commands", func() {
	Describe("When the diff command is run", func() {
		It("should let the user know that they have not described a certification and show how to use the command", func() {
			output := Masonry("diff")
			Eventually(output.Err.Contents).Should(ContainSubstring("certification type not specified\n"))
		})
	})
	Describe("When the CLI is run with the `diff` command without opencontrols dir", func() {
		It("should let the user know that there is no opencontrols/certifications directory", func() {
			output := Masonry("diff", "LATO")
			Eventually(output.Err.Contents).Should(ContainSubstring("Error: `" + filepath.Join("opencontrols", "certifications") + "` directory does exist\n"))
		})
	})
	Describe("When the CLI is run with the `diff` command with a certification", func() {
		It("should print the number of missing controls", func() {
			output := Masonry(
				"diff", "LATO",
				"-o", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures")).Wait(1 * time.Second)
			Eventually(output.Out.Contents).Should(ContainSubstring("Number of missing controls:"))
		})
	})
})

func Masonry(args ...string) *Session {
	path, err := Build("github.com/opencontrol/compliance-masonry/cmd/compliance-masonry")
	Expect(err).NotTo(HaveOccurred())
	cmd := exec.Command(path, args...)
	stdin, err := cmd.StdinPipe()
	Expect(err).ToNot(HaveOccurred())
	buffer := bufio.NewWriter(stdin)
	_, _ = buffer.WriteString(strings.Join(args, " "))
	_ = buffer.Flush()
	session, err := Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	return session
}
