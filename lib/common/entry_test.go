package common

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/tools/constants"
)

var _ = Describe("Entry", func() {
	Describe("Retrieving the config file", func() {
		table.DescribeTable("GetConfigFile", func(e Entry, expectedPath string) {
			assert.Equal(GinkgoT(), e.GetConfigFile(), expectedPath)
		},
			table.Entry("Empty / new base struct to return default", Entry{}, constants.DefaultConfigYaml),
			table.Entry("overriden config file path", Entry{Path: "samplepath"}, "samplepath"),
		)
	})
})
