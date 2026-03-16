//go:build !integration

package calc_test

import (
	"errors"
	"testing"

	"github.com/example/calc-app/internal/calc"
	errs "github.com/example/calc-app/internal/errors"
)

// TestEngineCalculateAdd verifies that Engine.Calculate correctly adds two numbers.
func TestEngineCalculateAdd(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{name: "positive integers", left: 3, right: 4, expected: "7"},
		{name: "zero plus zero", left: 0, right: 0, expected: "0"},
		{name: "negative plus positive", left: -5, right: 10, expected: "5"},
		{name: "negative plus negative", left: -3, right: -7, expected: "-10"},
		{name: "large integers", left: 1000, right: 2000, expected: "3000"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpAdd}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v + %v) = %q, want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestEngineCalculateSubtract verifies that Engine.Calculate correctly subtracts two numbers.
func TestEngineCalculateSubtract(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{name: "positive integers", left: 10, right: 3, expected: "7"},
		{name: "zero minus zero", left: 0, right: 0, expected: "0"},
		{name: "negative minus positive", left: -5, right: 3, expected: "-8"},
		{name: "positive minus negative", left: 5, right: -3, expected: "8"},
		{name: "result is zero", left: 7, right: 7, expected: "0"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpSubtract}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v - %v) = %q, want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestEngineCalculateMultiply verifies that Engine.Calculate correctly multiplies two numbers.
func TestEngineCalculateMultiply(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{name: "positive integers", left: 3, right: 4, expected: "12"},
		{name: "multiply by zero", left: 5, right: 0, expected: "0"},
		{name: "negative times positive", left: -3, right: 4, expected: "-12"},
		{name: "negative times negative", left: -3, right: -4, expected: "12"},
		{name: "multiply by one", left: 7, right: 1, expected: "7"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpMultiply}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v * %v) = %q, want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestEngineCalculateDivide verifies that Engine.Calculate correctly divides two numbers.
func TestEngineCalculateDivide(t *testing.T) {
	tests := []struct {
		name     string
		left     float64
		right    float64
		expected string
	}{
		{name: "positive integers", left: 12, right: 4, expected: "3"},
		{name: "divide by one", left: 7, right: 1, expected: "7"},
		{name: "negative divided by positive", left: -12, right: 4, expected: "-3"},
		{name: "negative divided by negative", left: -12, right: -4, expected: "3"},
		{name: "zero divided by nonzero", left: 0, right: 5, expected: "0"},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: tt.left, Right: tt.right, Operator: calc.OpDivide}
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calculate(%v / %v) = %q, want %q", tt.left, tt.right, result, tt.expected)
			}
		})
	}
}

// TestEngineCalculateDivideByZeroReturnsError verifies that dividing by zero returns
// an error that wraps ErrDivisionByZero.
func TestEngineCalculateDivideByZeroReturnsError(t *testing.T) {
	engine := calc.NewEngine()
	op := calc.Operation{Left: 10, Right: 0, Operator: calc.OpDivide}

	result, err := engine.Calculate(op)
	if err == nil {
		t.Fatal("expected an error for division by zero, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("expected error to wrap ErrDivisionByZero, got: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty string result on error, got %q", result)
	}
}

// TestEngineCalculateDivideByZeroNegativeDividend verifies that dividing a negative
// number by zero also returns an error wrapping ErrDivisionByZero.
func TestEngineCalculateDivideByZeroNegativeDividend(t *testing.T) {
	engine := calc.NewEngine()
	op := calc.Operation{Left: -5, Right: 0, Operator: calc.OpDivide}

	_, err := engine.Calculate(op)
	if err == nil {
		t.Fatal("expected an error for division by zero, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("expected error to wrap ErrDivisionByZero, got: %v", err)
	}
}

// TestEngineCalculateUnknownOperatorReturnsError verifies that an unrecognized
// operator returns an error wrapping ErrUnknownOperator.
func TestEngineCalculateUnknownOperatorReturnsError(t *testing.T) {
	tests := []struct {
		name     string
		operator calc.Operator
	}{
		{name: "percent operator", operator: calc.Operator("%")},
		{name: "caret operator", operator: calc.Operator("^")},
		{name: "empty operator", operator: calc.Operator("")},
		{name: "word operator", operator: calc.Operator("mod")},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := calc.Operation{Left: 5, Right: 3, Operator: tt.operator}
			result, err := engine.Calculate(op)
			if err == nil {
				t.Fatalf("expected an error for unknown operator %q, got nil", tt.operator)
			}
			if !errors.Is(err, errs.ErrUnknownOperator) {
				t.Errorf("expected error to wrap ErrUnknownOperator, got: %v", err)
			}
			if result != "" {
				t.Errorf("expected empty string result on error, got %q", result)
			}
		})
	}
}

// TestEngineCalculateResultMatchesFormatResult verifies that the string returned
// by Calculate matches what FormatResult would produce for each operation.
func TestEngineCalculateResultMatchesFormatResult(t *testing.T) {
	tests := []struct {
		name     string
		op       calc.Operation
		expected string
	}{
		{
			name:     "add integers produces integer format",
			op:       calc.Operation{Left: 2, Right: 3, Operator: calc.OpAdd},
			expected: calc.FormatResult(5),
		},
		{
			name:     "subtract integers produces integer format",
			op:       calc.Operation{Left: 10, Right: 4, Operator: calc.OpSubtract},
			expected: calc.FormatResult(6),
		},
		{
			name:     "multiply integers produces integer format",
			op:       calc.Operation{Left: 4, Right: 5, Operator: calc.OpMultiply},
			expected: calc.FormatResult(20),
		},
		{
			name:     "divide integers produces integer format",
			op:       calc.Operation{Left: 8, Right: 4, Operator: calc.OpDivide},
			expected: calc.FormatResult(2),
		},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := engine.Calculate(tt.op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Calculate result = %q, want FormatResult output %q", result, tt.expected)
			}
		})
	}
}

// TestNewEngineReturnsNonNilEngine verifies that NewEngine() returns a usable Engine.
func TestNewEngineReturnsNonNilEngine(t *testing.T) {
	engine := calc.NewEngine()
	if engine == nil {
		t.Fatal("NewEngine() returned nil")
	}
}

// BenchmarkEngineCalculate measures the performance of Engine.Calculate.
func BenchmarkEngineCalculate(b *testing.B) {
	engine := calc.NewEngine()
	op := calc.Operation{Left: 10, Right: 3, Operator: calc.OpAdd}
	for i := 0; i < b.N; i++ {
		engine.Calculate(op)
	}
}

// TestEngineCalculateNoErrorOnValidOperations verifies no error is returned for
// all four valid operators.
func TestEngineCalculateNoErrorOnValidOperations(t *testing.T) {
	tests := []struct {
		name string
		op   calc.Operation
	}{
		{name: "addition", op: calc.Operation{Left: 1, Right: 2, Operator: calc.OpAdd}},
		{name: "subtraction", op: calc.Operation{Left: 5, Right: 3, Operator: calc.OpSubtract}},
		{name: "multiplication", op: calc.Operation{Left: 3, Right: 4, Operator: calc.OpMultiply}},
		{name: "division", op: calc.Operation{Left: 10, Right: 2, Operator: calc.OpDivide}},
	}

	engine := calc.NewEngine()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := engine.Calculate(tt.op)
			if err != nil {
				t.Errorf("unexpected error for operator %q: %v", tt.op.Operator, err)
			}
		})
	}
}
