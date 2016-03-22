package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSchemaVersion(t *testing.T) {
	var SchemaVersionTests = []struct {
		b               Base
		expectedVersion string
	}{
		{
			b:               Base{},
			expectedVersion: "",
		},
		{
			b:               Base{SchemaVersion: "1.0.0"},
			expectedVersion: "1.0.0",
		},
	}
	for _, test := range SchemaVersionTests {
		assert.Equal(t, test.expectedVersion, test.b.GetSchemaVersion())
	}
}
