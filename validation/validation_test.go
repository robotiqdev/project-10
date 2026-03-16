package validation_test

import (
	"errors"
	"testing"

	"calculator/validation"
)

// --- ValidateOperator tests ---

// TestValidateOperator_ValidPlus verifies '+' is accepted.
func TestValidateOperator_ValidPlus(t *testing.T) {
	if err := validation.ValidateOperator("+"); err != nil {
		t.Errorf("ValidateOperator(%q) = %v, want nil", "+", err)
	}
}

// TestValidateOperator_ValidMinus verifies '-' is accepted.
func TestValidateOperator_ValidMinus(t *testing.T) {
	if err := validation.ValidateOperator("-"); err != nil {
		t.Errorf("ValidateOperator(%q) = %v, want nil", "-", err)
	}
}

// TestValidateOperator_ValidMultiply verifies '*' is accepted.
func TestValidateOperator_ValidMultiply(t *testing.T) {
	if err := validation.ValidateOperator("*"); err != nil {
		t.Errorf("ValidateOperator(%q) = %v, want nil", "*", err)
	}
}

// TestValidateOperator_ValidDivide verifies '/' is accepted.
func TestValidateOperator_ValidDivide(t *testing.T) {
	if err := validation.ValidateOperator("/"); err != nil {
		t.Errorf("ValidateOperator(%q) = %v, want nil", "/", err)
	}
}

// TestValidateOperator_InvalidPercent verifies '%' returns ErrInvalidOperator.
func TestValidateOperator_InvalidPercent(t *testing.T) {
	err := validation.ValidateOperator("%")
	if err == nil {
		t.Fatalf("ValidateOperator(%q) = nil, want ErrInvalidOperator", "%")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "%", err)
	}
}

// TestValidateOperator_InvalidCaret verifies '^' returns ErrInvalidOperator.
func TestValidateOperator_InvalidCaret(t *testing.T) {
	err := validation.ValidateOperator("^")
	if err == nil {
		t.Fatalf("ValidateOperator(%q) = nil, want ErrInvalidOperator", "^")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "^", err)
	}
}

// TestValidateOperator_InvalidEmptyString verifies empty string returns ErrInvalidOperator.
func TestValidateOperator_InvalidEmptyString(t *testing.T) {
	err := validation.ValidateOperator("")
	if err == nil {
		t.Fatalf("ValidateOperator(%q) = nil, want ErrInvalidOperator", "")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "", err)
	}
}

// TestValidateOperator_InvalidLetter verifies a letter returns ErrInvalidOperator.
func TestValidateOperator_InvalidLetter(t *testing.T) {
	err := validation.ValidateOperator("a")
	if err == nil {
		t.Fatalf("ValidateOperator(%q) = nil, want ErrInvalidOperator", "a")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "a", err)
	}
}

// TestValidateOperator_InvalidWord verifies a word returns ErrInvalidOperator.
func TestValidateOperator_InvalidWord(t *testing.T) {
	err := validation.ValidateOperator("add")
	if err == nil {
		t.Fatalf("ValidateOperator(%q) = nil, want ErrInvalidOperator", "add")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "add", err)
	}
}

// TestValidateOperator_Table runs a table-driven set of valid and invalid cases.
func TestValidateOperator_Table(t *testing.T) {
	tests := []struct {
		input   string
		wantErr bool
	}{
		{"+", false},
		{"-", false},
		{"*", false},
		{"/", false},
		{"%", true},
		{"^", true},
		{"", true},
		{"a", true},
		{"x", true},
		{"add", true},
		{"++", true},
		{" ", true},
	}

	for _, tc := range tests {
		t.Run("op="+tc.input, func(t *testing.T) {
			err := validation.ValidateOperator(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("ValidateOperator(%q) = nil, want error", tc.input)
				}
				if !errors.Is(err, validation.ErrInvalidOperator) {
					t.Errorf("ValidateOperator(%q) error = %v, want errors.Is ErrInvalidOperator", tc.input, err)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateOperator(%q) = %v, want nil", tc.input, err)
				}
			}
		})
	}
}

// --- ExtractOperator tests ---

// TestExtractOperator_ValidPlus verifies '+' is returned unchanged.
func TestExtractOperator_ValidPlus(t *testing.T) {
	op, err := validation.ExtractOperator("+")
	if err != nil {
		t.Fatalf("ExtractOperator(%q) error = %v, want nil", "+", err)
	}
	if op != "+" {
		t.Errorf("ExtractOperator(%q) = %q, want %q", "+", op, "+")
	}
}

