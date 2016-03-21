package main

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry-go/tools/vcs"
	"github.com/opencontrol/compliance-masonry-go/yaml"
	"github.com/opencontrol/compliance-masonry-go/yaml/parser"
	"io/ioutil"
	"os"
)

const (
	DefaultDestination = "opencontrols"
	DefaultConfigYaml  = "opencontrol.yaml"
)

func Get(destination string, config string) {
	err := vcs.Clone("github.com/18F/cg-deck", "atdd", destination)
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
	configSchema, err := yaml.Parse(parser.Parser{}, configBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	configSchema.GetSchemaVersion()
}
