package parser

import (
	"fmt"
	"strconv"

	errs "github.com/example/calc-app/internal/errors"
)

// ValidateOperator checks that op is one of the four supported operators (+, -, *, /).
// Returns a *errs.CalcError wrapping errs.ErrUnknownOperator if invalid.
func ValidateOperator(op string) error {
	switch op {
	case "+", "-", "*", "/":
		return nil
	default:
		return &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: op}
	}
}

// ValidateNumber checks that s can be parsed as a float64.
// Returns a *errs.CalcError wrapping errs.ErrInvalidNumber if invalid.
func ValidateNumber(s string) error {
	if s == "" {
		return &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: s}
	}
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: s}
	}
	return nil
}

// ValidateArgCount checks that exactly 3 args are provided (num1, op, num2).
// Returns a *errs.CalcError wrapping errs.ErrInvalidArgCount if count != 3.
func ValidateArgCount(args []string) error {
	if len(args) != 3 {
		return &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: fmt.Sprintf("%d", len(args))}
	}
	return nil
}
