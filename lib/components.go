package lib

import (
	"sync"

	"fmt"
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/components"
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
func (components *componentsMap) get(key string) (component common.Component, found bool) {
	components.RLock()
	defer components.RUnlock()
	component, found = components.mapping[key]
	return
}

// compareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// Returns true if the component was added, returns false if the component was not added.
// This function is thread-safe.
func (components *componentsMap) compareAndAdd(component common.Component) bool {
	components.Lock()
	defer components.Unlock()
	_, exists := components.mapping[component.GetKey()]
	if !exists {
		components.mapping[component.GetKey()] = component
		return true
	}
	return false
}

// getAll retrieves all the components without giving directly to the map.
func (components *componentsMap) getAll() []common.Component {
	components.RLock()
	defer components.RUnlock()
	result := make([]common.Component, len(components.mapping))
	idx := 0
	for _, value := range components.mapping {
		result[idx] = value
		idx++
	}
	return result
}

// LoadComponent imports components into a Component struct and adds it to the
// Components map.
func (ws *localWorkspace) LoadComponent(componentDir string) error {
	component, err := components.Load(componentDir)
	if err != nil {
		return err
	}
	// If the component is new, make sure we load the justifications as well.
	if ws.components.compareAndAdd(component) {
		ws.justifications.LoadMappings(component)
	} else {
		return fmt.Errorf("Component: %s exists!\n", component.GetKey())
	}
	return nil
}

func (ws *localWorkspace) GetAllComponents() []common.Component {
	return ws.components.getAll()
}

func (ws *localWorkspace) GetComponent(component string) (common.Component, bool) {
	return ws.components.get(component)
}
