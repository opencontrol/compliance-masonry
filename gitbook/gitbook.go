package gitbook

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry-go/helpers"
	"github.com/opencontrol/compliance-masonry-go/models"
)

// OpenControlGitBook struct is an extension of models.OpenControl that adds
// an exportPath
type OpenControlGitBook struct {
	*models.OpenControl
	markdownPath string
	exportPath   string
}

// ComponentGitbook struct is an extension of models.Component that adds
// an exportPath
type ComponentGitbook struct {
	*models.Component
	exportPath string
}

// ControlGitbook struct is an extension of models.Control that adds
// an exportPath
type ControlGitbook struct {
	*models.Control
	exportPath  string
	standardKey string
	controlKey  string
}

func exportLink(text string, location string) string {
	return fmt.Sprintf("* [%s](%s)  \n", text, location)
}

func replaceParentheses(text string) string {
	return strings.Replace(strings.Replace(text, "(", "", -1), ")", "", -1)
}

// BuildGitbook entry point for creating gitbook
func BuildGitbook(opencontrolDir string, certificationPath string, markdownPath string, exportPath string) {
	openControl := OpenControlGitBook{
		models.LoadData(opencontrolDir, certificationPath),
		markdownPath,
		exportPath,
	}
	helpers.CreateDirectory(exportPath)
	helpers.CreateDirectory(filepath.Join(exportPath, "components"))
	helpers.CreateDirectory(filepath.Join(exportPath, "standards"))
	openControl.buildSummaries()
	openControl.exportComponents()
	openControl.exportStandards()
}
