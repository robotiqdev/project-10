//go:build !integration

package calc

import (
	"errors"
	"math"
	"testing"

	errs "github.com/example/calc-app/internal/errors"
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
		{"mixed sign integers", -3, 4, 1},
		{"decimals", 1.5, 2.5, 4.0},
		{"zero left operand", 0, 5, 5},
		{"zero right operand", 5, 0, 5},
		{"both operands zero", 0, 0, 0},
		{"negative decimal", -1.1, -2.2, -3.3000000000000003},
		{"large positive numbers", 1e12, 2e12, 3e12},
		{"large negative numbers", -1e12, -2e12, -3e12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
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
		{"decimals", 5.5, 2.2, 3.3000000000000003},
		{"zero left operand", 0, 5, -5},
		{"zero right operand", 5, 0, 5},
		{"both operands zero", 0, 0, 0},
		{"large numbers", 1e15, 5e14, 5e14},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Subtract(tt.a, tt.b)
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
		{"fractional result", 0.1, 0.2, 0.020000000000000004},
		{"large numbers", 1e6, 1e6, 1e12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Multiply(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestDivide verifies Divide returns the correct quotient for valid inputs.
func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"normal division", 10, 2, 5},
		{"decimal result", 10, 3, 10.0 / 3.0},
		{"exact integer result", 9, 3, 3},
		{"negative dividend", -10, 2, -5},
		{"negative divisor", 10, -2, -5},
		{"both negative", -10, -2, 5},
		{"zero dividend", 0, 5, 0},
		{"decimal operands", 7.5, 2.5, 3.0},
		{"large numbers", 1e12, 1e6, 1e6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Divide(tt.a, tt.b)
			if err != nil {
				t.Fatalf("Divide(%v, %v) returned unexpected error: %v", tt.a, tt.b, err)
			}
			if math.Abs(got-tt.expected) >= 1e-9 && got != tt.expected {
				t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// TestDivideByZeroReturnsError verifies that dividing by zero returns a non-nil error
// that satisfies errors.Is(err, errs.ErrDivisionByZero).
func TestDivideByZeroReturnsError(t *testing.T) {
	_, err := Divide(5, 0)

	if err == nil {
		t.Fatal("Divide(5, 0) expected an error, got nil")
	}

	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("Divide(5, 0) error = %v; want errors.Is(err, errs.ErrDivisionByZero) to be true", err)
	}
}

// TestDivideByZeroWithNegativeDividend verifies that dividing a negative number by zero
// also returns the ErrDivisionByZero error.
func TestDivideByZeroWithNegativeDividend(t *testing.T) {
	_, err := Divide(-5, 0)

	if err == nil {
		t.Fatal("Divide(-5, 0) expected an error, got nil")
	}

	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("Divide(-5, 0) error = %v; want errors.Is(err, errs.ErrDivisionByZero) to be true", err)
	}
}

// TestDivideByZeroReturnsZeroResult verifies that when dividing by zero,
// the returned result value is 0 (not some undefined float).
func TestDivideByZeroReturnsZeroResult(t *testing.T) {
	result, err := Divide(5, 0)

	if err == nil {
		t.Fatal("Divide(5, 0) expected an error, got nil")
	}

	if result != 0 {
		t.Errorf("Divide(5, 0) result = %v; want 0 when error is returned", result)
	}
}

// TestDivideByZeroMultipleDividends verifies that Divide(x, 0) always returns
// ErrDivisionByZero regardless of the dividend x value.
func TestDivideByZeroMultipleDividends(t *testing.T) {
	tests := []struct {
		name     string
		dividend float64
	}{
		{"positive integer", 1},
		{"positive large integer", 1000000},
		{"negative integer", -7},
		{"negative large integer", -1000000},
		{"zero dividend", 0},
		{"positive decimal", 3.14},
		{"negative decimal", -2.71},
		{"very small positive", 0.0001},
		{"very large positive", 1e15},
		{"very large negative", -1e15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.dividend, 0)

			if err == nil {
				t.Fatalf("Divide(%v, 0) = %v, nil; want an error", tt.dividend, result)
			}
			if !errors.Is(err, errs.ErrDivisionByZero) {
				t.Errorf("Divide(%v, 0) error = %v; want errors.Is(err, ErrDivisionByZero) to be true", tt.dividend, err)
			}
			if result != 0 {
				t.Errorf("Divide(%v, 0) result = %v; want 0 when error is returned", tt.dividend, result)
			}
		})
	}
}
