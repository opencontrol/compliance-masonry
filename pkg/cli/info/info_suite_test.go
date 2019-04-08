/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package info_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestInfo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Info Suite")
}
