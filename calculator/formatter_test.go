package calculator_test

import (
	"testing"

	"github.com/example/calc-app/calculator"
)

// TestFormatResultIntegerValues verifies that integer-valued results are
// formatted without a decimal point (e.g. 6.0 → "6").
func TestFormatResultIntegerValues(t *testing.T) {
	tests := []struct {
		name     string
		result   calculator.Result
		expected string
	}{
		{"positive integer 6", calculator.Result{Value: 6.0}, "6"},
		{"positive integer 1", calculator.Result{Value: 1.0}, "1"},
		{"positive integer 100", calculator.Result{Value: 100.0}, "100"},
		{"large integer 10000", calculator.Result{Value: 10000.0}, "10000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.FormatResult(tt.result)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q", tt.result, got, tt.expected)
			}
		})
	}
}

// TestFormatResultDecimalValues verifies that decimal results are formatted
// with the minimum number of digits needed (e.g. 3.5 → "3.5").
func TestFormatResultDecimalValues(t *testing.T) {
	tests := []struct {
		name     string
		result   calculator.Result
		expected string
	}{
		{"half 3.5", calculator.Result{Value: 3.5}, "3.5"},
		{"decimal 0.1", calculator.Result{Value: 0.1}, "0.1"},
		{"decimal 0.25", calculator.Result{Value: 0.25}, "0.25"},
		{"decimal 1.23456789", calculator.Result{Value: 1.23456789}, "1.23456789"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.FormatResult(tt.result)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q", tt.result, got, tt.expected)
			}
		})
	}
}

// TestFormatResultNegativeValues verifies that negative results are formatted
// correctly with a leading minus sign and no unnecessary trailing zeros.
func TestFormatResultNegativeValues(t *testing.T) {
	tests := []struct {
		name     string
		result   calculator.Result
		expected string
	}{
		{"negative integer -3", calculator.Result{Value: -3.0}, "-3"},
		{"negative integer -10", calculator.Result{Value: -10.0}, "-10"},
		{"negative decimal -2.5", calculator.Result{Value: -2.5}, "-2.5"},
		{"negative decimal -0.75", calculator.Result{Value: -0.75}, "-0.75"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.FormatResult(tt.result)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q", tt.result, got, tt.expected)
			}
		})
	}
}

// TestFormatResultZero verifies that zero is formatted as "0".
func TestFormatResultZero(t *testing.T) {
	result := calculator.Result{Value: 0.0}
	got := calculator.FormatResult(result)
	if got != "0" {
		t.Errorf("FormatResult(%v) = %q, want %q", result, got, "0")
	}
}

// TestFormatResultNoTrailingZeros verifies that the formatter strips trailing
// zeros from whole-number floats (e.g. 6.0 becomes "6", not "6.0").
func TestFormatResultNoTrailingZeros(t *testing.T) {
	tests := []struct {
		name     string
		result   calculator.Result
		expected string
	}{
		{"6.0 should not have trailing zero", calculator.Result{Value: 6.0}, "6"},
		{"10.0 should not have trailing zeros", calculator.Result{Value: 10.0}, "10"},
		{"-5.0 should not have trailing zero", calculator.Result{Value: -5.0}, "-5"},
		{"0.0 should be just zero", calculator.Result{Value: 0.0}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.FormatResult(tt.result)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q (trailing zeros must be stripped)", tt.result, got, tt.expected)
			}
		})
	}
}

// TestFormatResultIgnoresExpression verifies that the Expression field on
// the Result struct does not affect the formatted value output.
func TestFormatResultIgnoresExpression(t *testing.T) {
	tests := []struct {
		name     string
		result   calculator.Result
		expected string
	}{
		{
			"value 7 with expression set",
			calculator.Result{Value: 7.0, Expression: "3 + 4 = 7"},
			"7",
		},
		{
			"decimal value with expression set",
			calculator.Result{Value: 3.5, Expression: "7 / 2 = 3.5"},
			"3.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculator.FormatResult(tt.result)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q", tt.result, got, tt.expected)
			}
		})
	}
}
