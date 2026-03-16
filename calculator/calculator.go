package calculator

import (
	"fmt"
	"strconv"

	calcerrors "github.com/example/calc-app/internal/errors"
)

func formatNum(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

// Calculate performs the arithmetic operation described by op and returns
// a Result containing the computed value and a human-readable expression string.
func Calculate(op Operation) (Result, error) {
	if op.Operator == "/" && op.Right == 0 {
		return Result{}, &calcerrors.CalcError{Sentinel: calcerrors.ErrDivisionByZero, Input: "0"}
	}

	var val float64
	switch op.Operator {
	case "+":
		val = Add(op.Left, op.Right)
	case "-":
		val = Subtract(op.Left, op.Right)
	case "*":
		val = Multiply(op.Left, op.Right)
	case "/":
		val = Divide(op.Left, op.Right)
	default:
		return Result{}, &calcerrors.CalcError{Sentinel: calcerrors.ErrUnknownOperator, Input: op.Operator}
	}

	expr := fmt.Sprintf("%s %s %s", formatNum(op.Left), op.Operator, formatNum(op.Right))
	return Result{Value: val, Expression: expr}, nil
}
