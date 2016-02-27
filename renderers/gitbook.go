package renderers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/geramirez/masonry-go/models"
)

type OpenControlGitBook struct {
	*models.OpenControl
	exportPath string
}

func (openControl *OpenControlGitBook) BuildReadMe() {
	var readme string
	readme += "## Standards  \n"

	var old_family, new_family string
	for standard_key, cert_standard := range openControl.Certification.Standards {

		standard := openControl.Standards[standard_key]
		for control_key, _ := range cert_standard.Controls {
			new_family = standard.Controls[control_key].Family
			// Write Family Summary
			if new_family != old_family {
				readme += fmt.Sprintf(
					"* [%s](%s)  \n",
					standard_key,
					filepath.Join("standards", standard_key+"-"+new_family+".md"),
				)
			}
			readme += fmt.Sprintf(
				"\t* [%s](%s)  \n",
				control_key,
				filepath.Join("standards", standard_key+"-"+control_key+".md"),
			)
			old_family = new_family
		}
	}

	readme += "## Systems  \n"
	for _, system := range openControl.Systems {

		readme += fmt.Sprintf(
			"* [%s](%s)  \n",
			system.Name,
			filepath.Join("standards", system.Key+".md"),
		)

		for _, component := range system.Components {
			readme += fmt.Sprintf(
				"\t* [%s](%s)  \n",
				component.Name,
				filepath.Join("standards", system.Key+"-"+component.Key+".md"),
			)
		}
	}
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "SUMMARY.md"), []byte(readme), 0700)
	ioutil.WriteFile(filepath.Join(openControl.exportPath, "README.md"),  []byte(readme), 0700)

}

func BuildGitbook(opencontrol_dir string, certification_path string, export_path string) {
	openControl := OpenControlGitBook{
		models.LoadData(opencontrol_dir, certification_path),
		export_path,
	}
	if _, err := os.Stat(export_path); os.IsNotExist(err) {
		os.MkdirAll(export_path, 0700)
	}

	openControl.BuildReadMe()
}
