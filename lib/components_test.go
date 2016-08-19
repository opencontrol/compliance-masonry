package lib

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAddComponent(t *testing.T) {
	// Setup map
	m := newComponents()
	// Get nil component.
	component := m.Get("test")
	assert.Nil(t, component)
	// Create mock component
	newComponent := new(mocks.Component)
	newComponent.On("GetKey").Return("test")
	// Test add method
	m.add(newComponent)
	// Try to retrieve the component again.
	component = m.Get("test")
	assert.Equal(t, component.GetKey(), "test")
}

func TestCompareAndAddComponent(t *testing.T) {
	m := newComponents()
	// Get nil component.
	component := m.Get("test")
	assert.Nil(t, component)
	// Create mock component
	newComponent := new(mocks.Component)
	newComponent.On("GetKey").Return("test")
	// Use compare and add initially.
	added := m.CompareAndAdd(newComponent)
	assert.True(t, added)
	// Use compare and add again to show failure.
	added = m.CompareAndAdd(newComponent)
	assert.False(t, added)
}