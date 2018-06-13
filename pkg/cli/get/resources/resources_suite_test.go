/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package resources

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestResources(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Resources Suite")
}
