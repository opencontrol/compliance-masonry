package main_test

import (
	"bufio"
	"os/exec"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var usage = `
Usage:
  compliance-masonry [flags]
  compliance-masonry [command]

Available Commands:
  diff        Compliance Diff Gap Analysis
  docs        Create compliance documentation
  get         Install compliance dependencies
  help        Help about any command

Flags:
  -h, --help      help for compliance-masonry
      --verbose   Run with verbosity
  -v, --version   Print the version
`

var _ = Describe("Masonry CLI", func() {
	Describe("When the CLI is run with no commands", func() {
		It("should list the available commands", func() {
			output := Masonry()
			Eventually(output.Out.Contents).Should(ContainSubstring(usage))
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
