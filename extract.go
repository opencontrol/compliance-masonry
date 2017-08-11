package main

import (
	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry/commands/extract"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

const (
	extractCommandName  = "extract"
	extractCommandUsage = "Extract as JSON output"
)

var (
	extractCommandAliases = []string{"x"}
	extractCommandFlags   = []cli.Flag{
		cli.StringFlag{
			Name:  "opencontrols, o",
			Value: constants.DefaultDestination,
			Usage: "Set opencontrols directory",
		},
		cli.StringFlag{
			Name:  "json, j",
			Value: constants.DefaultJsonFile,
			Usage: "JSON file for output",
		},
	}
	extractCommand = cli.Command{
		Name:    extractCommandName,
		Aliases: extractCommandAliases,
		Usage:   extractCommandUsage,
		Flags:   extractCommandFlags,
		Action:  extractCommandAction,
	}
)

func extractCommandAction(c *cli.Context) error {
	// read parms
	opencontrolsDir := c.String("opencontrols")
	jsonFile := c.String("json")

	// construct args
	config := extract.Config{
		Certification:  c.Args().First(),
		OpencontrolDir: opencontrolsDir,
		JsonFile:       jsonFile,
	}

	// invoke command
	errs := extract.Extract(config)
	if errs != nil && len(errs) > 0 {
		err := cli.NewMultiError(errs...)
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
