// +build linux,darwin

package errors

import (
	"io"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEOFioEOF(t *testing.T) {
	assert.True(t, EOF(io.EOF))
}

func TestEOFonECONNRESET(t *testing.T) {
	assert.True(t, EOF(syscall.ECONNRESET))
}

func TestEOFonENOTCONN(t *testing.T) {
	assert.True(t, EOF(syscall.ENOTCONN))
}

func TestEOFonECONNREFUSED(t *testing.T) {
	assert.True(t, EOF(syscall.ECONNREFUSED))
}

func TestEOFonENETDOWN(t *testing.T) {
	assert.True(t, EOF(syscall.ENETDOWN))
}

func TestEOFonENETUNREACH(t *testing.T) {
	assert.True(t, EOF(syscall.ENETUNREACH))
}

func TestEOFonETIMEDOUT(t *testing.T) {
	assert.True(t, EOF(syscall.ETIMEDOUT))
}
