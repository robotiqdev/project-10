package calculator_test

import (
	"testing"

	"github.com/example/calc-app/calculator"
)

// TestCalculateAddition verifies that Calculate correctly adds two operands,
// returning the right Value and Expression in the Result struct.
func TestCalculateAddition(t *testing.T) {
	tests := []struct {
		name           string
		op             calculator.Operation
		expectedValue  float64
		expectedExpr   string
	}{
		{
			name:          "integer operands",
			op:            calculator.Operation{Left: 3, Operator: "+", Right: 4},
			expectedValue: 7,
			expectedExpr:  "3 + 4",
		},
		{
			name:          "decimal operands",
			op:            calculator.Operation{Left: 1.5, Operator: "+", Right: 2.5},
			expectedValue: 4,
			expectedExpr:  "1.5 + 2.5",
		},
		{
			name:          "negative left operand",
			op:            calculator.Operation{Left: -3, Operator: "+", Right: 4},
			expectedValue: 1,
			expectedExpr:  "-3 + 4",
		},
		{
			name:          "negative right operand",
			op:            calculator.Operation{Left: 10, Operator: "+", Right: -4},
			expectedValue: 6,
			expectedExpr:  "10 + -4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Value != tt.expectedValue {
				t.Errorf("Calculate(%v).Value = %v, want %v", tt.op, got.Value, tt.expectedValue)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}

// TestCalculateSubtraction verifies that Calculate correctly subtracts operands.
func TestCalculateSubtraction(t *testing.T) {
	tests := []struct {
		name          string
		op            calculator.Operation
		expectedValue float64
		expectedExpr  string
	}{
		{
			name:          "integer operands",
			op:            calculator.Operation{Left: 10, Operator: "-", Right: 3},
			expectedValue: 7,
			expectedExpr:  "10 - 3",
		},
		{
			name:          "decimal operands",
			op:            calculator.Operation{Left: 5.5, Operator: "-", Right: 2.5},
			expectedValue: 3,
			expectedExpr:  "5.5 - 2.5",
		},
		{
			name:          "negative left operand",
			op:            calculator.Operation{Left: -5, Operator: "-", Right: 3},
			expectedValue: -8,
			expectedExpr:  "-5 - 3",
		},
		{
			name:          "negative right operand",
			op:            calculator.Operation{Left: 5, Operator: "-", Right: -3},
			expectedValue: 8,
			expectedExpr:  "5 - -3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Value != tt.expectedValue {
				t.Errorf("Calculate(%v).Value = %v, want %v", tt.op, got.Value, tt.expectedValue)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}

// TestCalculateMultiplication verifies that Calculate correctly multiplies operands.
func TestCalculateMultiplication(t *testing.T) {
	tests := []struct {
		name          string
		op            calculator.Operation
		expectedValue float64
		expectedExpr  string
	}{
		{
			name:          "integer operands",
			op:            calculator.Operation{Left: 3, Operator: "*", Right: 4},
			expectedValue: 12,
			expectedExpr:  "3 * 4",
		},
		{
			name:          "decimal operands",
			op:            calculator.Operation{Left: 1.5, Operator: "*", Right: 2},
			expectedValue: 3,
			expectedExpr:  "1.5 * 2",
		},
		{
			name:          "negative left operand",
			op:            calculator.Operation{Left: -3, Operator: "*", Right: 4},
			expectedValue: -12,
			expectedExpr:  "-3 * 4",
		},
		{
			name:          "both negative operands",
			op:            calculator.Operation{Left: -3, Operator: "*", Right: -4},
			expectedValue: 12,
			expectedExpr:  "-3 * -4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Value != tt.expectedValue {
				t.Errorf("Calculate(%v).Value = %v, want %v", tt.op, got.Value, tt.expectedValue)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}

// TestCalculateDivision verifies that Calculate correctly divides operands,
// including cases where the result is a decimal.
func TestCalculateDivision(t *testing.T) {
	tests := []struct {
		name          string
		op            calculator.Operation
		expectedValue float64
		expectedExpr  string
	}{
		{
			name:          "integer operands exact result",
			op:            calculator.Operation{Left: 10, Operator: "/", Right: 2},
			expectedValue: 5,
			expectedExpr:  "10 / 2",
		},
		{
			name:          "division resulting in decimal",
			op:            calculator.Operation{Left: 1, Operator: "/", Right: 4},
			expectedValue: 0.25,
			expectedExpr:  "1 / 4",
		},
		{
			name:          "decimal operands",
			op:            calculator.Operation{Left: 7.5, Operator: "/", Right: 2.5},
			expectedValue: 3,
			expectedExpr:  "7.5 / 2.5",
		},
		{
			name:          "negative dividend",
			op:            calculator.Operation{Left: -10, Operator: "/", Right: 2},
			expectedValue: -5,
			expectedExpr:  "-10 / 2",
		},
		{
			name:          "decimal result from integer division",
			op:            calculator.Operation{Left: 7, Operator: "/", Right: 2},
			expectedValue: 3.5,
			expectedExpr:  "7 / 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Value != tt.expectedValue {
				t.Errorf("Calculate(%v).Value = %v, want %v", tt.op, got.Value, tt.expectedValue)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}

// TestCalculateDivisionByZero verifies that Calculate returns an error when
// the right operand is zero and the operator is division.
func TestCalculateDivisionByZero(t *testing.T) {
	op := calculator.Operation{Left: 5, Operator: "/", Right: 0}
	_, err := calculator.Calculate(op)
	if err == nil {
		t.Errorf("Calculate(%v) expected an error for division by zero, got nil", op)
	}
}

// TestCalculateResultStructFields verifies that Calculate populates both
// the Value and Expression fields of the returned Result struct.
func TestCalculateResultStructFields(t *testing.T) {
	tests := []struct {
		name          string
		op            calculator.Operation
		expectedValue float64
		expectedExpr  string
	}{
		{
			name:          "addition result has value and expression",
			op:            calculator.Operation{Left: 2, Operator: "+", Right: 3},
			expectedValue: 5,
			expectedExpr:  "2 + 3",
		},
		{
			name:          "subtraction result has value and expression",
			op:            calculator.Operation{Left: 8, Operator: "-", Right: 5},
			expectedValue: 3,
			expectedExpr:  "8 - 5",
		},
		{
			name:          "multiplication result has value and expression",
			op:            calculator.Operation{Left: 6, Operator: "*", Right: 7},
			expectedValue: 42,
			expectedExpr:  "6 * 7",
		},
		{
			name:          "division result has value and expression",
			op:            calculator.Operation{Left: 9, Operator: "/", Right: 3},
			expectedValue: 3,
			expectedExpr:  "9 / 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Value != tt.expectedValue {
				t.Errorf("Calculate(%v).Value = %v, want %v", tt.op, got.Value, tt.expectedValue)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}

// TestCalculateDecimalOperandsExpression verifies that the expression string
// uses the correct decimal representation (no trailing zeros) for operands.
func TestCalculateDecimalOperandsExpression(t *testing.T) {
	tests := []struct {
		name         string
		op           calculator.Operation
		expectedExpr string
	}{
		{
			name:         "whole number operands have no trailing zeros in expression",
			op:           calculator.Operation{Left: 6.0, Operator: "+", Right: 2.0},
			expectedExpr: "6 + 2",
		},
		{
			name:         "decimal operands preserved in expression",
			op:           calculator.Operation{Left: 1.5, Operator: "+", Right: 0.25},
			expectedExpr: "1.5 + 0.25",
		},
		{
			name:         "negative decimal in expression",
			op:           calculator.Operation{Left: -2.5, Operator: "*", Right: 4},
			expectedExpr: "-2.5 * 4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculator.Calculate(tt.op)
			if err != nil {
				t.Fatalf("Calculate(%v) returned unexpected error: %v", tt.op, err)
			}
			if got.Expression != tt.expectedExpr {
				t.Errorf("Calculate(%v).Expression = %q, want %q", tt.op, got.Expression, tt.expectedExpr)
			}
		})
	}
}
