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
	"github.com/opencontrol/compliance-masonry/pkg/cli/get"
	"github.com/opencontrol/compliance-masonry/version"
	"github.com/spf13/cobra"
)

// Verbose boolean for turning on/off verbosity
var Verbose bool

// Version for cli variable holding program version
var Version bool

// NewMasonryCommand Main Masonry command cli
// Add new commands/subcommands for new verbs in this function
func NewMasonryCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "masonry",
		Short: "OpenControl CLI Tool",
		Long: `Compliance Masonry is a command-line interface (CLI) that
allows users to construct certification documentation using
the OpenControl Schema`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				os.Exit(0)
			}
			err := RunGlobalFlags(out, cmd)
			clierrors.CheckError(err)
		},
	}

	cmds.ResetFlags()

	cmds.AddCommand(diff.NewCmdDiff(out))
	cmds.AddCommand(docs.NewCmdDocs(out))
	cmds.AddCommand(get.NewCmdGet(out))

	// Global Options
	cmds.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "Run with verbosity")
	cmds.PersistentFlags().BoolVarP(&Version, "version", "v", false, "Print the version")

	return cmds
}

// RunGlobalFlags runs global options when specified in cli
func RunGlobalFlags(out io.Writer, cmd *cobra.Command) error {

	var flagVersion = cmd.Flag("version").Value.String()
	var flagVerbose = cmd.Flag("verbose").Value.String()

	log.SetOutput(ioutil.Discard)
	if flagVerbose == "true" {
		log.SetOutput(os.Stderr)
		log.Println("Running with verbosity")
	}

	if flagVersion == "true" {
		fmt.Printf("compliance-masonry: version %s\n", version.Version)
	}

	return nil

}
