package yaml

import (
	"github.com/opencontrol/compliance-masonry-go/yaml/common/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/vektra/errors"
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

	// Bad base parse
	data = []byte("schema_version: @")
	_, err = Parse(parser, data)
	assert.Contains(t, err.Error(), ErrMalformedBaseYamlPrefix)

	// Malformed yaml - wrong value type
	data = []byte("schema_version: versionone")
	_, err = Parse(parser, data)
	assert.Equal(t, err, ErrCantParseSemver)

	// Non-string version 0.0
	data = []byte("schema_version: 1.0")
	_, err = Parse(parser, data)
	assert.Equal(t, ErrCantParseSemver, err)
}

func TestParseUnknownVersion(t *testing.T) {
	parser := new(mocks.SchemaParser)
	// Unknown version 0.0
	data := []byte(`schema_version: "0.0.0"`)
	_, err := Parse(parser, data)
	assert.Equal(t, ErrUnknownSchemaVersion, err)
}

func TestParseV1_0_0(t *testing.T) {
	// Test that ParseV1_0_0 is not called when only specifying 1.0 as version not the full 1.0.0
	parser := new(mocks.SchemaParser)
	data := []byte(`schema_version: "1.0"`)
	mockSchema := new(mocks.BaseSchema)
	parser.On("ParseV1_0_0", data).Return(mockSchema, nil)
	_, _ = Parse(parser, data)
	parser.AssertNotCalled(t, "ParseV1_0_0", data)

	// Test that ParseV1_0_0 is called
	parser = new(mocks.SchemaParser)
	data = []byte(`schema_version: "1.0.0"`)
	mockSchema = new(mocks.BaseSchema)
	parser.On("ParseV1_0_0", data).Return(mockSchema, nil)
	_, _ = Parse(parser, data)
	parser.AssertCalled(t, "ParseV1_0_0", data)

	// Test that ParseV1_0_0 is called but returns a failure
	expectedError := errors.New("Can't parse yaml")
	parser = new(mocks.SchemaParser)
	data = []byte(`schema_version: "1.0.0"`)
	mockSchema = new(mocks.BaseSchema)
	parser.On("ParseV1_0_0", data).Return(mockSchema, expectedError)
	_, err := Parse(parser, data)
	parser.AssertCalled(t, "ParseV1_0_0", data)
	assert.Equal(t, expectedError, err)
}
