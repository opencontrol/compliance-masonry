package errors

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorDetails(t *testing.T) {
	dets := map[string]string{
		"error": "this is an error",
	}

	assert.Equal(t, dets, Details(New("this is an error")))
}

func TestErrorDetailsHere(t *testing.T) {
	err := New("this is an error")

	here := Here(err)
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	dets := map[string]string{
		"error":    "this is an error",
		"location": fmt.Sprintf("%s:%d", file, line-1),
	}

	assert.Equal(t, dets, Details(here))
}

func TestErrorDetailsCause(t *testing.T) {
	cas := New("this is a cause")
	err := New("this is an error")

	dets := map[string]string{
		"error": "this is an error",
		"cause": "this is a cause",
	}

	assert.Equal(t, dets, Details(Cause(err, cas)))
}

func TestErrorDetailsMultipleCauses(t *testing.T) {
	top := New("this is the top")
	cas := New("this is a cause")
	err := New("this is an error")

	dets := map[string]string{
		"error":  "this is an error",
		"cause":  "this is a cause",
		"cause2": "this is the top",
	}

	assert.Equal(t, dets, Details(Cause(err, Cause(cas, top))))
}

func TestErrorDetailsTrace(t *testing.T) {
	err := New("this is an error")
	trace := Trace(err).(*TraceError)

	dets := map[string]string{
		"error": "this is an error",
		"trace": trace.Trace(),
	}

	assert.Equal(t, dets, Details(trace))
}

func TestErrorDetailsHereInSubject(t *testing.T) {
	err := New("this is an error")

	here := Here(err)
	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	dets := map[string]string{
		"error":    "this is an error",
		"location": fmt.Sprintf("%s:%d", file, line-1),
		"subject":  "this is a subject",
	}

	sub := Subject(here, "this is a subject")

	assert.Equal(t, dets, Details(sub))
}
