/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package common

//go:generate mockery -name Certification

// Certification is the interface for getting all the attributes for a given certification.
// Schema info: https://github.com/opencontrol/schemas#certifications
//
// GetKey returns the the unique key that represents the name of the certification.
//
// GetSortedStandards returns the list of sorted standard keys.
//
// GetControlKeysFor returns the list of control keys for a given standard key.
type Certification interface {
	GetKey() string
	GetSortedStandards() []string
	GetControlKeysFor(standardKey string) []string
}
