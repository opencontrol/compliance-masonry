package export

import (
	"strconv"
	"strings"
)

////////////////////////////////////////////////////////////////////////
// Package functions

// escapedPrefixMatch - escape a string in prep for regexp
func escapeStringForRegexp(a string) string {
	specialChars := []string{"(", ")", ".", "?"}
	for i := range specialChars {
		specialChar := specialChars[i]
		a = strings.Replace(a, specialChar, "\\"+specialChar, -1)
	}
	return a
}

// stringInSlice - is a string in a list?
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// utility function to return an error list from single error
func returnErrors(err error) []error {
	var result []error
	result = append(result, err)
	return result
}

// debugHook - adds onto dlv with specific condition
func debugHook(config *Config, _ *map[string]interface{}) {
	// could assume called only in debug mode; check anyway
	if !config.Debug {
		return
	}
	// insert whatever you need for a hook into stopping the program
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
