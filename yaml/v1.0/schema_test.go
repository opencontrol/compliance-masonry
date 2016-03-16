package v1_0

import (
	"testing"
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var parseTests = []struct {
		data []byte
		expectedSchema Schema
		expectedError error
	}{
		{
			data: []byte(`
schema_version: 1.0
name: "test-schema"
maintainers:
  - test@test.com
`),
			expectedSchema: Schema{
				Base: common.Base{SchemaVersion: 1.0},
				Metadata: Metadata{
					Name:"test-schema",
					Maintainers: []string {
						"test@test.com",
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range parseTests {
		s := Schema{}
		err := s.Parse(test.data)
		assert.Equal(t, test.expectedSchema, s)
		assert.Equal(t, test.expectedError, err)
	}
}