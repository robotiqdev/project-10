//go:build !integration

package calculator_test

import (
	"math"
	"testing"

	"github.com/example/calc-app/calculator"
	"github.com/example/calc-app/internal/calc"
)

// TestAdd verifies Add returns the correct sum for various inputs.
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive integers", 3, 4, 7},
		{"negative integers", -3, -4, -7},
		{"mixed sign: negative plus positive", -3, 4, 1},
		{"mixed sign: positive plus negative", 3, -4, -1},
		{"decimals", 1.5, 2.5, 4.0},
		{"zero left operand", 0, 5, 5},
		{"zero right operand", 5, 0, 5},
		{"both operands zero", 0, 0, 0},
		{"negative decimals", -1.5, -2.5, -4.0},
		{"large integers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestSubtract verifies Subtract returns the correct difference for various inputs.
func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive integers", 10, 3, 7},
		{"negative integers", -10, -3, -7},
		{"mixed sign: positive minus negative", 5, -3, 8},
		{"mixed sign: negative minus positive", -5, 3, -8},
		{"decimals", 5.5, 2.5, 3.0},
		{"zero left operand", 0, 5, -5},
		{"zero right operand", 5, 0, 5},
		{"both operands zero", 0, 0, 0},
		{"negative decimals", -5.5, -2.5, -3.0},
		{"subtract larger from smaller", 3, 10, -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestMultiply verifies Multiply returns the correct product for various inputs.
func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive integers", 3, 4, 12},
		{"negative integers", -3, -4, 12},
		{"mixed sign: positive times negative", 3, -4, -12},
		{"mixed sign: negative times positive", -3, 4, -12},
		{"decimals", 1.5, 2.0, 3.0},
		{"zero left operand", 0, 5, 0},
		{"zero right operand", 5, 0, 0},
		{"both operands zero", 0, 0, 0},
		{"multiply by one", 7, 1, 7},
		{"negative decimal times positive", -2.5, 4.0, -10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestDivide verifies Divide returns the correct quotient for non-zero denominators.
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"normal integer division", 10, 2, 5},
		{"exact integer result", 9, 3, 3},
		{"decimal result", 10, 3, 10.0 / 3.0},
		{"negative dividend", -10, 2, -5},
		{"negative divisor", 10, -2, -5},
		{"both negative", -10, -2, 5},
		{"zero dividend", 0, 5, 0},
		{"decimal operands", 7.5, 2.5, 3.0},
		{"decimal dividend integer divisor", 5.5, 2, 2.75},
		{"integer dividend decimal divisor", 6, 1.5, 4.0},
		{"negative decimals", -7.5, -2.5, 3.0},
		{"divide by one", 42, 1, 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.Divide(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestAddBasicCases verifies Add returns the correct sum for the required basic cases.
func TestAddBasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
		approx   bool
	}{
		{"Add(2,3)=5", 2, 3, 5, false},
		{"Add(-1,1)=0", -1, 1, 0, false},
		{"Add(0.1,0.2)≈0.3", 0.1, 0.2, 0.3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Add(tt.a, tt.b)
			if tt.approx {
				if math.Abs(got-tt.expected) > 1e-9 {
					t.Errorf("Add(%v, %v) = %v; want approximately %v", tt.a, tt.b, got, tt.expected)
				}
			} else {
				if got != tt.expected {
					t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
				}
			}
		})
	}
}

// TestSubtractBasicCases verifies Subtract returns the correct difference for the required basic cases.
func TestSubtractBasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"Subtract(5,3)=2", 5, 3, 2},
		{"Subtract(-2,-3)=1", -2, -3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Subtract(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestMultiplyBasicCases verifies Multiply returns the correct product for the required basic cases.
func TestMultiplyBasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"Multiply(3,4)=12", 3, 4, 12},
		{"Multiply(0,100)=0", 0, 100, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestDivideBasicCases verifies Divide returns the correct quotient for the required basic cases.
func TestDivideBasicCases(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"Divide(10,4)=2.5", 10, 4, 2.5},
		{"Divide(-6,2)=-3", -6, 2, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calc.Divide(tt.a, tt.b)
			if err != nil {
				t.Fatalf("Divide(%v, %v) returned unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// BenchmarkAdd measures the performance of the Add function.
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Add(2, 3)
	}
}

// BenchmarkSubtract measures the performance of the Subtract function.
func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Subtract(5, 3)
	}
}

// BenchmarkMultiply measures the performance of the Multiply function.
func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Multiply(3, 4)
	}
}

// BenchmarkDivide measures the performance of the Divide function.
func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calc.Divide(10, 4)
	}
}
