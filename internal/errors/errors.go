package errors

import "errors"

// Sentinel errors — placeholder values; implementation will provide real messages.
var (
	ErrUnknownOperator = errors.New("unknown operator")
	ErrInvalidNumber   = errors.New("invalid number")
	ErrDivisionByZero  = errors.New("division by zero")
	ErrInvalidArgCount = errors.New("invalid argument count")
)

// CalcError wraps a sentinel error with the offending input value.
type CalcError struct {
	Sentinel error
	Input    string
}

// Error returns an empty string — not yet implemented.
func (e *CalcError) Error() string { return "" }

// Unwrap returns nil — not yet implemented.
func (e *CalcError) Unwrap() error { return nil }
