package lib

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStandard(t *testing.T) {
	// Setup map
	m := newStandards()
	// Get nil component.
	standard := m.Get("test")
	assert.Nil(t, standard)
	// Create mock component
	newStandard := new(mocks.Standard)
	newStandard.On("GetName").Return("test")
	// Test add method
	m.Add(newStandard)
	// Try to retrieve the component again.
	standard = m.Get("test")
	assert.Equal(t, standard.GetName(), "test")
}

func TestBadLoadStandard(t *testing.T) {
	// Setup map
	m := newStandards()
	ws := LocalWorkspace{Standards: m}
	err := ws.LoadStandard("fake.file")
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrStandardSchema, err)
}
