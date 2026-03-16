package calc

import (
	"fmt"

	errs "calculator/internal/errors"
)

// Divide returns a/b, or an error if b == 0.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &errs.CalcError{Sentinel: errs.ErrDivisionByZero, Input: fmt.Sprintf("%g", b)}
	}
	return a / b, nil
}
