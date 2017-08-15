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
	parmOpencontrols := c.String("opencontrols")
	parmDestination := c.String("destination")
	parmOutputFormat := c.String("format")
	parmFlatten := c.Bool("flatten")
	parmInferKeys := c.Bool("infer-keys")
	parmDocxtemplater := c.Bool("docxtemplater")
	parmKeySeparator := c.String("key-separator")

	// convert to enum
	outputFormat, err := extract.ToOutputFormat(parmOutputFormat)
	if err != nil {
		return cli.NewExitError(err.Error(), 1)
	}

	// --docxtemplater always forces --flatten
	if parmDocxtemplater {
		parmFlatten = true
	}

	// construct args
	config := extract.Config{
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
	errs := extract.Extract(config)
	if errs != nil && len(errs) > 0 {
		err := cli.NewMultiError(errs...)
		return cli.NewExitError(err.Error(), 1)
	}
	return nil
}
