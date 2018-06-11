package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/compliance-masonry/pkg/tests"
)

var usage = `
Usage:
  masonry [flags]
  masonry [command]

Available Commands:
  diff        Compliance Diff Gap Analysis
  docs        Create compliance documentation
  export      Export to consolidated output
  get         Install compliance dependencies
  help        Help about any command

Flags:
  -h, --help      help for masonry
      --verbose   Run with verbosity
  -v, --version   Print the version
`

var _ = Describe("Masonry CLI", func() {
	Describe("When the CLI is run with no commands", func() {
		It("should list the available commands", func() {
			output := masonry_test.Masonry()
			Eventually(output.Out.Contents).Should(ContainSubstring(usage))
		})
	})
})
