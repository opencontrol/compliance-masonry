/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package implementationstatus

import (
	"fmt"
	"io"

	"github.com/opencontrol/compliance-masonry/pkg/cli/clierrors"
	"github.com/opencontrol/compliance-masonry/pkg/lib"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/spf13/cobra"
	"github.com/tg/gosortmap"
)

// NewCmdImplementationStatus allows you to query for implementation_status.
func NewCmdImplementationStatus(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "implementationstatus",
		Short: "Compliance implementation status search",
		Run: func(cmd *cobra.Command, args []string) {
			err := RunImplementationStatus(out, cmd, args)
			clierrors.CheckError(err)
		},
	}
	cmd.Flags().StringP("opencontrol", "o", constants.DefaultOpenControlsFolder, "Set opencontrol directory")
	cmd.Flags().StringP("implementation_status", "i", constants.DefaultImplementationStatus, "implementation_status to search for")
	return cmd
}

// RunImplementationStatus allows you to query for implementation_status when specified in cli
func RunImplementationStatus(out io.Writer, cmd *cobra.Command, args []string) error {
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
	inventory, errs := FindImplementationStatus(config, cmd.Flag("implementation_status").Value.String())
	if errs != nil && len(errs) > 0 {
		return clierrors.NewExitError(clierrors.NewMultiError(errs...).Error(), 1)
	}
	fmt.Fprintf(out, "# Components with implementation_status: %s\n", cmd.Flag("implementation_status").Value.String())
	for _, control := range sortmap.ByKey(inventory.SatisfiesMap) {
		fmt.Fprintf(out, "%s\n", control.Key)
	}
	return nil
}

// Config contains the settings
type Config struct {
	Certification  string
	OpencontrolDir string
}

type ComponentsInventory struct {
	common.Workspace
	ComponentList []common.Component
	SatisfiesMap  map[string]common.Satisfies
}

func FindImplementationStatus(config Config, statustype string) (ComponentsInventory, []error) {
	// Initialize inventory with certification
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if certificationPath == "" || errs != nil {
		return ComponentsInventory{}, errs
	}
	workspace, errs := lib.LoadData(config.OpencontrolDir, certificationPath)
	if errs != nil {
		return ComponentsInventory{}, errs
	}
	i := ComponentsInventory{
		Workspace:    workspace,
		SatisfiesMap: make(map[string]common.Satisfies),
	}

	i.ComponentList = i.GetAllComponents()
	if i.GetCertification() == nil || i.ComponentList == nil {
		return ComponentsInventory{}, []error{fmt.Errorf("Unable to load data in %s for certification %s", config.OpencontrolDir, config.Certification)}
	}

	for _, component := range i.ComponentList {
		for _, satisfiedControl := range component.GetAllSatisfies() {
			for _, status := range satisfiedControl.GetImplementationStatuses() {
				if status == statustype {
					key := component.GetName() + "@" + satisfiedControl.GetControlKey()
					if _, exists := i.SatisfiesMap[key]; !exists {
						i.SatisfiesMap[key] = satisfiedControl
					}
				}
			}
		}
	}
	return i, nil
}
