package main

import (
	"testing"
	"github.com/opencontrol/compliance-masonry/lib/mocks"
	commonMocks"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/opencontrol/compliance-masonry/lib"
)

func TestSimpleDataExtractAndFormat(t *testing.T) {
	// Test the case when there is no data

	// create mock workspace
	ws := new(mocks.Workspace)
	ws.On("GetJustification", "standard", "control").Return(lib.Verifications{})
	// test function expecting "no data"
	p := plugin{ws}
	data := simpleDataExtractAndFormat(p)
	assert.Equal(t, data, "no data")

	// Test the case when there is data.

	// create mock workspace
	ws = new(mocks.Workspace)
	satisfies := new(commonMocks.Satisfies)
	satisfies.On("GetImplementationStatus").Return("IMPLEMENTED")
	ws.On("GetJustification", "standard", "control").Return(lib.Verifications{lib.Verification{SatisfiesData: satisfies}})
	// test function expecting "no data"
	p = plugin{ws}
	data = simpleDataExtractAndFormat(p)
	assert.Equal(t, data, "IMPLEMENTED")
}