package parser_test

import (
	"errors"
	"testing"

	errs "calculator/internal/errors"
	"calculator/internal/parser"
)

// TestValidateNumber_ValidInputs verifies that well-formed number strings return nil.
func TestValidateNumber_ValidInputs(t *testing.T) {
	validCases := []struct {
		name  string
		input string
	}{
		{"positive integer", "3"},
		{"negative integer", "-3"},
		{"positive decimal", "3.14"},
		{"negative decimal", "-0.5"},
		{"zero", "0"},
		{"large positive integer", "123456"},
		{"large negative decimal", "-999.999"},
		{"positive zero decimal", "0.0"},
	}

	for _, tc := range validCases {
		t.Run(tc.name, func(t *testing.T) {
			err := parser.ValidateNumber(tc.input)
			if err != nil {
				t.Errorf("ValidateNumber(%q) returned unexpected error: %v", tc.input, err)
			}
		})
	}
}

// TestValidateNumber_InvalidInputs verifies that malformed number strings return an error
// wrapping errs.ErrInvalidNumber.
func TestValidateNumber_InvalidInputs(t *testing.T) {
	invalidCases := []struct {
		name  string
		input string
	}{
		{"alphabetic string", "abc"},
		{"empty string", ""},
		{"trailing dot", "3."},
		{"invalid exponent suffix", "1e2x"},
		{"double negative", "--3"},
		{"letters with numbers", "12abc"},
		{"leading plus sign with text", "+abc"},
		{"just a dot", "."},
	}

	for _, tc := range invalidCases {
		t.Run(tc.name, func(t *testing.T) {
			err := parser.ValidateNumber(tc.input)
			if err == nil {
				t.Errorf("ValidateNumber(%q) expected error, got nil", tc.input)
				return
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("ValidateNumber(%q): got error %v, want errors.Is(err, ErrInvalidNumber) == true", tc.input, err)
			}
		})
	}
}

// TestValidateNumber_ReturnsCalcError verifies that the returned error is a *errs.CalcError
// and contains the input string.
func TestValidateNumber_ReturnsCalcError(t *testing.T) {
	input := "abc"
	err := parser.ValidateNumber(input)
	if err == nil {
		t.Fatalf("ValidateNumber(%q) expected error, got nil", input)
	}

	var calcErr *errs.CalcError
	if !errors.As(err, &calcErr) {
		t.Fatalf("ValidateNumber(%q): expected *errs.CalcError, got %T", input, err)
	}
	if calcErr.Input != input {
		t.Errorf("CalcError.Input = %q, want %q", calcErr.Input, input)
	}
	if calcErr.Sentinel != errs.ErrInvalidNumber {
		t.Errorf("CalcError.Sentinel = %v, want ErrInvalidNumber", calcErr.Sentinel)
	}
}

// TestParseNumber_ValidInputs verifies that valid number strings return the correct float64.
func TestParseNumber_ValidInputs(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected float64
	}{
		{"positive integer", "3", 3.0},
		{"negative integer", "-3", -3.0},
		{"positive decimal", "3.14", 3.14},
		{"negative decimal", "-0.5", -0.5},
		{"zero", "0", 0.0},
		{"large value", "123456", 123456.0},
		{"small negative decimal", "-999.999", -999.999},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parser.ParseNumber(tc.input)
			if err != nil {
				t.Fatalf("ParseNumber(%q) returned unexpected error: %v", tc.input, err)
			}
			if got != tc.expected {
				t.Errorf("ParseNumber(%q) = %v, want %v", tc.input, got, tc.expected)
			}
		})
	}
}

// TestParseNumber_InvalidInputs verifies that invalid strings return an error
// wrapping errs.ErrInvalidNumber and a zero float64.
func TestParseNumber_InvalidInputs(t *testing.T) {
	invalidCases := []struct {
		name  string
		input string
	}{
		{"alphabetic string", "abc"},
		{"empty string", ""},
		{"trailing dot", "3."},
		{"invalid exponent suffix", "1e2x"},
		{"double negative", "--3"},
	}

	for _, tc := range invalidCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parser.ParseNumber(tc.input)
			if err == nil {
				t.Errorf("ParseNumber(%q) expected error, got nil (value=%v)", tc.input, got)
				return
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("ParseNumber(%q): got error %v, want errors.Is(err, ErrInvalidNumber) == true", tc.input, err)
			}
		})
	}
}

// TestParseNumber_ErrorIsCalcError verifies that ParseNumber's error is a *errs.CalcError.
func TestParseNumber_ErrorIsCalcError(t *testing.T) {
	input := "--3"
	_, err := parser.ParseNumber(input)
	if err == nil {
		t.Fatalf("ParseNumber(%q) expected error, got nil", input)
	}

	var calcErr *errs.CalcError
	if !errors.As(err, &calcErr) {
		t.Fatalf("ParseNumber(%q): expected *errs.CalcError, got %T", input, err)
	}
	if calcErr.Input != input {
		t.Errorf("CalcError.Input = %q, want %q", calcErr.Input, input)
	}
	if calcErr.Sentinel != errs.ErrInvalidNumber {
		t.Errorf("CalcError.Sentinel = %v, want ErrInvalidNumber", calcErr.Sentinel)
	}
}

// TestParseNumber_ReturnValueOnSuccess ensures the returned float64 matches strconv.ParseFloat.
func TestParseNumber_ReturnValueOnSuccess(t *testing.T) {
	// Verify floating point precision is preserved
	input := "3.14159265358979"
	got, err := parser.ParseNumber(input)
	if err != nil {
		t.Fatalf("ParseNumber(%q) unexpected error: %v", input, err)
	}
	// The value should be exactly what strconv.ParseFloat would return
	const expected = 3.14159265358979
	if got != expected {
		t.Errorf("ParseNumber(%q) = %.15f, want %.15f", input, got, expected)
	}
}
