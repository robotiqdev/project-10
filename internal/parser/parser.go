package parser

import "strconv"

// Parse parses the provided args slice into an *Input.
// args must contain exactly three elements: [left, operator, right].
func Parse(args []string) (*Input, error) {
	if err := ValidateArgCount(args); err != nil {
		return nil, err
	}

	op := args[1]
	if err := ValidateOperator(op); err != nil {
		return nil, err
	}

	if err := ValidateNumber(args[0]); err != nil {
		return nil, err
	}
	left, _ := strconv.ParseFloat(args[0], 64)

	if err := ValidateNumber(args[2]); err != nil {
		return nil, err
	}
	right, _ := strconv.ParseFloat(args[2], 64)

	return &Input{Left: left, Right: right, Operator: op}, nil
}
