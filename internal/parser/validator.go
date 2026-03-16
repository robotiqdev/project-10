package parser

import (
	"fmt"
	"strconv"

	errs "github.com/example/calc-app/internal/errors"
)

// ValidateArgCount returns nil if args has exactly 3 elements, otherwise an error.
func ValidateArgCount(args []string) error {
	if len(args) != 3 {
		return &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: fmt.Sprintf("got %d args", len(args))}
	}
	return nil
}

// ValidateOperator returns nil if op is a recognized operator, otherwise an error.
func ValidateOperator(op string) error {
	switch op {
	case "+", "-", "*", "/":
		return nil
	}
	return &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: op}
}

// ValidateNumber returns nil if s can be parsed as a float64, otherwise an error.
func ValidateNumber(s string) error {
	_, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: s}
	}
	return nil
}
