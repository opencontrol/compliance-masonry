package lib

import (
	"log"
	"sync"

	"github.com/opencontrol/compliance-masonry/lib/components/versions/base"
	"github.com/opencontrol/compliance-masonry/tools/fs"
	"path/filepath"
	"github.com/opencontrol/compliance-masonry/tools/constants"
	"github.com/opencontrol/compliance-masonry/lib/components/versions"
	"errors"
)

// componentsMap struct is a thread-safe structure mapping for components
type componentsMap struct {
	mapping map[string]base.Component
	sync.RWMutex
}

// newComponents creates an instance of Components struct
func newComponents() *componentsMap {
	return &componentsMap{mapping: make(map[string]base.Component)}
}

// add adds a new component to the component map
func (components *componentsMap) add(component base.Component) {
	components.Lock()
	components.mapping[component.GetKey()] = component
	components.Unlock()
}

// get retrieves a new component from the component map
func (components *componentsMap) get(key string) base.Component {
	components.RLock()
	defer components.RUnlock()
	return components.mapping[key]
}

// compareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *componentsMap) compareAndAdd(component base.Component) bool {
	components.Lock()
	defer components.Unlock()
	added := false
	if _, exists := components.mapping[component.GetKey()]; !exists {
		components.mapping[component.GetKey()] = component
		added = true
	} else {
		log.Fatalf("Component: %s exisits!\n", component.GetKey())
	}
	return added
}

// getAll retrieves all the components without giving directly to the map.
func (components *componentsMap) getAll() []base.Component {
	components.RLock()
	defer components.RUnlock()
	c := make([]base.Component, len(components.mapping))
	idx := 0
	for _, value := range components.mapping {
		c[idx] = value
		idx++
	}
	return c
}

// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (ws *LocalWorkspace) LoadComponent(componentDir string) error {
	// Get file system assistance.
	fs := fs.OSUtil{}
	// Read the component file.
	fileName := filepath.Join(componentDir, "component.yaml")
	componentData, err := fs.OpenAndReadFile(fileName)
	if err != nil {
		return errors.New(constants.ErrComponentFileDNE)
	}
	// Parse the component.
	var component base.Component
	component, err = versions.ParseComponent(componentData,fileName)
	if err != nil {
		return err
	}
	// Ensure we have a key for the component.
	if component.GetKey() == "" {
		component.SetKey(getKey(componentDir))
	}
	// If the component is new, make sure we load the justifications as well.
	if ws.components.compareAndAdd(component) {
		ws.justifications.LoadMappings(component)
	}
	return nil
}