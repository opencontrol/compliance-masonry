package parser

import (
	"github.com/opencontrol/compliance-masonry/config/common"
	v1_0_0 "github.com/opencontrol/compliance-masonry/config/versions/1.0.0"
)

// Parser is the concrete implementation of parsing different schema versions.
type Parser struct{}

// ParseV1_0_0 parses data using the v1.0.0 schema.
func (p Parser) ParseV1_0_0(data []byte) (common.BaseSchema, error) {
	config := new(v1_0_0.Schema)
	parseError := config.Parse(data)
	return config, parseError
}
