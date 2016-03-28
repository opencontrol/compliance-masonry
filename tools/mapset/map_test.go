package mapset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReserve(t *testing.T) {
	m := Init()

	// Regular reservation
	res := m.Reserve("key1", "value")
	assert.Equal(t, res.Error, nil)
	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.Value, "value")

	// Repeat
	res = m.Reserve("key1", "value")
	assert.Equal(t, res.Error, ErrAlreadyExists)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")

	// bad inputs: no key
	res = m.Reserve("", "value")
	assert.Equal(t, res.Error, ErrEmptyInput)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")

	// bad inputs: no Value
	res = m.Reserve("key1", "")
	assert.Equal(t, res.Error, ErrEmptyInput)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")
}
