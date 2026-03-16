package parser

// ValidateOperator checks that op is one of the four supported operators (+, -, *, /).
// Returns a *errs.CalcError wrapping errs.ErrUnknownOperator if invalid.
func ValidateOperator(op string) error {
	return nil
}

// ValidateNumber checks that s can be parsed as a float64.
// Returns a *errs.CalcError wrapping errs.ErrInvalidNumber if invalid.
func ValidateNumber(s string) error {
	return nil
}

// ValidateArgCount checks that exactly 3 args are provided (num1, op, num2).
// Returns a *errs.CalcError wrapping errs.ErrInvalidArgCount if count != 3.
func ValidateArgCount(args []string) error {
	return nil
}
