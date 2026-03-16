package validation

// ValidateOperator checks whether op is one of the four supported operators: +, -, *, /.
// Returns nil if valid, or ErrInvalidOperator otherwise.
func ValidateOperator(op string) error {
	switch op {
	case "+", "-", "*", "/":
		return nil
	default:
		return ErrInvalidOperator
	}
}

// ExtractOperator validates op and returns it on success.
// Returns ("", ErrInvalidOperator) for unsupported operators.
func ExtractOperator(op string) (string, error) {
	if err := ValidateOperator(op); err != nil {
		return "", err
	}
	return op, nil
}
