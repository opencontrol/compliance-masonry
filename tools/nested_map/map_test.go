package nestedmap

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReserve(t *testing.T) {
	m := Init()

	// Regular reservation
	res := m.Reserve("key1", "key2", "value")
	assert.Equal(t, res.Error, nil)
	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.Value, "value")

	// Repeat
	res = m.Reserve("key1", "key2", "value")
	assert.Equal(t, res.Error, ErrAlreadyExists)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")

	// bad inputs: no OuterKey
	res = m.Reserve("", "key2", "value")
	assert.Equal(t, res.Error, ErrEmptyInput)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")

	// bad inputs: no InnerKey
	res = m.Reserve("key1", "", "value")
	assert.Equal(t, res.Error, ErrEmptyInput)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")

	// bad inputs: no Value
	res = m.Reserve("key1", "key2", "")
	assert.Equal(t, res.Error, ErrEmptyInput)
	assert.Equal(t, res.Success, false)
	assert.Equal(t, res.Value, "")
}