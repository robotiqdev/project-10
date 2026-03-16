package calc

import errs "github.com/example/calc-app/internal/errors"

// Add returns the sum of a and b.
func Add(a, b float64) float64 {
	return 0
}

// Subtract returns the difference of a and b.
func Subtract(a, b float64) float64 {
	return 0
}

// Multiply returns the product of a and b.
func Multiply(a, b float64) float64 {
	return 0
}

// Divide returns a divided by b. Returns an error if b is zero.
func Divide(a, b float64) (float64, error) {
	_ = errs.ErrDivisionByZero // ensure import is used
	return 0, nil
}
