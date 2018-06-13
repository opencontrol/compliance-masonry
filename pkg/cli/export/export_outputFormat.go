/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package export

import (
	"errors"
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
