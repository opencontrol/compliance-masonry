/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package masonry

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/pkg/cli/diff"
	"github.com/opencontrol/compliance-masonry/pkg/cli/docs"
	"github.com/opencontrol/compliance-masonry/pkg/cli/export"
	"github.com/opencontrol/compliance-masonry/pkg/cli/get"
	"github.com/opencontrol/compliance-masonry/pkg/cli/info"
	cliversion "github.com/opencontrol/compliance-masonry/pkg/cli/version"
	"github.com/opencontrol/compliance-masonry/version"
	"github.com/spf13/cobra"
)

// Verbose boolean for turning on/off verbosity
var Verbose bool

// Version for cli variable holding program version
var Version bool

func init() {
	log.SetOutput(ioutil.Discard)
}

// NewMasonryCommand Main Masonry command cli
// Add new commands/subcommands for new verbs in this function
func NewMasonryCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "masonry",
		Short: "OpenControl CLI Tool",
		Long: `Compliance Masonry is a command-line interface (CLI) that
allows users to construct certification documentation using
the OpenControl Schema`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := RunGlobalFlags(out, cmd)
			clierrors.CheckError(err)
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmds.SetUsageTemplate(usageTemplate)
	cmds.ResetFlags()
	// Global Options
	cmds.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Run with verbosity")
	cmds.PersistentFlags().BoolVarP(&Version, "version", "v", false, "Print the version")

	// Add new main commands here
	cmds.AddCommand(diff.NewCmdDiff(out))
	cmds.AddCommand(info.NewCmdInfo(out))
	cmds.AddCommand(docs.NewCmdDocs(out))
	cmds.AddCommand(export.NewCmdExport(out))
	cmds.AddCommand(get.NewCmdGet(out))
	cmds.AddCommand(cliversion.NewCmdVersion(out))

	disableFlagsInUseLine(cmds)

	return cmds
}

// RunGlobalFlags runs global options when specified in cli
func RunGlobalFlags(out io.Writer, cmd *cobra.Command) error {
	flagVersion := Version
	flagVerbose := Verbose

	if flagVerbose {
		log.SetOutput(os.Stderr)
		log.Println("Running with verbosity")
	}

	if flagVersion {
		version.PrintVersion()
	}

	return nil

}

// disableFlagsInUseLine do not add a `[flags]` to the end of the usage line.
func disableFlagsInUseLine(cmd *cobra.Command) {
	visitAll(cmd, func(cmds *cobra.Command) {
		cmds.DisableFlagsInUseLine = true
	})
}

// visitAll will traverse all commands from the root.
// This is different from the VisitAll of cobra.Command where only parents
// are checked.
func visitAll(cmds *cobra.Command, fn func(*cobra.Command)) {
	for _, cmd := range cmds.Commands() {
		visitAll(cmd, fn)
	}
	fn(cmds)
}
