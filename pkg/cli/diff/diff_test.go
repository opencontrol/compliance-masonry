/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package diff_test

import (
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opencontrol/compliance-masonry/pkg/tests"
)

var _ = Describe("Diff Commands", func() {
	Describe("When the diff command is run", func() {
		It("should let the user know that they have not described a certification and show how to use the command", func() {
			output := masonry_test.Masonry("diff")
			Eventually(output.Err.Contents).Should(ContainSubstring("certification type not specified\n"))
		})
	})
	Describe("When the CLI is run with the `diff` command without opencontrols dir", func() {
		It("should let the user know that there is no opencontrols/certifications directory", func() {
			output := masonry_test.Masonry("diff", "LATO")
			Eventually(output.Err.Contents).Should(ContainSubstring("Error: `" + filepath.Join("opencontrols", "certifications") + "` directory does exist\n"))
		})
	})
	Describe("When the CLI is run with the `diff` command with a certification", func() {
		It("should print the number of missing controls", func() {
			output := masonry_test.Masonry(
				"diff", "LATO",
				"-o", filepath.Join("..", "..", "..", "test", "fixtures", "opencontrol_fixtures")).Wait(1 * time.Second)
			Eventually(output.Out.Contents).Should(ContainSubstring("Number of missing controls:"))
		})
	})
})
