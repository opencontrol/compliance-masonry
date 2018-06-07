package get

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/pkg/cli/get/resources"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/pkg/lib/opencontrol"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"github.com/spf13/cobra"
)

// NewCmdGet gets all the compliance dependencies
func NewCmdGet(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Install compliance dependencies",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGet(out, cmd)
			clierrors.CheckError(err)
		},
	}
	cmd.Flags().StringP("config", "c", constants.DefaultConfigYaml, "Location of system-level yaml configuration file")
	cmd.Flags().StringP("dest", "d", constants.DefaultDestination, "Location to download the compliance repositories")
	return cmd
}

// RunGet runs get when specified in cli
func RunGet(out io.Writer, cmd *cobra.Command) error {
	f := fs.OSUtil{}
	config := cmd.Flag("config").Value.String()
	configBytes, err := f.OpenAndReadFile(config)
	if err != nil {
		fmt.Fprintf(out, "%v\n", err.Error())
		os.Exit(1)
	}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(out, "%v\n", err.Error())
		os.Exit(1)
	}
	destination := filepath.Join(wd, cmd.Flag("dest").Value.String())
	err = Get(destination, configBytes)
	if err != nil {
		return clierrors.NewExitError(err.Error(), 1)
	}
	fmt.Fprintf(out, "%v\n", "Compliance Dependencies Installed")
	return nil
}

// Get will retrieve all of the resources for the schemas and the resources for all the dependent schemas.
func Get(destination string, configData []byte) error {
	// Check the data.
	if configData == nil || len(configData) == 0 {
		return common.ErrNoDataToParse
	}
	// Parse it.
	parser := opencontrol.YAMLParser{}
	configSchema, err := parser.Parse(configData)
	if err != nil {
		return err
	}
	// Get Resources
	getter := resources.NewVCSAndLocalGetter(parser)
	err = resources.GetResources("", destination, configSchema, getter)
	if err != nil {
		return err
	}
	return nil
}
