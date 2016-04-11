package config_test

import (
	. "github.com/opencontrol/compliance-masonry/config"

	. "github.com/onsi/ginkgo"
	"github.com/opencontrol/compliance-masonry/config/common"
	"github.com/opencontrol/compliance-masonry/config/common/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/vektra/errors"
)

var _ = Describe("Parse", func() {
	var (
		parser *mocks.SchemaParser
		err    error
		schema common.BaseSchema
	)

	BeforeEach(func() {
		parser = new(mocks.SchemaParser)
	})

	Describe("bad input scenarios", func() {
		It("should detect there's no data to parse when given nil data", func() {
			schema, err = Parse(parser, nil)
			assert.Equal(GinkgoT(), ErrNoDataToParse, err)
		})
		It("should detect there's no data to parse when given empty data", func() {
			schema, err = Parse(parser, []byte(""))
			assert.Equal(GinkgoT(), ErrNoDataToParse, err)
		})
		It("should detect when it's unable to unmarshal into the base type", func() {
			schema, err = Parse(parser, []byte("schema_version: @"))
			assert.Contains(GinkgoT(), err.Error(), ErrMalformedBaseYamlPrefix)
		})
		It("should detect when it's unable to determine the semver version because it is not in the format", func() {
			schema, err = Parse(parser, []byte("schema_version: versionone"))
			assert.Equal(GinkgoT(), err, ErrCantParseSemver)
		})
		It("should detect when it's unable to determine the semver version because the version is not in string quotes", func() {
			schema, err = Parse(parser, []byte(`schema_version: 1.0`))
			assert.Equal(GinkgoT(), err, ErrCantParseSemver)
		})
		It("should detect when the version is unknown", func() {
			schema, err = Parse(parser, []byte(`schema_version: "0.0.0"`))
			assert.Equal(GinkgoT(), err, ErrUnknownSchemaVersion)
		})
	})
	Describe("ParseV1_0_0 scenarios", func() {
		var (
			data          []byte
			expectedError error
			mockSchema    *mocks.BaseSchema
		)
		BeforeEach(func() {
			expectedError = nil
			mockSchema = new(mocks.BaseSchema)
		})
		JustBeforeEach(func() {
			parser.On("ParseV1_0_0", data).Return(mockSchema, expectedError)
		})
		Context("when the ParseV1_0_0 will not be called", func() {
			It("should not call it when passing in 1.0", func() {
				Parse(parser, []byte(`schema_version: "1.0"`))
				parser.AssertNotCalled(GinkgoT(), "ParseV1_0_0", data)
			})
		})
		Context("when the ParseV1_0_0 will is called", func() {
			BeforeEach(func() {
				data = []byte(`schema_version: "1.0.0"`)
			})
			Context("when ParseV1_0_0 is passed valid data", func() {
				It("should call ParseV1_0_0", func() {
					_, err = Parse(parser, data)
					parser.AssertCalled(GinkgoT(), "ParseV1_0_0", data)
					assert.Equal(GinkgoT(), expectedError, err)
				})
			})
			Context("when ParseV1_0_0 is passed invalid data", func() {
				BeforeEach(func() {
					expectedError = errors.New("can't parse yaml")
				})
				It("should call ParseV1_0_0 but return an error", func() {
					_, err = Parse(parser, data)
					parser.AssertCalled(GinkgoT(), "ParseV1_0_0", data)
					assert.Equal(GinkgoT(), expectedError, err)
				})
			})
		})
	})
})
