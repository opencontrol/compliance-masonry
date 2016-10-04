package main

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

type plugin struct {
	common.Workspace
}

func simpleDataExtract(p plugin) string {
	selectJustifications := p.GetAllVerificationsWith("standard", "control")
	if len(selectJustifications) == 0 {
		return "no data"
	}
	return selectJustifications[0].SatisfiesData.GetImplementationStatus()
}

func main() {
	workspace, _ := lib.LoadData("sample opencontrol directory", "cert path")
	sampleData := simpleDataExtract(plugin{workspace})
	fmt.Println(sampleData)
}
