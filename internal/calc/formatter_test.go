package calc

import "testing"

// TestFormatResult verifies the output contract for FormatResult using table-driven tests.
func TestFormatResult(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{"whole number 6.0 formats without decimal point", 6.0, "6"},
		{"decimal 3.14 preserves precision", 3.14, "3.14"},
		{"negative whole number -3.0 formats without decimal point", -3.0, "-3"},
		{"zero 0.0 formats as '0'", 0.0, "0"},
		{"repeating decimal 1/3 preserves all digits", 0.3333333333333333, "0.3333333333333333"},
		{"large number 1e10 uses decimal notation", 1e10, "10000000000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatResult(tt.input)
			if got != tt.expected {
				t.Errorf("FormatResult(%v) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

// TestFormatResultNoTrailingZeros verifies that whole numbers never have trailing zeros.
func TestFormatResultNoTrailingZeros(t *testing.T) {
	cases := []struct {
		input    float64
		expected string
	}{
		{1.0, "1"},
		{10.0, "10"},
		{100.0, "100"},
		{-5.0, "-5"},
		{0.0, "0"},
	}

	for _, c := range cases {
		got := FormatResult(c.input)
		if got != c.expected {
			t.Errorf("FormatResult(%v): got %q, want %q (no trailing zeros expected)", c.input, got, c.expected)
		}
	}
}

// TestFormatResultDecimalPrecision verifies that decimal precision is preserved
// for non-terminating decimals without unnecessary trailing zeros.
func TestFormatResultDecimalPrecision(t *testing.T) {
	cases := []struct {
		input    float64
		expected string
	}{
		{0.1, "0.1"},
		{0.2, "0.2"},
		{0.3333333333333333, "0.3333333333333333"},
		{1.23456789, "1.23456789"},
		{-2.5, "-2.5"},
	}

	for _, c := range cases {
		got := FormatResult(c.input)
		if got != c.expected {
			t.Errorf("FormatResult(%v): got %q, want %q", c.input, got, c.expected)
		}
	}
}
