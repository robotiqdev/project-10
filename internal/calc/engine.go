package calc

import errs "github.com/example/calc-app/internal/errors"

// Engine orchestrates arithmetic operations and result formatting.
type Engine struct{}

// NewEngine constructs and returns a new Engine.
func NewEngine() *Engine {
	return &Engine{}
}

// Calculate performs the operation described by op and returns a formatted
// result string. It returns an error if the operation is invalid (e.g.
// division by zero or unknown operator).
func (e *Engine) Calculate(op Operation) (string, error) {
	var result float64
	switch op.Operator {
	case OpAdd:
		result = Add(op.Left, op.Right)
	case OpSubtract:
		result = Subtract(op.Left, op.Right)
	case OpMultiply:
		result = Multiply(op.Left, op.Right)
	case OpDivide:
		var err error
		result, err = Divide(op.Left, op.Right)
		if err != nil {
			return "", err
		}
	default:
		return "", &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: string(op.Operator)}
	}
	return FormatResult(result), nil
}
