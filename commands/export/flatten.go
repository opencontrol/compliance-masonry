package export

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

////////////////////////////////////////////////////////////////////////
// Package functions

// flattenScalar - handle scalar flatten if possible
func flattenScalar(config *Config, value interface{}, key string, flattened *map[string]interface{}) bool {
	// first, check all supported simple types
	result := true
	if _, okStr := value.(string); okStr {
		if config.Debug {
			log.Printf("flatten:Scalar(string): %s=%s\n", key, value.(string))
		}
		(*flattened)[key] = value.(string)
	} else if _, okFloat64 := value.(float64); okFloat64 {
		if config.Debug {
			log.Printf("flatten:Scalar(float64): %s=%f\n", key, value.(float64))
		}
		(*flattened)[key] = value.(float64)
	} else if _, okBool := value.(bool); okBool {
		if config.Debug {
			log.Printf("flatten:Scalar(bool): %s=%t\n", key, value.(bool))
		}
		(*flattened)[key] = value.(bool)
	} else {
		result = false
	}
	if config.Debug {
		debugHook(config, flattened)
	}
	return result
}

// flattenArray - handle embedded arrays
func flattenArray(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// are we an array?
	input, okArray := value.([]interface{})
	if !okArray {
		return false, nil
	}
	if config.Debug {
		log.Printf("flatten:Array:process %s\n", key)
	}

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
		if embedArray && config.Debug {
			log.Printf("flatten:Array:embedArray %s\n", key)
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
			if config.Debug {
				log.Printf("flatten:Array:discoverKey %s=%s\n", key, arrayKeyToUse)
			}
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
		if config.Debug {
			debugHook(config, flattenedToUse)
		}

		// docxtemplater: simple arrays are embedded (not flattened)
		if embedArray {
			// account for single elements with no key; use 'name' as the key to match docxtemplater
			if len(*flattenedToUse) == 1 {
				if val, mapHasEmptyKey := (*flattenedToUse)[""]; mapHasEmptyKey {
					if config.Debug {
						log.Printf("flatten:Array:embedArray:replaceEmptyKey %s\n", key)
					}
					(*flattenedToUse)["name"] = val
					delete((*flattenedToUse), "")
				}
			}
			targetArray = append(targetArray, *flattenedToUse)
			if config.Debug {
				debugHook(config, flattenedToUse)
			}
		}
	}

	// if we are using docxtemplater format, append targetArray as single value for this key
	if config.Docxtemplater && (targetArray != nil) {
		if config.Debug {
			debugHook(config, flattened)
			log.Printf("flatten:Array:useTargetArray %s\n", key)
		}
		(*flattened)[key] = targetArray
		if config.Debug {
			debugHook(config, flattened)
		}
	}

	// all is well
	if config.Debug {
		debugHook(config, flattened)
	}
	return true, nil
}

// flattenMap - handle dictionary
func flattenMap(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// must be a map type
	input, okMapType := value.(map[string]interface{})
	if !okMapType {
		return false, nil
	}
	if config.Debug {
		log.Printf("flatten:Map:process %s\n", key)
	}

	// iterate over key-value pairs
	var newKey string
	for rkey, subValue := range input {
		// first-time logic
		if key != "" {
			newKey = key + config.KeySeparator + rkey
		} else {
			if config.Debug {
				log.Printf("flatten:Map:isFirstTime %s\n", key)
			}
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
	if config.Debug {
		debugHook(config, flattened)
	}
	return true, nil
}

// flattenDriver - handle all known types for flattening
func flattenDriver(config *Config, value interface{}, key string, flattened *map[string]interface{}) (bool, error) {
	// account for unset value - just ignore (?)
	if value == nil {
		if config.Debug {
			log.Printf("flatten: No value for %s\n", key)
		}
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
	if config.Debug {
		debugHook(config, flattened)
	}
	return false, fmt.Errorf("key '%s': unknown value '%v'", key, value)
}

// flattenNormalize - called after everything else, handles control normalization
func flattenNormalize(config *Config, flattened *map[string]interface{}) error {
	// discover all controls
	var allControls []string
	regexControlKeyPattern := "^(?P<prefix_match>data" + config.KeySeparator +
		"components" + config.KeySeparator + "(?P<comp_name>.*?)" + config.KeySeparator +
		"satisfies" + config.KeySeparator + "(?P<control_key>.*?)" +
		config.KeySeparator + ")control_key$"
	regexControlKeyExp, _ := regexp.Compile(regexControlKeyPattern)
	for key, value := range *flattened {
		// must be a string
		valueStr, okStr := value.(string)
		if !okStr {
			continue
		}

		// anything to do?
		regexControlKeyMatch := regexControlKeyExp.FindStringSubmatch(key)
		if len(regexControlKeyMatch) == 0 {
			continue
		}

		// in the list?
		if !stringInSlice(valueStr, allControls) {
			allControls = append(allControls, valueStr)
		}
	}

	// for each control, find the single "winner"
	for i := range allControls {
		control := allControls[i]

		// iterate over the flattened map specifically for this control
		for key, value := range *flattened {
			// anything to do?
			regexControlKeyMatch := regexControlKeyExp.FindStringSubmatch(key)
			if len(regexControlKeyMatch) == 0 {
				continue
			}
			if value.(string) != control {
				continue
			}

			// we simply take the *first* one as the winner. probably stupid.
			normalizedKeyPrefix := fmt.Sprintf("controls%s%s", config.KeySeparator, control)

			// export the actual prefix to steal from the flattened map
			regexControlKeyResult := make(map[string]string)
			for i, controlKeyName := range regexControlKeyExp.SubexpNames() {
				if i != 0 {
					regexControlKeyResult[controlKeyName] = regexControlKeyMatch[i]
				}
			}
			prefixMatch := regexControlKeyResult["prefix_match"]

			// iterate over the flattened map...again
			for key2, value2 := range *flattened {
				// check for and export suffix
				if !strings.HasPrefix(key2, prefixMatch) {
					continue
				}
				suffixMatch := key2[len(prefixMatch):]

				// add normalized entry "as-is"
				newControlKey := fmt.Sprintf("%s%s%s", normalizedKeyPrefix, config.KeySeparator, suffixMatch)
				(*flattened)[newControlKey] = value2
			}
		}
	}

	// we really don't error check here
	return nil
}

// flatten - generic function to flatten JSON or YAML
func flatten(config *Config, input map[string]interface{}, lkey string, flattened *map[string]interface{}) error {
	// start the ball rolling
	processed, err := flattenDriver(config, input, lkey, flattened)
	if err != nil {
		return err
	}
	if !processed {
		return fmt.Errorf("flatten could not process '%v'", input)
	}

	// the final part of flatten is to normalize control output
	return flattenNormalize(config, flattened)
}
