/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package common

//go:generate mockery -name Workspace -testonly

// Workspace represents all the information such as components, standards, and certification as well as
// the result information such as the justifications.
type Workspace interface {
	LoadComponents(string) []error
	LoadStandards(string) []error
	LoadCertification(string) error
	GetCertification() Certification
	GetAllComponents() []Component
	GetComponent(componentKey string) (Component, bool)
	GetAllStandards() []Standard
	GetStandard(standardKey string) (Standard, bool)
	GetAllVerificationsWith(standardKey string, controlKey string) Verifications
}
