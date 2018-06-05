package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry/commands/diff"
	"github.com/tg/gosortmap"
)

const (
	diffCommandName  = "diff"
	diffCommandUsage = "Compute Gap Analysis"
)

var (
	diffCommandAliases = []string{"d"}
	diffCommandFlags   = []cli.Flag{
		cli.StringFlag{
			Name:        "opencontrols, o",
			Value:       "opencontrols",
			Usage:       "Set opencontrols directory",
			Destination: &opencontrolDir,
		},
	}
	diffCommand = cli.Command{
		Name:    diffCommandName,
		Aliases: diffCommandAliases,
		Usage:   diffCommandUsage,
		Flags:   diffCommandFlags,
		Action:  diffCommandAction,
	}
)

func diffCommandAction(c *cli.Context) error {
	config := diff.Config{
		Certification:  c.Args().First(),
		OpencontrolDir: opencontrolDir,
	}
	inventory, errs := diff.ComputeGapAnalysis(config)
	if errs != nil && len(errs) > 0 {
		return cli.NewExitError(cli.NewMultiError(errs...).Error(), 1)
	}
	fmt.Fprintf(c.App.Writer, "\nNumber of missing controls: %d\n", len(inventory.MissingControlList))
	for _, standardAndControl := range sortmap.ByKey(inventory.MissingControlList) {
		fmt.Fprintf(c.App.Writer, "%s\n", standardAndControl.Key)
	}
	return nil
}
