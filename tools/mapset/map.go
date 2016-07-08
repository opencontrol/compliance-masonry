package mapset

import (
	"github.com/vektra/errors"
	"gopkg.in/fatih/set.v0"
)

// MapSet is the map with each value being a set.
type MapSet struct {
	mapOfSet map[string]*set.Set
}

// Init returns an initialized NestedMap
func Init() MapSet {
	return MapSet{mapOfSet: make(map[string]*set.Set)}
}

// Result is the result from any NestedMap operations
type Result struct {
	Value   string
	Error   error
	Success bool
}

var (
	// ErrEmptyInput represents that the input into the operation is not complete
	ErrEmptyInput = errors.New("One or more inputs are empty")
)

// Reserve will put a space into the map for the value given that key. Will return false if there is already an entry.
func (m *MapSet) Reserve(key string, value string) (result Result) {
	if key == "" || value == "" {
		result.Error = ErrEmptyInput
		return
	}
	var innerSet *set.Set
	if _, ok := m.mapOfSet[key]; !ok {
		innerSet = set.New()
		m.mapOfSet[key] = innerSet
	}
	if m.mapOfSet[key].Has(value) {
		result.Success = false
		result.Value = value
		return
	}
	m.mapOfSet[key].Add(value)
	result.Success = true
	result.Value = value
	return
}
