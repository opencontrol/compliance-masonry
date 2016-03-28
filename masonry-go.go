package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/opencontrol/compliance-masonry-go/config/common"
	"github.com/opencontrol/compliance-masonry-go/config/parser"
	"github.com/opencontrol/compliance-masonry-go/tools/constants"
	"github.com/opencontrol/compliance-masonry-go/tools/fs"
	"github.com/opencontrol/compliance-masonry-go/tools/nested_map"
)

var markdownPath, opencontrolDir, exportPath string

// NewCLIApp creates a new instances of the CLI
func NewCLIApp() *cli.App {
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
				fmt.Println("Documentation Initialized")
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
				wd, err := os.Getwd()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				destination := filepath.Join(wd, c.String("dest"))
				err = Get(destination,
					configBytes,
					&common.ConfigWorker{Downloader: common.NewVCSDownloader(), Parser: parser.Parser{}, ResourceMap: nestedmap.Init()})
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				println("Compliance Dependencies Installed")
			},
		},
		{
			Name:    "docs",
			Aliases: []string{"d"},
			Usage:   "Create Documentation",
			Subcommands: []cli.Command{
				{
					Name:    "gitbook",
					Aliases: []string{"g"},
					Usage:   "Create Gitbook Documentation",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "opencontrols, o",
							Value:       "opencontrols",
							Usage:       "Set opencontrols directory",
							Destination: &opencontrolDir,
						},
						cli.StringFlag{
							Name:        "exports, e",
							Value:       "exports",
							Usage:       "Sets the export directory",
							Destination: &exportPath,
						},
						cli.StringFlag{
							Name:        "markdowns, m",
							Value:       "markdowns",
							Usage:       "Sets the markdowns directory",
							Destination: &markdownPath,
						},
					},
					Action: func(c *cli.Context) {
						config := gitbookConfig{
							certification:  c.Args().First(),
							opencontrolDir: opencontrolDir,
							exportPath:     exportPath,
							markdownPath:   markdownPath,
						}
						messages := config.makeGitbook()
						fmt.Println(strings.Join(messages, "\n"))
					},
				},
			},
		},
	}
	return app
}

func main() {
	app := NewCLIApp()
	app.Run(os.Args)
}
