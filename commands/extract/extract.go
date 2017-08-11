package extract

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/opencontrol/compliance-masonry/lib"
	lib_certifications "github.com/opencontrol/compliance-masonry/lib/certifications"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strings"
)

////////////////////////////////////////////////////////////////////////
// OutputFormat enumeration support

// OutputFormat is the format to use for output file
type OutputFormat int

// local variables to map to / from the enumeration
var outputFormatStrings []string
var outputFormats []OutputFormat

// add string as new mapped enumeration
func ciota(s string) OutputFormat {
	var result = OutputFormat(len(outputFormatStrings) - 1)
	outputFormatStrings = append(outputFormatStrings, s)
	outputFormats = append(outputFormats, result)
	return result
}

// create the enumerations
var (
	FormatUnset = ciota("")
	FormatJSON  = ciota("json")
	FormatYAML  = ciota("yaml")
)

// Convert OutputFormat to string
func (f OutputFormat) String() string {
	return outputFormatStrings[int(f)]
}

// ToOutputFormat converts a string to an OutputFormat enum
func ToOutputFormat(s string) (OutputFormat, error) {
	// sanity
	if len(strings.TrimSpace(s)) == 0 {
		return FormatUnset, errors.New("empty string")
	}

	// scan
	for i, v := range outputFormatStrings {
		if v == s {
			return outputFormats[i], nil
		}
	}

	// not found
	return FormatUnset, errors.New("invalid value")
}

////////////////////////////////////////////////////////////////////////
// Package structures

// Config contains settings for this object
type Config struct {
	Certification   string
	OpencontrolDir  string
	DestinationFile string
	OutputFormat    OutputFormat
}

// internal - structure for JSON / YAML output
type extractData struct {
	Certification common.Certification
	Components    []common.Component
	Standards     []common.Standard
}

// MarshalJSON provides JSON support
func (p *extractData) MarshalJSON() (b []byte, e error) {
	// start the output
	buffer := bytes.NewBufferString("{")

	// certification
	buffer.WriteString("\"certification\":")
	bytesJSON, err := lib_certifications.MarshalJSON(p.Certification)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesJSON))

	// iterate over components
	if len(p.Components) > 0 {
		buffer.WriteString(",\"components\":[")
		for i, v := range p.Components {
			bytesJSON, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(string(bytesJSON))
		}
		buffer.WriteString("]")
	}

	// iterate over standards
	if len(p.Standards) > 0 {
		buffer.WriteString(",\"standards\":[")
		for i, v := range p.Standards {
			bytesJSON, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(string(bytesJSON))
		}
		buffer.WriteString("]")
	}

	// finish json
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

// internal - structure for all extracted data
type extractOutput struct {
	Config *Config
	Data   extractData
}

// MarshalJSON provides JSON support
func (p *extractOutput) MarshalJSON() (b []byte, e error) {
	// start the output
	buffer := bytes.NewBufferString("{")

	// config section
	buffer.WriteString("\"config\":")
	bytesConfig, err := json.Marshal(p.Config)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesConfig))

	// data section
	buffer.WriteString(",\"data\":")
	bytesData, err := json.Marshal(&p.Data)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(string(bytesData))

	// close output
	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

////////////////////////////////////////////////////////////////////////
// Package functions

// utility function to return an error list from single error
func returnErrors(err error) []error {
	var result []error
	result = append(result, err)
	return result
}

// internal - JSON output
func extractJSON(config *Config, workspace common.Workspace, output *extractOutput, writer io.Writer) []error {
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
	writer.Write(byteSlice)

	return errors
}

// internal - YAML output
func extractYAML(config *Config, workspace common.Workspace, output *extractOutput, writer io.Writer) []error {
	// result
	var errors []error

	// work vars
	var byteSlice []byte
	var err error

	// do the work
	byteSlice, err = yaml.Marshal(output)
	if err != nil {
		return returnErrors(err)
	}
	writer.Write(byteSlice)

	return errors
}

// internal - handle extraction
func extract(config *Config, workspace common.Workspace) []error {
	// sanity
	if len(strings.TrimSpace(config.DestinationFile)) == 0 {
		return returnErrors(errors.New("empty destination files"))
	}

	// create our work object
	var output extractOutput
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
		return extractJSON(config, workspace, &output, writer)
	case FormatYAML:
		return extractYAML(config, workspace, &output, writer)
	default:
		return returnErrors(fmt.Errorf("unsupported OutputFormat '%s'", config.OutputFormat))
	}
}

// Extract loads the inventory and writes output to destinaation
func Extract(config Config) []error {
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
	return extract(&config, workspace)
}
