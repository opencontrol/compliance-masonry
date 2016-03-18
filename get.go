package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/opencontrol/compliance-masonry-go/yaml"
	"github.com/opencontrol/compliance-masonry-go/yaml/parser"
	"github.com/opencontrol/compliance-masonry-go/tools/vcs"
)

const (
	DefaultDestination = "opencontrols"
	DefaultConfigYaml  = "opencontrol.yaml"
)

func Get(destination string, config string, verbose bool) {
	log.SetOutput(ioutil.Discard)
	if verbose {
		log.SetOutput(os.Stderr)
	}
	err := vcs.Clone("github.com/18F/cg-deck", "atdd", "deck")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if _, err := os.Stat(config); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", config)
		os.Exit(1)
	}
	configBytes, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	configSchema, err := yaml.Parse(parser.Parser{},configBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	configSchema.GetSchemaVersion()
	//log.Println("in get")
}
