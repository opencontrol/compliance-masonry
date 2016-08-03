package main

import (
	"github.com/opencontrol/compliance-masonry/lib"
	"fmt"
)

type plugin struct {
	lib.Workspace
}

func simpleDataExtractAndFormat(p plugin) string {
	selectJustifications := p.GetJustification("standard", "control")
	if len(selectJustifications) == 0 {
		return "no data"
	}
	return selectJustifications[0].SatisfiesData.GetImplementationStatus()
}

func main() {
	workspace, _ := lib.LoadData("sample opencontrol directory", "cert path")
	fmt.Println(simpleDataExtractAndFormat(plugin{workspace}))
}
