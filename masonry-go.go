package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry-go/gitbook"
)

func main() {
	app := cli.NewApp()
	app.Name = "masonry-go"
	app.Usage = "Open Control CLI Tool"
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
					Value: DefaultDestination,
					Usage: "Location to download the repos.",
				},
				cli.StringFlag{
					Name:  "config",
					Value: DefaultConfigYaml,
					Usage: "Location of system yaml",
				},
				cli.BoolFlag{
					Name:  "verbose, v",
					Usage: "Indicates whether to run the command with verbosity.",
				},
			},
			Action: func(c *cli.Context) {
				Get(c.String("dest"), c.String("config"), c.Bool("verbose"))
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
