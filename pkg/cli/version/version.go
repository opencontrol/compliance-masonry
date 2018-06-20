/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package version

import (
	"io"

	"github.com/opencontrol/compliance-masonry/version"
	"github.com/spf13/cobra"
)

// NewCmdVersion prints out the version
func NewCmdVersion(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display version",
		Run: func(cmd *cobra.Command, args []string) {
			version.PrintVersion()
		},
	}
	return cmd
}
