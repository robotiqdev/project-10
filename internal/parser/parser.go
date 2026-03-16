package parser

import (
	"errors"
	"strconv"

	"github.com/example/calc-app/internal/calc"
)

// Sentinel errors returned by Parse.
var (
	ErrInvalidArgCount = errors.New("invalid argument count")
	ErrInvalidOperand  = errors.New("invalid operand")
	ErrUnknownOperator = errors.New("unknown operator")
)

// Input holds the parsed operands and operator from CLI arguments.
type Input struct {
	Left     float64
	Operator calc.Operator
	Right    float64
}

// Parse converts a slice of CLI arguments into an Input.
// Expected format: [left_operand, operator, right_operand]
// Returns an error if the argument count is wrong, operands cannot be parsed
// as numbers, or the operator is not one of +, -, *, /.
func Parse(args []string) (*Input, error) {
	if err := ValidateArgCount(args); err != nil {
		return nil, err
	}

	left, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, ValidateNumber(args[0])
	}

	if err := ValidateOperator(args[1]); err != nil {
		return nil, err
	}

	right, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, ValidateNumber(args[2])
	}

	return &Input{
		Left:     left,
		Operator: calc.Operator(args[1]),
		Right:    right,
	}, nil
}
