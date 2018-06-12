package masonry

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/pkg/cli/diff"
	"github.com/opencontrol/compliance-masonry/pkg/cli/docs"
	"github.com/opencontrol/compliance-masonry/pkg/cli/export"
	"github.com/opencontrol/compliance-masonry/pkg/cli/get"
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
	}

	cmds.ResetFlags()
	// Global Options
	cmds.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Run with verbosity")
	cmds.PersistentFlags().BoolVarP(&Version, "version", "v", false, "Print the version")

	// Add new main commands here
	cmds.AddCommand(diff.NewCmdDiff(out))
	cmds.AddCommand(docs.NewCmdDocs(out))
	cmds.AddCommand(export.NewCmdExport(out))
	cmds.AddCommand(get.NewCmdGet(out))

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
		fmt.Printf("compliance-masonry: version %s\n", version.Version)
		os.Exit(0)
	}

	return nil

}
