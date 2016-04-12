// +build !linux,!darwin

package errors

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEOFioEOF(t *testing.T) {
	assert.True(t, EOF(io.EOF))
}

func TestEOFonMatchClosed(t *testing.T) {
	assert.True(t, EOF(New("this is closed")))
}

func TestEOFonMatchResetByPeer(t *testing.T) {
	assert.True(t, EOF(New("connection reset by peer")))
}
