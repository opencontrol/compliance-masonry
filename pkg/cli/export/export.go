package export

import (
	"fmt"
	"io"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/spf13/cobra"
)

// docxtemplater boolean variable
var docxtemplater bool

// flatten boolean variable
var flattenFlag bool

// keys boolen flag
var keysFlag bool

// NewCmdExport exports to consolidated output
func NewCmdExport(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export",
		Short: "Export to consolidated output",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunExport(out, cmd, args)
			clierrors.CheckError(err)
		},
	}
	cmd.Flags().StringP("opencontrol", "o", constants.DefaultDestination, "Set opencontrol directory")
	cmd.Flags().StringP("dest", "d", constants.DefaultJSONFile, "Destination file for output")
	cmd.Flags().BoolVarP(&flattenFlag, "flatten", "n", false, "Flatten results file")
	cmd.Flags().StringP("format", "f", constants.DefaultOutputFormat, "Output format for destination file")
	cmd.Flags().BoolVarP(&keysFlag, "keys", "k", false, "Keys to use when processing arrays while flattening")
	cmd.Flags().BoolVarP(&docxtemplater, "docxtemplater", "x", false, "Use docxtemplater format")
	cmd.Flags().StringP("separator", "s", constants.DefaultKeySeparator, "Separator to use when flattening keys")
	return cmd
}

// RunExport runs export when specified in cli
func RunExport(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("certification type not specified")
	}

	// read parms
	parmOpencontrols := cmd.Flag("opencontrols").Value.String()
	parmDestination := cmd.Flag("destination").Value.String()
	parmOutputFormat := cmd.Flag("format").Value.String()
	parmFlatten := false
	parmInferKeys := false
	parmDocxtemplater := false
	parmKeySeparator := cmd.Flag("separator").Value.String()

	// convert to enum
	outputFormat, err := ToOutputFormat(parmOutputFormat)
	if err != nil {
		return clierrors.NewExitError(err.Error(), 1)
	}

	// --docxtemplater always forces --flatten
	if parmDocxtemplater {
		parmFlatten = true
	}

	// construct args
	config := Config{
		Certification:   args[0],
		OpencontrolDir:  parmOpencontrols,
		DestinationFile: parmDestination,
		OutputFormat:    outputFormat,
		Flatten:         parmFlatten,
		InferKeys:       parmInferKeys,
		Docxtemplater:   parmDocxtemplater,
		KeySeparator:    parmKeySeparator,
	}

	// invoke command
	errs := Export(config)
	if errs != nil && len(errs) > 0 {
		err := clierrors.NewMultiError(errs...)
		return clierrors.NewExitError(err.Error(), 1)
	}
	return nil
}
