package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/geramirez/masonry-go/models"
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
			Action: func(c *cli.Context) {
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
						opencontrol_dir := "opencontrols"
						certification := c.Args().First()
						if certification == "" {
							println("Error: New Missing Certification Argument")
							println("Usage: masonry-go docs gitbook LATO")
						} else{
							certification_path := filepath.Join(
								opencontrol_dir,
								"certifications",
								certification + ".yaml",
							)
							if _, err := os.Stat(certification_path); os.IsNotExist(err) {
								println("Error: %s does not exist", certification_path)
							} else {
							models.LoadData(opencontrol_dir, certification_path)
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
