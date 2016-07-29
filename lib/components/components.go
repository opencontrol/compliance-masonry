package components

import (
	"log"
	"sync"

	"github.com/opencontrol/compliance-masonry/lib/components/versions/base"
)

// Components struct is a thread-safe structure mapping for components
type Components struct {
	mapping map[string]base.Component
	sync.RWMutex
}

// NewComponents creates an instance of Components struct
func NewComponents() *Components {
	return &Components{mapping: make(map[string]base.Component)}
}

// Add adds a new component to the component map
func (components *Components) Add(component base.Component) {
	components.Lock()
	components.mapping[component.GetKey()] = component
	components.Unlock()
}

// Get retrieves a new component from the component map
func (components *Components) Get(key string) base.Component {
	components.RLock()
	defer components.RUnlock()
	return components.mapping[key]
}

// CompareAndAdd compares to see if the component exists in the map. If not, it adds the component.
// This function is thread-safe.
func (components *Components) CompareAndAdd(component base.Component) bool {
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

// GetAll retrieves all the components without giving directly to the map.
func (components *Components) GetAll() []base.Component {
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