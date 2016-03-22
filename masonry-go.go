package main

import (
	"os"
	"path/filepath"

	"fmt"
	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry-go/gitbook"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/config/parser"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"io/ioutil"
	"log"
)

func main() {
	app := cli.NewApp()
	app.Name = "masonry-go"
	app.Usage = "Open Control CLI Tool"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Indicates whether to run the command with verbosity.",
		},
	}
	app.Before = func(c *cli.Context) error {
		// Resets the log to output to nothing
		log.SetOutput(ioutil.Discard)
		if c.Bool("verbose") {
			log.SetOutput(os.Stderr)
			log.Println("Running with verbosity")
		}
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Initialize Open Control documentation repository",
			Action: func(c *cli.Context) {
				println("Documentation Initialized")
			},
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Install compliance dependencies",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "dest",
					Value: constants.DefaultDestination,
					Usage: "Location to download the repos.",
				},
				cli.StringFlag{
					Name:  "config",
					Value: constants.DefaultConfigYaml,
					Usage: "Location of system yaml",
				},
			},
			Action: func(c *cli.Context) {
				config := c.String("config")
				configBytes, err := fs.OpenAndReadFile(config)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				Get(c.String("dest"),
					configBytes,
					&common.ConfigWorker{Downloader: common.NewVCSDownloader(), Parser: parser.Parser{}})
				println("Compliance Dependencies Installed")
			},
		},
		{
			Name:    "docs",
			Aliases: []string{"d"},
			Usage:   "Create Documentation",
			Subcommands: []cli.Command{
				{
					Name:  "gitbook",
					Usage: "Create Gitbook Documentation",
					Action: func(c *cli.Context) {
						opencontrolDir := "opencontrols"
						certification := c.Args().First()
						if certification == "" {
							println("Error: New Missing Certification Argument")
							println("Usage: masonry-go docs gitbook LATO")
						} else {
							certificationPath := filepath.Join(
								opencontrolDir,
								"certifications",
								certification+".yaml",
							)
							if _, err := os.Stat(certificationPath); os.IsNotExist(err) {
								println("Error: %s does not exist", certificationPath)
							} else {
								gitbook.BuildGitbook(
									opencontrolDir,
									certificationPath,
									"markdowns",
									"exports",
								)
								println("New Gitbook Documentation Created")
							}

						}
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
