/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package common

// Verification struct holds data for a specific component and verification
// This is an internal data structure that helps map standards and controls to components
type Verification struct {
	ComponentKey  string
	SatisfiesData Satisfies
}

// Verifications is a slice of type Verifications
type Verifications []Verification

// Len returns the length of the GeneralReferences slice
func (slice Verifications) Len() int {
	return len(slice)
}

// Less returns true if a GeneralReference is less than another reference
func (slice Verifications) Less(i, j int) bool {
	return slice[i].ComponentKey < slice[j].ComponentKey
}

// Swap swaps the two GeneralReferences
func (slice Verifications) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
