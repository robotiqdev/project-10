//go:build !integration

package calculator_test

import (
	"testing"

	"github.com/example/calc-app/internal/calc"
)

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
