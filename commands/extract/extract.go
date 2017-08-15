package extract

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/opencontrol/compliance-masonry/lib"
	lib_certifications "github.com/opencontrol/compliance-masonry/lib/certifications"
	"github.com/opencontrol/compliance-masonry/lib/common"
	my_logger "github.com/opencontrol/compliance-masonry/logger"
	"github.com/opencontrol/compliance-masonry/tools/certifications"
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
	Flatten         bool
	InferKeys       bool
	Docxtemplater   bool
	KeySeparator    string
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

// debugHook - adds onto dlv with specific condition
func debugHook(config *Config, flattened *map[string]interface{}) {
	if value, hasKey := (*flattened)["data:components"]; hasKey {
		my_logger.Debugf("Hit debugHook: %v", value)
	}
}

// isScalar - is a given value a supported scalar?
func isScalar(value interface{}) bool {
	// first, check all supported simple types
	result := false
	if _, okStr := value.(string); okStr {
		result = true
	} else if _, okFloat64 := value.(float64); okFloat64 {
		result = true
	} else if _, okBool := value.(bool); okBool {
		result = true
	}
	return result
}

// flattenDiscoverKey - handle what a flattened array key should be
func discoverKey(config *Config, value interface{}, lkey string, index int) string {
	// default value
	defaultKey := lkey + strconv.Itoa(index)

	// only process if we must infer keys
	if !config.InferKeys {
		return defaultKey
	}

	// we can only handle maps
	input, okMapType := value.(map[string]interface{})
	if !okMapType {
		return defaultKey
	}

	// determine weights for the keyname to use
	const invalidKeyName = "invalid"
	const invalidKeyWeight = 99
	keyWeights := make(map[string]int)
	keyWeights["key"] = 0
	keyWeights["control_key"] = 1
	keyWeights["name"] = 2

	// iterate over the map, just looking at the keys
	foundKeyWeight := invalidKeyWeight
	foundKeyName := invalidKeyName
	for rkey, rvalue := range input {
		// must be a string scalar
		_, isStr := rvalue.(string)
		if !isStr {
			continue
		}

		if curKeyWeight, hasKey := keyWeights[rkey]; hasKey {
			if curKeyWeight < foundKeyWeight {
				foundKeyWeight = curKeyWeight
				foundKeyName = rvalue.(string)
			}
		}
	}

	// return the bestest key we can find
	if foundKeyWeight != invalidKeyWeight {
		return lkey + foundKeyName
	}

	// return the default
	return defaultKey
}

// flattenScalar - handle scalar flatten if possible
func flattenScalar(config *Config, value interface{}, key string, flattened *map[string]interface{}) bool {
	// first, check all supported simple types
	result := true
	if _, okStr := value.(string); okStr {
		my_logger.Debugf("flatten:Scalar(string): %s=%s", key, value.(string))
		(*flattened)[key] = value.(string)
	} else if _, okFloat64 := value.(float64); okFloat64 {
		my_logger.Debugf("flatten:Scalar(float64): %s=%f", key, value.(float64))
		(*flattened)[key] = value.(float64)
	} else if _, okBool := value.(bool); okBool {
		my_logger.Debugf("flatten:Scalar(bool): %s=%t", key, value.(bool))
		(*flattened)[key] = value.(bool)
	} else {
		result = false
	}
	debugHook(config, flattened)
	return result
}

