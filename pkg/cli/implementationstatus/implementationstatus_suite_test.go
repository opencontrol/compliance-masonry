/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package implementationstatus_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImplementationstatus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Implementation Status Suite")
}
