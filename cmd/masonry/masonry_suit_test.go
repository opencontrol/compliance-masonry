/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMasonryGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ComplianceMasonryGo Suite")
}
