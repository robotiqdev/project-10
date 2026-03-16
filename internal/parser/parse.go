package parser

import (
	errs "calculator/internal/errors"
)

// Input holds the parsed values from command-line arguments.
type Input struct {
	Left     float64
	Right    float64
	Operator string
}

// Parse parses command-line arguments into an Input.
// Expected format: [leftNumber, operator, rightNumber].
func Parse(args []string) (Input, error) {
	if len(args) != 3 {
		return Input{}, &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: ""}
	}

	left, err := ParseNumber(args[0])
	if err != nil {
		return Input{}, err
	}

	operator := args[1]
	switch operator {
	case "+", "-", "*", "/":
		// valid
	default:
		return Input{}, &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: operator}
	}

	right, err := ParseNumber(args[2])
	if err != nil {
		return Input{}, err
	}

	return Input{Left: left, Right: right, Operator: operator}, nil
}
