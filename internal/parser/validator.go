package parser

import (
	"strconv"

	errs "calculator/internal/errors"
)

// ValidateNumber checks whether s can be parsed as a float64.
// Returns nil if valid, or a *errs.CalcError wrapping errs.ErrInvalidNumber if not.
func ValidateNumber(s string) error {
	if len(s) > 0 && s[len(s)-1] == '.' {
		return &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: s}
	}
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: s}
	}
	return nil
}

// ParseNumber parses s as a float64.
// Returns the parsed value and nil error on success, or 0 and a *errs.CalcError
// wrapping errs.ErrInvalidNumber on failure.
func ParseNumber(s string) (float64, error) {
	if err := ValidateNumber(s); err != nil {
		return 0, err
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: s}
	}
	return v, nil
}
