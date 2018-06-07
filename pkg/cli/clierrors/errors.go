package clierrors

import (
	"context"
	"fmt"
	"os"
	"strings"
)

// Sinlge check failure
func CheckError(err error) {
	if err != nil {
		if err != context.Canceled {
			fmt.Fprintf(os.Stderr, fmt.Sprintf("An error occurred: %v\n", err))
		}
		os.Exit(1)
	}
}

// The following code was taken from https://github.com/urfave/cli/blob/master/errors.go
// MultiError is an error that wraps multiple errors.
type MultiError struct {
	Errors []error
}

// NewMultiError creates a new MultiError. Pass in one or more errors.
func NewMultiError(err ...error) MultiError {
	return MultiError{Errors: err}
}

// Error implements the error interface.
func (m MultiError) Error() string {
	errs := make([]string, len(m.Errors))
	for i, err := range m.Errors {
		errs[i] = err.Error()
	}

	return strings.Join(errs, "\n")
}

// ExitError fulfills both the builtin `error` interface and `ExitCoder`
type ExitError struct {
	exitCode int
	message  interface{}
}

// NewExitError makes a new *ExitError
func NewExitError(message interface{}, exitCode int) *ExitError {
	return &ExitError{
		exitCode: exitCode,
		message:  message,
	}
}

// Error returns the string message, fulfilling the interface required by
// `error`
func (ee *ExitError) Error() string {
	return fmt.Sprintf("%v", ee.message)
}

// ExitCode returns the exit code, fulfilling the interface required by
// `ExitCoder`
func (ee *ExitError) ExitCode() int {
	return ee.exitCode
}

// End code taken from https://github.com/urfave/cli/blob/master/errors.go
