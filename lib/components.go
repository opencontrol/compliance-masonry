package lib

import (
	"log"
	"sync"

	"github.com/opencontrol/compliance-masonry/lib/components"
	"github.com/opencontrol/compliance-masonry/lib/common"
)

// componentsMap struct is a thread-safe structure mapping for components
type componentsMap struct {
	mapping map[string]common.Component
	sync.RWMutex
}

// newComponents creates an instance of Components struct
func newComponents() *componentsMap {
	return &componentsMap{mapping: make(map[string]common.Component)}
}

// add adds a new component to the component map
func (components *componentsMap) add(component common.Component) {
	components.Lock()
	components.mapping[component.GetKey()] = component
	components.Unlock()
}

// get retrieves a new component from the component map
func (components *componentsMap) get(key string) common.Component {
	components.RLock()
	defer components.RUnlock()
	return components.mapping[key]
}

// compareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *componentsMap) compareAndAdd(component common.Component) bool {
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
func (components *componentsMap) getAll() []common.Component {
	components.RLock()
	defer components.RUnlock()
	c := make([]common.Component, len(components.mapping))
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
	component, err := components.Load(componentDir)
	if err != nil {
		return err
	}
	// If the component is new, make sure we load the justifications as well.
	if ws.components.compareAndAdd(component) {
		ws.justifications.LoadMappings(component)
	}
	return nil
}