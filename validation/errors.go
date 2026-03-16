package validation

// CalcError is a custom error type that wraps a string message.
// Stub: Error() returns empty string until implementation fills in sentinel values.
type CalcError string

// Error satisfies the error interface.
func (e CalcError) Error() string { return "" }

// Sentinel errors — zero values until implementation provides real messages.
var (
	ErrInvalidOperator CalcError
	ErrInvalidNumber   CalcError
	ErrDivisionByZero  CalcError
	ErrInvalidArgCount CalcError
)
