package parser
import (
	"github.com/opencontrol/compliance-masonry-go/yaml/common"
	"github.com/opencontrol/compliance-masonry-go/yaml/v1.0"
)

type Parser struct {
}

func (p Parser) ParseV1_0(data[] byte) (common.BaseSchema, error) {
	schema := v1_0.Schema{}
	parseError := schema.Parse(data)
	return schema, parseError
}