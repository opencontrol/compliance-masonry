package lib

import (
	"github.com/opencontrol/compliance-masonry/pkg/lib/common"
	"github.com/opencontrol/compliance-masonry/pkg/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStandard(t *testing.T) {
	// Setup map
	m := newStandards()
	// Get nil component.
	standard, found := m.get("test")
	assert.False(t, found)
	assert.Nil(t, standard)
	// Create mock component
	newStandard := new(mocks.Standard)
	newStandard.On("GetName").Return("test")
	// Test add method
	m.add(newStandard)
	// Try to retrieve the component again.
	standard, found = m.get("test")
	assert.True(t, found)
	assert.Equal(t, standard.GetName(), "test")
}

func TestBadLoadStandard(t *testing.T) {
	// Setup map
	m := newStandards()
	ws := localWorkspace{standards: m}
	err := ws.LoadStandard("fake.file")
	assert.NotNil(t, err)
	assert.Equal(t, common.ErrStandardSchema, err)
}
