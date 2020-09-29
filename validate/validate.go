package validate

import (
	"fmt"
	"github.com/opencontrol/compliance-masonry/pkg/lib"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"os"
)

// Validate validates opencontrol masonry repository that has been previously obtained by masonry get
func Validate() {
	problems := make([]string, 0)
	workspace, errors := lib.LoadData("opencontrols/", "opencontrols/certifications/fedramp-high.yaml")
	if errors != nil {
		fmt.Println(errors)
		os.Exit(1)
	}
	for _, component := range workspace.GetAllComponents() {
		problems = append(problems, validateComponent(workspace, component)...)
	}
	for _, problem := range problems {
		fmt.Println(problem)
	}
	os.Exit(len(problems))
}

func validateComponent(workspace common.Workspace, component common.Component) []string {
	problems := make([]string, 0)
	uniq := make(map[string]map[string]common.Satisfies)

	for _, satisfy := range component.GetAllSatisfies() {
		standardKey := satisfy.GetStandardKey()
		_, ok := uniq[standardKey]
		if !ok {
			_, found := workspace.GetStandard(standardKey)
			if !found {
				problems = append(problems, fmt.Sprintf("Component %s references standard %s, however that cannot be found in the workspace.", component.GetName(), standardKey))
				continue
			}
			uniq[standardKey] = make(map[string]common.Satisfies)
		}
		standard, _ := workspace.GetStandard(standardKey)

		standardControl := standard.GetControl(satisfy.GetControlKey())
		if standardControl == nil {
			problems = append(problems, fmt.Sprintf("Could not find reference %s in the standard %s", satisfy.GetControlKey(), standardKey))
		}

		_, found := uniq[standardKey][satisfy.GetControlKey()]
		if found {
			problems = append(problems, fmt.Sprintf("Component %s: Duplicate items found: %s", component.GetKey(), satisfy.GetControlKey()))
		}
		uniq[standardKey][satisfy.GetControlKey()] = satisfy

		switch satisfy.GetImplementationStatus() {
		case "complete", "partial", "not applicable", "planned", "unsatisfied", "unknown", "none":
			break
		default:
			problems = append(problems, fmt.Sprintf("Found non-standard implementation_status: %s.", satisfy.GetImplementationStatus()))
			break
		}
		problems = append(problems, validateNarratives(component, satisfy)...)

	}
	return problems
}

func validateNarratives(component common.Component, satisfy common.Satisfies) []string {
	problems := make([]string, 0)

	requireKey := len(satisfy.GetNarratives()) > 1
	uniqNarratives := make(map[string]bool)
	for _, narrative := range satisfy.GetNarratives() {
		key := narrative.GetKey()
		if requireKey && key == "" {
			problems = append(problems, fmt.Sprintf("Component %s: Satisfy '%s': Narrative key is required when multiple narratives are present.", component.GetKey(), satisfy.GetControlKey()))
		}

		if len(key) > 6 {
			problems = append(problems, fmt.Sprintf("Component %s: Satisfy '%s': Long narrative key probably malformed: '%s'", component.GetKey(), satisfy.GetControlKey(), key))

		}

		if key != "" {
			_, found := uniqNarratives[key]
			if found {
				problems = append(problems, fmt.Sprintf("Component %s:, Satisfy '%s': Duplicate narratives sequence found: %s", component.GetKey(), satisfy.GetControlKey(), key))

			}
		}
		uniqNarratives[key] = true
	}
	return problems
}
