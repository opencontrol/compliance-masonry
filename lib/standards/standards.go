package standards

import (
	"io/ioutil"
	"github.com/opencontrol/compliance-masonry/lib/common"
	v1_0_0 "github.com/opencontrol/compliance-masonry/lib/standards/versions/1_0_0"
	"gopkg.in/yaml.v2"
)

type Standard interface {
	GetName() string
	GetControls() map[string]common.Control
	GetControl(string) common.Control
	GetSortedControls() []string
}

func Load(path string) (Standard, error) {
	var standard v1_0_0.Standard
	standardData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, common.ErrReadFile
	}
	err = yaml.Unmarshal(standardData, &standard)
	if err != nil {
		return nil, common.ErrStandardSchema
	}
	return standard, nil
}