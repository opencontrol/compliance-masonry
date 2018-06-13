/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package opencontrol_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCommon(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opencontrol Versions Suite")
}
