//go:build !integration

package calculator_test

import (
	"testing"

	"github.com/example/calc-app/calculator"
	"github.com/example/calc-app/internal/calc"
)

// TestCalculateAddition verifies that Calculate correctly adds two operands,
// returning the right Value and Expression in the Result struct.
func TestCalculateAddition(t *testing.T) {
	tests := []struct {
		name          string
		op            calculator.Operation
		expectedValue float64
		expectedExpr  string
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

// TestCalculateAdd verifies Calculate end-to-end for addition with required basic cases.
func TestCalculateAdd(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{"Add(2,3)=5", 2, 3, "5"},
		{"Add(-1,1)=0", -1, 1, "0"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpAdd}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("Calculate(%v + %v) returned unexpected error: %v", tt.left, tt.right, err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v + %v) = %q; want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestCalculateSubtract verifies Calculate end-to-end for subtraction with required basic cases.
func TestCalculateSubtract(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{"Subtract(5,3)=2", 5, 3, "2"},
		{"Subtract(-2,-3)=1", -2, -3, "1"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpSubtract}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("Calculate(%v - %v) returned unexpected error: %v", tt.left, tt.right, err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v - %v) = %q; want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestCalculateMultiply verifies Calculate end-to-end for multiplication with required basic cases.
func TestCalculateMultiply(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{"Multiply(3,4)=12", 3, 4, "12"},
		{"Multiply(0,100)=0", 0, 100, "0"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpMultiply}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("Calculate(%v * %v) returned unexpected error: %v", tt.left, tt.right, err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v * %v) = %q; want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestCalculateDivide verifies Calculate end-to-end for division with required basic cases.
func TestCalculateDivide(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{"Divide(10,4)=2.5", 10, 4, "2.5"},
		{"Divide(-6,2)=-3", -6, 2, "-3"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpDivide}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("Calculate(%v / %v) returned unexpected error: %v", tt.left, tt.right, err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v / %v) = %q; want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// BenchmarkCalculate measures the performance of the Engine.Calculate method
// across all supported arithmetic operations.
func BenchmarkCalculate(b *testing.B) {
	engine := calc.NewEngine()

	b.Run("Add", func(b *testing.B) {
		op := calc.Operation{Left: 2, Right: 3, Operator: calc.OpAdd}
		for i := 0; i < b.N; i++ {
			engine.Calculate(op)
		}
	})

	b.Run("Subtract", func(b *testing.B) {
		op := calc.Operation{Left: 5, Right: 3, Operator: calc.OpSubtract}
		for i := 0; i < b.N; i++ {
			engine.Calculate(op)
		}
	})

	b.Run("Multiply", func(b *testing.B) {
		op := calc.Operation{Left: 3, Right: 4, Operator: calc.OpMultiply}
		for i := 0; i < b.N; i++ {
			engine.Calculate(op)
		}
	})

	b.Run("Divide", func(b *testing.B) {
		op := calc.Operation{Left: 10, Right: 4, Operator: calc.OpDivide}
		for i := 0; i < b.N; i++ {
			engine.Calculate(op)
		}
	})
}
