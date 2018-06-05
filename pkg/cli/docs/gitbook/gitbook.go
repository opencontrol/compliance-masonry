package gitbook

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/opencontrol/compliance-masonry/lib"
	"github.com/opencontrol/compliance-masonry/lib/common"
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
	common.Workspace
	markdownPath string
	exportPath   string
	FSUtil       fs.Util
}

// ComponentGitbook struct is an extension of models.Component that adds
// an exportPath
type ComponentGitbook struct {
	common.Component
	exportPath string
}

// ControlGitbook struct is an extension of models.Control that adds
// an exportPath
type ControlGitbook struct {
	common.Control
	exportPath  string
	standardKey string
	controlKey  string
}

func exportLink(text string, location string) string {
	return fmt.Sprintf("* [%s](%s)\n", text, location)
}

func replaceParentheses(text string) string {
	return strings.Replace(strings.Replace(text, "(", "", -1), ")", "", -1)
}

// BuildGitbook entry point for creating gitbook
func (config Config) BuildGitbook() []error {
	var errs []error
	openControlData, err := lib.LoadData(config.OpencontrolDir, config.Certification)
	if err != nil && len(err) > 0 {
		return append(errs, err...)
	}
	openControl := OpenControlGitBook{
		openControlData,
		config.MarkdownPath,
		config.ExportPath,
		fs.OSUtil{},
	}
	openControl.FSUtil.Mkdirs(config.ExportPath)
	openControl.FSUtil.Mkdirs(filepath.Join(config.ExportPath, "components"))
	openControl.FSUtil.Mkdirs(filepath.Join(config.ExportPath, "standards"))
	if err := openControl.buildSummaries(); err != nil {
		return append(errs, err)
	}
	openControl.exportComponents()
	openControl.exportStandards()
	return nil
}
