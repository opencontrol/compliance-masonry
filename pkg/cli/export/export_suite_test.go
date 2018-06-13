/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package export_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Export Suite")
}
