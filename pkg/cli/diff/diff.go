package diff

import (
	"fmt"
	"io"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/spf13/cobra"
	"github.com/tg/gosortmap"
)

// NewCmdDiff provides a gap diff analysis.
func NewCmdDiff(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Compliance Diff Gap Analysis",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunDiff(out, cmd, args)
			clierrors.CheckError(err)
		},
	}
	cmd.Flags().StringP("opencontrol", "o", constants.DefaultOpenControlsFolder, "Set opencontrol directory")
	return cmd
}

// Runs diff when specified in cli
func RunDiff(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("certification type not specified")
	}

	if len(args) > 1 {
		return fmt.Errorf("too many arguments. expected only one certification type")
	}
	config := Config{
		Certification:  args[0],
		OpencontrolDir: cmd.Flag("opencontrol").Value.String(),
	}
	inventory, errs := ComputeGapAnalysis(config)
	if errs != nil && len(errs) > 0 {
		return clierrors.NewExitError(clierrors.NewMultiError(errs...).Error(), 1)
	}
	fmt.Fprintf(out, "\nNumber of missing controls: %d\n", len(inventory.MissingControlList))
	for _, standardAndControl := range sortmap.ByKey(inventory.MissingControlList) {
		fmt.Fprintf(out, "%s\n", standardAndControl.Key)
	}
	return nil
}
