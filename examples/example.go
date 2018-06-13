/*
 Copyright (C) 2018 OpenControl Contributors. See LICENSE.md for license.
*/

package main

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/pkg/lib"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"io"
	"os"
)

type plugin struct {
	common.Workspace
}

func simpleDataExtract(p plugin) string {
	selectJustifications := p.GetAllVerificationsWith("NIST-800-53", "CM-2")
	if len(selectJustifications) == 0 {
		return "no data"
	}
	return selectJustifications[0].SatisfiesData.GetImplementationStatus()
}

func run(workspacePath, certPath string, writer io.Writer) {
	workspace, _ := lib.LoadData(workspacePath, certPath)
	sampleData := simpleDataExtract(plugin{workspace})
	fmt.Fprint(writer, sampleData)
}

func main() {
	// in reality you would check the number of args.
	run(os.Args[1], os.Args[2], os.Stdout)
}
