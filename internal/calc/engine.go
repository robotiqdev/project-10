package calc

import (
	"fmt"

	errs "calculator/internal/errors"
)

// Operator represents a supported mathematical operator.
type Operator string

// Operation represents a mathematical operation to perform.
type Operation struct {
	Left     float64
	Right    float64
	Operator Operator
}

// Engine performs mathematical calculations.
type Engine struct{}

// NewEngine creates a new Engine instance.
func NewEngine() *Engine {
	return &Engine{}
}

// Calculate performs the operation and returns a formatted string result.
func (e *Engine) Calculate(op Operation) (string, error) {
	var result float64
	switch op.Operator {
	case "+":
		result = op.Left + op.Right
	case "-":
		result = op.Left - op.Right
	case "*":
		result = op.Left * op.Right
	case "/":
		var err error
		result, err = Divide(op.Left, op.Right)
		if err != nil {
			return "", err
		}
	default:
		return "", &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: string(op.Operator)}
	}
	return fmt.Sprintf("%g", result), nil
}
