/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/compliance-masonry/pkg/tests"
)

var usage = `
Usage:
  masonry [global-options] COMMAND [command-options]

Commands:
`

var _ = Describe("Masonry CLI", func() {
	Describe("When the CLI is run with no commands", func() {
		It("should list the available commands", func() {
			output := masonry_test.Masonry()
			Eventually(output.Out.Contents).Should(ContainSubstring(usage))
		})
	})
})