// flattenArray - handle embedded arrays
func flattenArray(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// are we an array?
	input, okArray := value.([]interface{})
	if !okArray {
		return false, nil
	}
	my_logger.Debugf("flatten:Array:process %s", key)

	// use a target array as the flattened value for this element
	var theArrayValue interface{}
	var targetArray []interface{}

	// docxtemplater: embed iff all elements are scalar
	embedArray := false
	if config.Docxtemplater {
		embedArray = true
		for i := 0; i < len(input); i++ {
			theArrayValue = input[i]
			if !isScalar(theArrayValue) {
				embedArray = false
				break
			}
		}
		if embedArray {
			my_logger.Debugf("flatten:Array:embedArray %s", key)
		}
	}

	// iterate over the array
	for i := 0; i < len(input); i++ {
		// the value to flatten
		theArrayValue = input[i]

		// what key / map will we use for flattening?
		var arrayKeyToUse string
		var flattenedToUse *map[string]interface{}

		// what should the target map be?
		if embedArray {
			// all scalar values mean we will use a simple map with a well-known data name
			var docxtemplaterArrayMap = make(map[string]interface{})
			arrayKeyToUse = "data"
			flattenedToUse = &docxtemplaterArrayMap
		} else {
			// handle the key name to use
			lkey := key + config.KeySeparator
			arrayKeyToUse = discoverKey(config, theArrayValue, lkey, i)
			my_logger.Debugf("flatten:Array:discoverKey %s=%s", key, arrayKeyToUse)
			flattenedToUse = flattened
		}

		// call the standard flatten function
		processed, err := flattenDriver(config, theArrayValue, arrayKeyToUse, flattenedToUse)
		if err != nil {
			return processed, err
		}
		if !processed {
			return false, fmt.Errorf("key '%s[%d]': flattenDriver returns not processed for '%v'", key, i, theArrayValue)
		}
		debugHook(config, flattenedToUse)

		// docxtemplater: simple arrays are embedded (not flattened)
		if embedArray {
			// account for single elements with no key; use 'name' as the key to match docxtemplater
			if len(*flattenedToUse) == 1 {
				if val, mapHasEmptyKey := (*flattenedToUse)[""]; mapHasEmptyKey {
					my_logger.Debugf("flatten:Array:embedArray:replaceEmptyKey %s", key)
					(*flattenedToUse)["name"] = val
					delete((*flattenedToUse), "")
				}
			}
			targetArray = append(targetArray, *flattenedToUse)
			debugHook(config, flattenedToUse)
		}
	}

	// if we are using docxtemplater format, append targetArray as single value for this key
	if config.Docxtemplater && (targetArray != nil) {
		debugHook(config, flattened)
		my_logger.Debugf("flatten:Array:useTargetArray %s", key)
		(*flattened)[key] = targetArray
		debugHook(config, flattened)
	}

	// all is well
	debugHook(config, flattened)
	return true, nil
}

// flattenMap - handle dictionary
func flattenMap(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// must be a map type
	input, okMapType := value.(map[string]interface{})
	if !okMapType {
		return false, nil
	}
	my_logger.Debugf("flatten:Map:process %s", key)

	// iterate over key-value pairs
	var newKey string
	for rkey, subValue := range input {
		// first-time logic
		if key != "" {
			newKey = key + config.KeySeparator + rkey
		} else {
			my_logger.Debugf("flatten:Map:isFirstTime %s", key)
			newKey = rkey
		}

		// check all of the known types
		processed, err := flattenDriver(config, subValue, newKey, flattened)
		if err != nil {
			return processed, err
		}
		if !processed {
			return false, fmt.Errorf("key '%s': flattenDriver returns not processed for '%v'", newKey, subValue)
		}
	}

	// all is well
	debugHook(config, flattened)
	return true, nil
}

// flattenDriver - handle all known types for flattening
func flattenDriver(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// account for unset value - just ignore (?)
	if value == nil {
		my_logger.Debugf("flatten: No value for %s", key)
		return true, nil
	}

	// some variables
	processed := false
	var err error

	// scalar is simplest - does not invoke anything lower
	processed = flattenScalar(config, value, key, flattened)
	if processed {
		return processed, nil
	}

	// array can recurse; trap error
	processed, err = flattenArray(config, value, key, flattened)
	if err != nil {
		return processed, err
	}
	if processed {
		return processed, nil
	}

	// map can recurse; trap error
	processed, err = flattenMap(config, value, key, flattened)
	if err != nil {
		return processed, err
	}
	if processed {
		return processed, nil
	}

	// we have a truly unknown type
	debugHook(config, flattened)
	return false, fmt.Errorf("key '%s': unknown value '%v'", key, value)
}

// flatten - generic function to flatten JSON or YAML
func flatten(config *Config, input map[string]interface{}, lkey string, flattened *map[string]interface{}) error {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	// start the ball rolling
	processed, err := flattenDriver(config, input, lkey, flattened)
	if err != nil {
		return err
	}
	if !processed {
		return fmt.Errorf("flatten could not process '%v'", input)
	}
	return nil
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

// internal - YAML output
func extractYAML(config *Config, workspace common.Workspace, output *extractOutput, writer io.Writer) []error {
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
