package docs

import (
	"fmt"
	"io"
	"os"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/pkg/cli/docs/gitbook"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/spf13/cobra"
)

// NewCmdDocs creates the compliance documentation.
func NewCmdDocs(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "docs",
		Short: "Create compliance documentation",
	}
	cmd.AddCommand(NewCmdDocsGitBook(out))
	return cmd
}

// NewCmdDocsGitBook creates the compliance documentation in Gitbook format.
func NewCmdDocsGitBook(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gitbook",
		Short: "Create compliance documentation in Gitbook format",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunGitBook(out, cmd, args)
			clierrors.CheckError(err)
		},
	}
	cmd.Flags().StringP("opencontrol", "o", constants.DefaultOpenControlsFolder, "Set opencontrol directory")
	cmd.Flags().StringP("export", "e", constants.DefaultExportsFolder, "Sets the export directory")
	cmd.Flags().StringP("markdown", "m", constants.DefaultMarkdownFolder, "Sets the markdown directory")
	return cmd
}

// RunGitBook generates GitBook style documentation when specified in cli
func RunGitBook(out io.Writer, cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("certification type not specified")
	}

	if len(args) > 1 {
		return fmt.Errorf("too many arguments. expected only one certification type")
	}
	config := gitbook.Config{
		Certification:  args[0],
		OpencontrolDir: cmd.Flag("opencontrol").Value.String(),
		ExportPath:     cmd.Flag("export").Value.String(),
		MarkdownPath:   cmd.Flag("markdown").Value.String(),
	}
	warning, errMessages := MakeGitbook(config)
	if warning != "" {
		fmt.Fprintf(out, "%v\n", warning)
	}
	if errMessages != nil && len(errMessages) > 0 {
		err := clierrors.NewMultiError(errMessages...)
		return clierrors.NewExitError(err.Error(), 1)
	}
	fmt.Fprintf(out, "%v\n", "New Gitbook Documentation Created")
	return nil

}

// MakeGitbook is the wrapper function that will create a gitbook for the specified certification.
func MakeGitbook(config gitbook.Config) (string, []error) {
	warning := ""
	certificationPath, err := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" {
		return warning, err
	}
	if _, err := os.Stat(config.MarkdownPath); os.IsNotExist(err) {
		warning = "Warning: markdown directory does not exist"
	}
	config.Certification = certificationPath
	if err := config.BuildGitbook(); err != nil {
		return warning, err
	}
	return warning, nil
}
