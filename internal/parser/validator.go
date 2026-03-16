package parser

// ValidateNumber checks whether s can be parsed as a float64.
// Returns nil if valid, or a *errs.CalcError wrapping errs.ErrInvalidNumber if not.
func ValidateNumber(s string) error {
	panic("not implemented")
}

// ParseNumber parses s as a float64.
// Returns the parsed value and nil error on success, or 0 and a *errs.CalcError
// wrapping errs.ErrInvalidNumber on failure.
func ParseNumber(s string) (float64, error) {
	panic("not implemented")
}
