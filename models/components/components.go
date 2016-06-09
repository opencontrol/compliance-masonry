package components

import (
	"log"
	"sync"

	"github.com/opencontrol/compliance-masonry/models/components/versions/base"
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

// GetAndApply get a component and apply the callback function inside while locking
// components
func (components *Components) GetAndApply(key string, callback func(component base.Component)) {
	components.RLock()
	callback(components.mapping[key])
	components.RUnlock()
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

// GetAll retrieves all the components
func (components *Components) GetAll() map[string]base.Component {
	return components.mapping
}