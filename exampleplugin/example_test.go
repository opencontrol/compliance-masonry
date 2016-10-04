package main

import (
	"github.com/opencontrol/compliance-masonry/lib/common"
	"github.com/opencontrol/compliance-masonry/lib/common/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleDatNoData(t *testing.T) {
	// Test the case when there is no data

	// create mock workspace
	ws := new(mocks.Workspace)
	ws.On("GetAllVerificationsWith", "standard", "control").Return(common.Verifications{})
	// test function expecting "no data"
	p := plugin{ws}
	data := simpleDataExtract(p)
	assert.Equal(t, data, "no data")
}

func TestSimpleDataWithData(t *testing.T) {
	// Test the case when there is data.

	// create mock workspace
	ws := new(mocks.Workspace)
	satisfies := new(mocks.Satisfies)
	satisfies.On("GetImplementationStatus").Return("IMPLEMENTED")
	ws.On("GetAllVerificationsWith", "standard", "control").Return(common.Verifications{common.Verification{SatisfiesData: satisfies}})
	// test function expecting "IMPLEMENTED"
	p := plugin{ws}
	data := simpleDataExtract(p)
	assert.Equal(t, data, "IMPLEMENTED")
}
