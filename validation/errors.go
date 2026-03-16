package validation

// CalcError is a custom error type that wraps a string message.
// Stub: Error() returns empty string until implementation fills in sentinel values.
type CalcError string

// Error satisfies the error interface.
func (e CalcError) Error() string { return string(e) }

// Sentinel errors.
var (
	ErrInvalidOperator CalcError = "invalid operator"
	ErrInvalidNumber   CalcError = "invalid number"
	ErrDivisionByZero  CalcError = "division by zero"
	ErrInvalidArgCount CalcError = "invalid argument count"
)
