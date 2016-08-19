package lib

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"github.com/opencontrol/compliance-masonry/lib/result"
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

func TestLoadSameComponentTwice(t *testing.T) {
	ws := LocalWorkspace{Components: newComponents(), Justifications: result.NewJustifications()}
	componentPath := filepath.Join("..", "fixtures", "component_fixtures", "v3_1_0", "EC2")
	err := ws.LoadComponent(componentPath)
	// Should load the component without a problem.
	assert.Nil(t, err)
	actualComponent := ws.Components.Get("EC2")
	assert.NotNil(t, actualComponent)
	// Try to load component again.
	err = ws.LoadComponent(componentPath)
	// Should return an error that this component was already loaded.
	assert.NotNil(t, err)
	assert.Equal(t, "Component: EC2 exists!\n", err.Error())
}

func TestBadLoadComponent(t *testing.T) {
	ws := LocalWorkspace{}
	err := ws.LoadComponent("fake.file")
	// Should return an error because it can't load the file.
	assert.Equal(t, "Component files does not exist", err.Error())
}