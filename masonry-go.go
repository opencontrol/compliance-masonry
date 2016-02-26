package main

import (
	"github.com/codegangsta/cli"
	"os"

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
						models.LoadData("opencontrols")
						println("New Gitbook Documentation Created")
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
