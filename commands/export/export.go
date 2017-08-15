package export

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"

	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/lib/common"
	my_logger "github.com/opencontrol/compliance-masonry/logger"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
)

////////////////////////////////////////////////////////////////////////
// Package functions

// exportJSON - JSON output
func exportJSON(config *Config, workspace common.Workspace, output *exportOutput, writer io.Writer) []error {
	// result
	var errors []error

	// work vars
	var byteSlice []byte
	var err error

	// do the work
	byteSlice, err = json.Marshal(output)
	if err != nil {
		return returnErrors(err)
	}

	// flatten output?
	if config.Flatten {
		my_logger.Debugf("JSON: Flatten")

		// decode json first
		mapped := map[string]interface{}{}
		err = json.Unmarshal(byteSlice, &mapped)
		if err != nil {
			return returnErrors(err)
		}

		// flatten the JSON (recursive)
		var flattened = make(map[string]interface{})
		var lkey string
		err := flatten(config, mapped, lkey, &flattened)
		if err != nil {
			return returnErrors(err)
		}
		var flattenedByteSlice []byte
		flattenedByteSlice, err = json.Marshal(flattened)
		if err != nil {
			return returnErrors(err)
		}
		writer.Write(flattenedByteSlice)
	} else {
		// direct output
		writer.Write(byteSlice)
	}

	return errors
}

// exportYAML - YAML output
func exportYAML(config *Config, workspace common.Workspace, output *exportOutput, writer io.Writer) []error {
	// result
	var dummyErrors []error

	// work vars
	var byteSlice []byte
	var err error

	// do the work
	byteSlice, err = yaml.Marshal(output)
	if err != nil {
		return returnErrors(err)
	}

	// flatten output?
	if config.Flatten {
		// we do not support flatten for YAML - returns 'map[interface {}]interface {}'
		return returnErrors(errors.New("--flatten unsupported for YAML"))
	}
	// direct output
	writer.Write(byteSlice)

	// just so we can return an empty array (not nil)
	return dummyErrors
}

// internal - handle export
func export(config *Config, workspace common.Workspace) []error {
	// sanity
	if len(strings.TrimSpace(config.DestinationFile)) == 0 {
		return returnErrors(errors.New("empty destination files"))
	}

	// create our work object
	var output exportOutput
	output.Config = config
	output.Data.Certification = workspace.GetCertification()
	output.Data.Components = workspace.GetAllComponents()
	output.Data.Standards = workspace.GetAllStandards()

	// handle output destination
	var writer io.Writer
	if config.DestinationFile == "-" {
		// send to stdout
		writer = os.Stdout
	} else {
		// send to file
		file, err := os.Create(config.DestinationFile)
		if err != nil {
			return returnErrors(err)
		}
		writer = file
		defer file.Close()
	}

	// handle the output
	switch config.OutputFormat {
	case FormatJSON:
		return exportJSON(config, workspace, &output, writer)
	case FormatYAML:
		return exportYAML(config, workspace, &output, writer)
	default:
		return returnErrors(fmt.Errorf("unsupported OutputFormat '%s'", config.OutputFormat))
	}
}

// Export loads the inventory and writes output to destinaation
func Export(config Config) []error {
	// resolve the actual certification to use
	certificationPath, errs := certifications.GetCertification(config.OpencontrolDir, config.Certification)
	if errs != nil && len(errs) > 0 {
		return errs
	}

	// load all workspace data
	workspace, errs := lib.LoadData(config.OpencontrolDir, certificationPath)
	if errs != nil && len(errs) > 0 {
		return errs
	}

	// retrieve workspace data and write to output
	return export(&config, workspace)
}
