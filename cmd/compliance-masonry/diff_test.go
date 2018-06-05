package main

import (
	"bytes"
	"github.com/codegangsta/cli"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"path/filepath"
)

var _ = Describe("Diff", func() {
	var (
		app       *cli.App
		buffer    *bytes.Buffer
		errBuffer *bytes.Buffer
	)
	BeforeEach(func() {
		app = NewCLIApp()
		buffer = new(bytes.Buffer)
		app.Writer = buffer
		errBuffer = new(bytes.Buffer)
		app.ErrWriter = errBuffer
		cli.ErrWriter = errBuffer
		cli.OsExiter = func(code int) {}
	})
	Describe("Running the diff command", func() {
		It("should return an exit error of 1 when running without certification argument.", func() {
			err := app.Run([]string{app.Name, "diff"})
			assert.NotNil(GinkgoT(), err)
			if assert.IsType(GinkgoT(), new(cli.ExitError), err) {
				exitErr, _ := err.(*cli.ExitError)
				assert.Equal(GinkgoT(), 1, exitErr.ExitCode())
				assert.Contains(GinkgoT(), errBuffer.String(), "Error: Missing Certification Argument")
			}
		})
		It("should return return the number a missing controls when given a certification", func() {
			err := app.Run([]string{app.Name, "diff", "LATO", "-o", filepath.Join("fixtures", "opencontrol_fixtures")})
			assert.Nil(GinkgoT(), err)
			assert.Contains(GinkgoT(), buffer.String(), "Number of missing controls:")
		})
	})
})
