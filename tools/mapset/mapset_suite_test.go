/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package mapset_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMapset(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapset Suite")
}
