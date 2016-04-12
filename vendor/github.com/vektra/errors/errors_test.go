package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorInterface(t *testing.T) {
	_ = error(&ErrorString{})
	_ = error(&HereError{})
	_ = error(&CauseError{})
	_ = error(&TraceError{})
}

func TestErrorsNew(t *testing.T) {
	err := New("this is an error")
	assert.Equal(t, err.Error(), "this is an error")
}

func TestErrorHere(t *testing.T) {
	err := New("this is an error")

	here := Here(err).(*HereError)
	_, _, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	loc := here.FullLocation()

	idx := strings.LastIndex(loc, "/")

	assert.Equal(t, fmt.Sprintf("errors_test.go:%d", line-1), loc[idx+1:])
	assert.Equal(t, fmt.Sprintf("errors/errors_test.go:%d", line-1), here.Location())
}

func TestErrorHereString(t *testing.T) {
	err := New("this is an error")

	here := Here(err)
	_, _, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	assert.Equal(t, fmt.Sprintf("errors/errors_test.go:%d: this is an error", line-1), here.Error())
}

func TestErrorHereChecksForHere(t *testing.T) {
	err := New("this is an error")

	here := Here(err)

	if here != Here(here) {
		t.Fatalf("Here re-wrapped another Here")
	}
}

func TestErrorCause(t *testing.T) {
	orig := New("this is the cause")
	err := New("this is an error")

	cause := Cause(err, orig).(*CauseError)

	assert.Equal(t, cause.Cause(), orig)
}

func TestErrorTrace(t *testing.T) {
	err := New("this is an error")

	here := Trace(err).(*TraceError)
	_, _, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	trace := here.Trace()

	lines := strings.Split(trace, "\n")

	loc := lines[4]
	sidx := strings.LastIndex(loc, "/")
	widx := strings.LastIndex(loc, " ")

	assert.Equal(t, fmt.Sprintf("errors_test.go:%d", line-1), loc[sidx+1:widx])
}

func TestErrorContext(t *testing.T) {
	err := New("this is an error")

	ctx := Context(err, "while testing").(*ContextError)

	assert.Equal(t, "while testing", ctx.Context())
	assert.Equal(t, "while testing: this is an error", ctx.Error())
}

func TestErrorSubject(t *testing.T) {
	err := New("this is an error")

	sub := Subject(err, "this is a subject").(*SubjectError)

	assert.Equal(t, "this is a subject", sub.Subject().(string))
	assert.Equal(t, "this is an error: this is a subject", sub.Error())
}

func TestErrorFormat(t *testing.T) {
	fmt := "error: %s"

	err := Format(fmt, "too many errors")

	assert.Equal(t, "error: too many errors", err.Error())
}

func TestErrorPrintGeneric(t *testing.T) {
	err := New("this is an error")

	var buf bytes.Buffer

	Print(err, &buf)

	assert.Equal(t, "error: this is an error\n", buf.String())
}

func TestErrorPrintHere(t *testing.T) {
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Here(err), &buf)

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	loc := fmt.Sprintf("%s:%d", file, line-2)

	assert.Equal(t, " from: "+loc+"\nerror: "+err.Error()+"\n", buf.String())
}

func TestErrorPrintHereAroundCause(t *testing.T) {
	cas := New("this is a cause")
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Here(Cause(err, cas)), &buf)

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	loc := fmt.Sprintf("%s:%d", file, line-2)

	assert.Equal(t, " from: "+loc+"\nerror: "+err.Error()+"\ncause: "+cas.Error()+"\n", buf.String())
}

func TestErrorPrintCauseAroundHere(t *testing.T) {
	cas := New("this is a cause")
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Cause(Here(err), cas), &buf)

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	loc := fmt.Sprintf("%s:%d", file, line-2)

	assert.Equal(t, " from: "+loc+"\nerror: "+err.Error()+"\ncause: "+cas.Error()+"\n", buf.String())
}

func TestErrorPrintCause(t *testing.T) {
	cas := New("this is the cause")
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Cause(err, cas), &buf)

	assert.Equal(t, "error: this is an error\ncause: this is the cause\n", buf.String())
}

func TestErrorPrintMultipleCauses(t *testing.T) {
	top := New("this is the real cause")
	cas := New("this is the cause")
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Cause(err, Cause(cas, top)), &buf)

	assert.Equal(t, "error: this is an error\ncause: this is the cause\ncause: this is the real cause\n", buf.String())
}

func TestErrorPrintMultipleCauses3(t *testing.T) {
	top := New("this is the real cause")
	hap := New("this just happened")
	cas := New("this is the cause")
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Cause(err, Cause(cas, Cause(hap, top))), &buf)

	assert.Equal(t, "error: this is an error\ncause: this is the cause\ncause: this just happened\ncause: this is the real cause\n", buf.String())
}

func TestErrorPrintTrace(t *testing.T) {
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Trace(err), &buf)

	lines := strings.Split(buf.String(), "\n")

	assert.Equal(t, "error: this is an error", lines[0])
	assert.Equal(t, "trace:", lines[1])
	assert.NotEqual(t, "", lines[2])
}

func TestErrorPrintTraceAroundHere(t *testing.T) {
	err := New("this is an error")

	var buf bytes.Buffer

	Print(Trace(Here(err)), &buf)

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		panic("caller is busted")
	}

	loc := fmt.Sprintf("%s:%d", file, line-2)

	lines := strings.Split(buf.String(), "\n")

	assert.Equal(t, " from: "+loc, lines[0])
	assert.Equal(t, "error: this is an error", lines[1])
	assert.Equal(t, "trace:", lines[2])
	assert.NotEqual(t, "", lines[3])
}

func TestErrorUnwrapHere(t *testing.T) {
	err := New("this is an error")

	assert.Equal(t, Unwrap(Here(err)), err)
}

func TestErrorUnwrapCause(t *testing.T) {
	cas := New("this is a cause error")
	err := New("this is an error")

	assert.Equal(t, Unwrap(Cause(err, cas)), err)
}

func TestErrorUnwrapTrace(t *testing.T) {
	err := New("this is an error")

	assert.Equal(t, Unwrap(Trace(err)), err)
}

func TestErrorUnwrapContext(t *testing.T) {
	err := New("this is an error")

	assert.Equal(t, Unwrap(Context(err, "blah")), err)
}

func TestErrorUnwrapSubject(t *testing.T) {
	err := New("this is an error")

	assert.Equal(t, Unwrap(Subject(err, "blah")), err)
}

func TestErrorUnwrapMultipleLevels(t *testing.T) {
	err := New("this is an error")

	assert.Equal(t, Unwrap(Subject(Here(err), "blah")), err)
}

func TestErrorEqualHere(t *testing.T) {
	err := New("this is an error")

	assert.True(t, Equal(Here(err), err))
}

func TestErrorEqualCause(t *testing.T) {
	cas := New("this is a cause error")
	err := New("this is an error")

	assert.True(t, Equal(Cause(err, cas), err))
}

func TestErrorEqualTrace(t *testing.T) {
	err := New("this is an error")

	assert.True(t, Equal(Trace(err), err))
}

func TestErrorEqualContext(t *testing.T) {
	err := New("this is an error")

	assert.True(t, Equal(Context(err, "blah"), err))
}

func TestErrorEqualSubject(t *testing.T) {
	err := New("this is an error")

	assert.True(t, Equal(Subject(err, "blah"), err))
}

func TestErrorEqualGeneric(t *testing.T) {
	err := New("this is an error")

	assert.True(t, Equal(err, err))
}