// TestExtractOperator_ValidMinus verifies '-' is returned unchanged.
func TestExtractOperator_ValidMinus(t *testing.T) {
	op, err := validation.ExtractOperator("-")
	if err != nil {
		t.Fatalf("ExtractOperator(%q) error = %v, want nil", "-", err)
	}
	if op != "-" {
		t.Errorf("ExtractOperator(%q) = %q, want %q", "-", op, "-")
	}
}

// TestExtractOperator_ValidMultiply verifies '*' is returned unchanged.
func TestExtractOperator_ValidMultiply(t *testing.T) {
	op, err := validation.ExtractOperator("*")
	if err != nil {
		t.Fatalf("ExtractOperator(%q) error = %v, want nil", "*", err)
	}
	if op != "*" {
		t.Errorf("ExtractOperator(%q) = %q, want %q", "*", op, "*")
	}
}

// TestExtractOperator_ValidDivide verifies '/' is returned unchanged.
func TestExtractOperator_ValidDivide(t *testing.T) {
	op, err := validation.ExtractOperator("/")
	if err != nil {
		t.Fatalf("ExtractOperator(%q) error = %v, want nil", "/", err)
	}
	if op != "/" {
		t.Errorf("ExtractOperator(%q) = %q, want %q", "/", op, "/")
	}
}

// TestExtractOperator_InvalidPercent verifies '%' returns empty string and ErrInvalidOperator.
func TestExtractOperator_InvalidPercent(t *testing.T) {
	op, err := validation.ExtractOperator("%")
	if err == nil {
		t.Fatalf("ExtractOperator(%q) = nil error, want ErrInvalidOperator", "%")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ExtractOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "%", err)
	}
	if op != "" {
		t.Errorf("ExtractOperator(%q) op = %q, want empty string on error", "%", op)
	}
}

// TestExtractOperator_InvalidCaret verifies '^' returns empty string and ErrInvalidOperator.
func TestExtractOperator_InvalidCaret(t *testing.T) {
	op, err := validation.ExtractOperator("^")
	if err == nil {
		t.Fatalf("ExtractOperator(%q) = nil error, want ErrInvalidOperator", "^")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ExtractOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "^", err)
	}
	if op != "" {
		t.Errorf("ExtractOperator(%q) op = %q, want empty string on error", "^", op)
	}
}

// TestExtractOperator_InvalidEmptyString verifies empty string returns ErrInvalidOperator.
func TestExtractOperator_InvalidEmptyString(t *testing.T) {
	op, err := validation.ExtractOperator("")
	if err == nil {
		t.Fatalf("ExtractOperator(%q) = nil error, want ErrInvalidOperator", "")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ExtractOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "", err)
	}
	if op != "" {
		t.Errorf("ExtractOperator(%q) op = %q, want empty string on error", "", op)
	}
}

// TestExtractOperator_InvalidLetter verifies a letter returns ErrInvalidOperator.
func TestExtractOperator_InvalidLetter(t *testing.T) {
	op, err := validation.ExtractOperator("a")
	if err == nil {
		t.Fatalf("ExtractOperator(%q) = nil error, want ErrInvalidOperator", "a")
	}
	if !errors.Is(err, validation.ErrInvalidOperator) {
		t.Errorf("ExtractOperator(%q) error = %v, want errors.Is ErrInvalidOperator", "a", err)
	}
	if op != "" {
		t.Errorf("ExtractOperator(%q) op = %q, want empty string on error", "a", op)
	}
}

// TestExtractOperator_Table runs a table-driven set of valid and invalid cases.
func TestExtractOperator_Table(t *testing.T) {
	tests := []struct {
		input   string
		wantOp  string
		wantErr bool
	}{
		{"+", "+", false},
		{"-", "-", false},
		{"*", "*", false},
		{"/", "/", false},
		{"%", "", true},
		{"^", "", true},
		{"", "", true},
		{"a", "", true},
		{"x", "", true},
		{"add", "", true},
		{"++", "", true},
		{" ", "", true},
	}

	for _, tc := range tests {
		t.Run("op="+tc.input, func(t *testing.T) {
			op, err := validation.ExtractOperator(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("ExtractOperator(%q) = nil error, want error", tc.input)
				}
				if !errors.Is(err, validation.ErrInvalidOperator) {
					t.Errorf("ExtractOperator(%q) error = %v, want errors.Is ErrInvalidOperator", tc.input, err)
				}
				if op != "" {
					t.Errorf("ExtractOperator(%q) op = %q, want empty string on error", tc.input, op)
				}
			} else {
				if err != nil {
					t.Errorf("ExtractOperator(%q) error = %v, want nil", tc.input, err)
				}
				if op != tc.wantOp {
					t.Errorf("ExtractOperator(%q) = %q, want %q", tc.input, op, tc.wantOp)
				}
			}
		})
	}
}
