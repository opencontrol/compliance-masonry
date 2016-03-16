package yaml

import (
	"github.com/opencontrol/compliance-masonry-go/yaml/common/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadInputsParse(t *testing.T) {
	parser := new(mocks.SchemaParser)
	// Nil data
	_, err := Parse(parser, nil)
	assert.Equal(t, ErrNoDataToParse, err)

	// Empty Data
	data := []byte("")
	_, err = Parse(parser, data)
	assert.Equal(t, ErrNoDataToParse, err)

	// Malformed yaml - wrong value type
	data = []byte("schema_version: versionone")
	_, err = Parse(parser, data)
	assert.Contains(t, err.Error(), "cannot unmarshal !!str `versionone` into float32")
}

func TestParseUnknownVersion(t *testing.T) {
	parser := new(mocks.SchemaParser)
	// Unknown version 0.0
	data := []byte("schema_version: 0.0")
	_, err := Parse(parser, data)
	assert.Equal(t, ErrUnknownSchemaVersion, err)
}

func TestParseV1_0(t *testing.T) {
	// Test that ParseV1_0 is called
	parser := new(mocks.SchemaParser)
	data := []byte("schema_version: 1.0")
	mockSchema := new(mocks.BaseSchema)
	parser.On("ParseV1_0", data).Return(mockSchema, nil)
	_, _ = Parse(parser, data)
	parser.AssertCalled(t, "ParseV1_0", data)
}
