package gitbook

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry/models"
	"github.com/opencontrol/compliance-masonry/tools/fs"
)

// Config contains data for gitbook export configurations
type Config struct {
	OpencontrolDir string
	Certification  string
	ExportPath     string
	MarkdownPath   string
}

// OpenControlGitBook struct is an extension of models.OpenControl that adds
// an exportPath
type OpenControlGitBook struct {
	*models.OpenControl
	markdownPath string
	exportPath   string
	FSUtil       fs.Util
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
func (config Config) BuildGitbook() {
	openControl := OpenControlGitBook{
		models.LoadData(config.OpencontrolDir, config.Certification),
		config.MarkdownPath,
		config.ExportPath,
		fs.OSUtil{},
	}
	openControl.FSUtil.Mkdirs(config.ExportPath)
	openControl.FSUtil.Mkdirs(filepath.Join(config.ExportPath, "components"))
	openControl.FSUtil.Mkdirs(filepath.Join(config.ExportPath, "standards"))
	openControl.buildSummaries()
	openControl.exportComponents()
	openControl.exportStandards()
}
