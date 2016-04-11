package common

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Base", func() {

	Describe("Retrieving the schema version", func(){
		table.DescribeTable("GetSchemaVersion", func(b Base, expectedVersion string) {
			assert.Equal(GinkgoT(), b.GetSchemaVersion(), expectedVersion)
		},
			table.Entry("Empty / new base struct", Base{}, ""),
			table.Entry("regular base struct", Base{SchemaVersion: "1.0.0"}, "1.0.0"),
		)
	})
})
