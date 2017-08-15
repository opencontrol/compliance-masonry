package main

import (
	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry/commands/export"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

const (
	exportCommandName  = "export"
	exportCommandUsage = "Export to consolidated output (JSON/YAML)"
)

var (
	exportCommandAliases = []string{"x"}
	exportCommandFlags   = []cli.Flag{
		cli.StringFlag{
			Name:  "opencontrols, o",
			Value: constants.DefaultDestination,
			Usage: "Set opencontrols directory",
		},
		cli.StringFlag{
			Name:  "destination, d",
			Value: constants.DefaultJSONFile,
			Usage: "Destination file for output",
		},
		cli.StringFlag{
			Name:  "format, f",
			Value: constants.DefaultOutputFormat,
			Usage: "Output format for destination file",
		},
		cli.BoolFlag{
			Name:  "flatten, n",
			Usage: "Flatten result file",
		},
		cli.BoolFlag{
			Name:  "infer-keys, k",
			Usage: "Infer keys to use when processing arrays while flattening",
		},
		cli.BoolFlag{
			Name:  "docxtemplater, x",
			Usage: "Use docxtemplater format",
		},
		cli.StringFlag{
			Name:  "key-separator, s",
			Value: constants.DefaultKeySeparator,
			Usage: "Separator to use when flattening keys",
		},
	}
	exportCommand = cli.Command{
		Name:    exportCommandName,
		Aliases: exportCommandAliases,
		Usage:   exportCommandUsage,
		Flags:   exportCommandFlags,
		Action:  exportCommandAction,
	}
)

func exportCommandAction(c *cli.Context) error {
	// read parms
	parmOpencontrols := c.String("opencontrols")
	parmDestination := c.String("destination")
	parmOutputFormat := c.String("format")
	parmFlatten := c.Bool("flatten")
	parmInferKeys := c.Bool("infer-keys")
	parmDocxtemplater := c.Bool("docxtemplater")
	parmKeySeparator := c.String("key-separator")

	// convert to enum
	outputFormat, err := export.ToOutputFormat(parmOutputFormat)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// --docxtemplater always forces --flatten
	if parmDocxtemplater {
		parmFlatten = true
	}

	// construct args
	config := export.Config{
		Certification:   c.Args().First(),
		OpencontrolDir:  parmOpencontrols,
		DestinationFile: parmDestination,
		OutputFormat:    outputFormat,
		Flatten:         parmFlatten,
		InferKeys:       parmInferKeys,
		Docxtemplater:   parmDocxtemplater,
		KeySeparator:    parmKeySeparator,
	}

	// invoke command
	errs := export.Export(config)
	if errs != nil && len(errs) > 0 {
		err := cli.NewMultiError(errs...)
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
