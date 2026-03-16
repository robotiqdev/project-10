//go:build !integration

package calculator_test

import (
	"math"
	"testing"

	"github.com/example/calc-app/internal/calc"
)

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
