package errors

import (
	"errors"
	"fmt"
)

// Sentinel errors
var (
	ErrDivisionByZero  = errors.New("division by zero")
	ErrInvalidArgCount = errors.New("invalid argument count")
)

// CalcError wraps a sentinel error with the offending input value.
type CalcError struct {
	Sentinel error
	Input    string
}

// Error returns a formatted message with the sentinel message and input value.
func (e *CalcError) Error() string {
	return fmt.Sprintf("%v: %q", e.Sentinel, e.Input)
}

// Unwrap returns the sentinel error so errors.Is() traverses the chain.
func (e *CalcError) Unwrap() error {
	return e.Sentinel
}
