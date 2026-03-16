//go:build !integration

package parser_test

import (
	"errors"
	"strings"
	"testing"

	errs "github.com/example/calc-app/internal/errors"
	"github.com/example/calc-app/internal/parser"
)

// TestValidateOperatorAcceptsValidOperators verifies that the four supported
// operators (+, -, *, /) produce no error.
func TestValidateOperatorAcceptsValidOperators(t *testing.T) {
	validOps := []string{"+", "-", "*", "/"}

	for _, op := range validOps {
		t.Run("operator "+op, func(t *testing.T) {
			if err := parser.ValidateOperator(op); err != nil {
				t.Errorf("ValidateOperator(%q) = %v; want nil", op, err)
			}
		})
	}
}

// TestValidateOperatorRejectsPercentSign verifies that "%" returns ErrUnknownOperator.
func TestValidateOperatorRejectsPercentSign(t *testing.T) {
	err := parser.ValidateOperator("%")

	if err == nil {
		t.Fatal(`ValidateOperator("%") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrUnknownOperator) {
		t.Errorf(`ValidateOperator("%%") error = %v; want errors.Is(err, ErrUnknownOperator) to be true`, err)
	}
}

// TestValidateOperatorRejectsCaretSign verifies that "^" returns ErrUnknownOperator.
func TestValidateOperatorRejectsCaretSign(t *testing.T) {
	err := parser.ValidateOperator("^")

	if err == nil {
		t.Fatal(`ValidateOperator("^") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrUnknownOperator) {
		t.Errorf(`ValidateOperator("^") error = %v; want errors.Is(err, ErrUnknownOperator) to be true`, err)
	}
}

// TestValidateOperatorRejectsEmptyString verifies that "" returns ErrUnknownOperator.
func TestValidateOperatorRejectsEmptyString(t *testing.T) {
	err := parser.ValidateOperator("")

	if err == nil {
		t.Fatal(`ValidateOperator("") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrUnknownOperator) {
		t.Errorf(`ValidateOperator("") error = %v; want errors.Is(err, ErrUnknownOperator) to be true`, err)
	}
}

// TestValidateOperatorRejectsWordOperators verifies that word-form operators
// (mod, div, add, sub, mul) are rejected with ErrUnknownOperator.
func TestValidateOperatorRejectsWordOperators(t *testing.T) {
	wordOps := []string{"mod", "div", "add", "sub", "mul", "plus"}

	for _, op := range wordOps {
		t.Run("word op "+op, func(t *testing.T) {
			err := parser.ValidateOperator(op)

			if err == nil {
				t.Fatalf("ValidateOperator(%q) = nil; want non-nil error", op)
			}
			if !errors.Is(err, errs.ErrUnknownOperator) {
				t.Errorf("ValidateOperator(%q) error = %v; want errors.Is(err, ErrUnknownOperator) to be true", op, err)
			}
		})
	}
}

// TestValidateOperatorRejectsMiscSymbols verifies various other symbols are rejected.
func TestValidateOperatorRejectsMiscSymbols(t *testing.T) {
	symbols := []string{"!", "@", "#", "$", "~", "|", "\\", "?", "=", "<", ">"}

	for _, op := range symbols {
		t.Run("symbol "+op, func(t *testing.T) {
			err := parser.ValidateOperator(op)

			if err == nil {
				t.Fatalf("ValidateOperator(%q) = nil; want non-nil error", op)
			}
			if !errors.Is(err, errs.ErrUnknownOperator) {
				t.Errorf("ValidateOperator(%q) error = %v; want errors.Is(err, ErrUnknownOperator) to be true", op, err)
			}
		})
	}
}

// TestValidateOperatorErrorContainsInput verifies that the error message includes
// the invalid operator value for debugging.
func TestValidateOperatorErrorContainsInput(t *testing.T) {
	invalidOps := []string{"%", "^", "mod", "!"}

	for _, op := range invalidOps {
		t.Run("contains "+op, func(t *testing.T) {
			err := parser.ValidateOperator(op)
			if err == nil {
				t.Fatalf("ValidateOperator(%q) = nil; want non-nil error", op)
			}
			if !strings.Contains(err.Error(), op) {
				t.Errorf("ValidateOperator(%q).Error() = %q; want it to contain %q", op, err.Error(), op)
			}
		})
	}
}

// TestValidateNumberAcceptsValidFormats verifies that well-formed number strings
// produce no error.
func TestValidateNumberAcceptsValidFormats(t *testing.T) {
	valid := []string{
		"0", "1", "42", "-1", "-42",
		"3.14", "-3.14", "1.0", "0.5",
		"1e10", "1E10", "1.5e3", "-2.5e-4",
		"100", "999999",
	}

	for _, s := range valid {
		t.Run("valid "+s, func(t *testing.T) {
			if err := parser.ValidateNumber(s); err != nil {
				t.Errorf("ValidateNumber(%q) = %v; want nil", s, err)
			}
		})
	}
}

// TestValidateNumberRejectsAlphaString verifies that a purely alphabetic string
// returns ErrInvalidNumber.
func TestValidateNumberRejectsAlphaString(t *testing.T) {
	err := parser.ValidateNumber("abc")

	if err == nil {
		t.Fatal(`ValidateNumber("abc") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrInvalidNumber) {
		t.Errorf(`ValidateNumber("abc") error = %v; want errors.Is(err, ErrInvalidNumber) to be true`, err)
	}
}

// TestValidateNumberRejectsEmptyString verifies that an empty string returns
// ErrInvalidNumber.
func TestValidateNumberRejectsEmptyString(t *testing.T) {
	err := parser.ValidateNumber("")

	if err == nil {
		t.Fatal(`ValidateNumber("") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrInvalidNumber) {
		t.Errorf(`ValidateNumber("") error = %v; want errors.Is(err, ErrInvalidNumber) to be true`, err)
	}
}

// TestValidateNumberRejectsMultipleDecimalPoints verifies that "1.2.3" returns
// ErrInvalidNumber.
func TestValidateNumberRejectsMultipleDecimalPoints(t *testing.T) {
	err := parser.ValidateNumber("1.2.3")

	if err == nil {
		t.Fatal(`ValidateNumber("1.2.3") = nil; want non-nil error`)
	}
	if !errors.Is(err, errs.ErrInvalidNumber) {
		t.Errorf(`ValidateNumber("1.2.3") error = %v; want errors.Is(err, ErrInvalidNumber) to be true`, err)
	}
}

// TestValidateNumberRejectsMixedAlphaNumeric verifies that mixed strings like "1a2"
// return ErrInvalidNumber.
func TestValidateNumberRejectsMixedAlphaNumeric(t *testing.T) {
	mixed := []string{"1a", "a1", "1a2", "12b", "1.2a", "e", "1e"}

	for _, s := range mixed {
		t.Run("mixed "+s, func(t *testing.T) {
			err := parser.ValidateNumber(s)

			if err == nil {
				t.Fatalf("ValidateNumber(%q) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("ValidateNumber(%q) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
		})
	}
}

// TestValidateNumberRejectsOperatorSymbols verifies that operator symbols are not
// valid numbers.
func TestValidateNumberRejectsOperatorSymbols(t *testing.T) {
	symbols := []string{"+", "-", "*", "/", "++", "--"}

	for _, s := range symbols {
		t.Run("symbol "+s, func(t *testing.T) {
			err := parser.ValidateNumber(s)

			if err == nil {
				t.Fatalf("ValidateNumber(%q) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("ValidateNumber(%q) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
		})
	}
}

// TestValidateNumberErrorContainsInput verifies that the error message includes
// the invalid number string for debugging.
func TestValidateNumberErrorContainsInput(t *testing.T) {
	invalidInputs := []string{"abc", "1.2.3", "1a2", "not_a_number"}

	for _, s := range invalidInputs {
		t.Run("contains "+s, func(t *testing.T) {
			err := parser.ValidateNumber(s)
			if err == nil {
				t.Fatalf("ValidateNumber(%q) = nil; want non-nil error", s)
			}
			if !strings.Contains(err.Error(), s) {
				t.Errorf("ValidateNumber(%q).Error() = %q; want it to contain %q", s, err.Error(), s)
			}
		})
	}
}

// TestValidateArgCountAcceptsThreeArgs verifies that exactly 3 arguments (num1, op, num2)
// produces no error.
func TestValidateArgCountAcceptsThreeArgs(t *testing.T) {
	args := []string{"5", "+", "3"}
	if err := parser.ValidateArgCount(args); err != nil {
		t.Errorf("ValidateArgCount(%v) = %v; want nil", args, err)
	}
}

// TestValidateArgCountRejectsZeroArgs verifies that 0 arguments returns ErrInvalidArgCount.
func TestValidateArgCountRejectsZeroArgs(t *testing.T) {
	args := []string{}
	err := parser.ValidateArgCount(args)

	if err == nil {
		t.Fatal("ValidateArgCount([]) = nil; want non-nil error")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount([]) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", err)
	}
}

// TestValidateArgCountRejectsOneArg verifies that 1 argument returns ErrInvalidArgCount.
func TestValidateArgCountRejectsOneArg(t *testing.T) {
	args := []string{"5"}
	err := parser.ValidateArgCount(args)

	if err == nil {
		t.Fatal("ValidateArgCount([5]) = nil; want non-nil error")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount([5]) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", err)
	}
}

// TestValidateArgCountRejectsTwoArgs verifies that 2 arguments returns ErrInvalidArgCount.
func TestValidateArgCountRejectsTwoArgs(t *testing.T) {
	args := []string{"5", "+"}
	err := parser.ValidateArgCount(args)

	if err == nil {
		t.Fatal("ValidateArgCount([5, +]) = nil; want non-nil error")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount([5, +]) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", err)
	}
}

// TestValidateArgCountRejectsFourArgs verifies that 4 arguments returns ErrInvalidArgCount.
func TestValidateArgCountRejectsFourArgs(t *testing.T) {
	args := []string{"5", "+", "3", "extra"}
	err := parser.ValidateArgCount(args)

	if err == nil {
		t.Fatal("ValidateArgCount([5, +, 3, extra]) = nil; want non-nil error")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount([5, +, 3, extra]) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", err)
	}
}

// TestValidateArgCountRejectsFiveArgs verifies that 5 arguments returns ErrInvalidArgCount.
func TestValidateArgCountRejectsFiveArgs(t *testing.T) {
	args := []string{"5", "+", "3", "extra1", "extra2"}
	err := parser.ValidateArgCount(args)

	if err == nil {
		t.Fatal("ValidateArgCount(5 args) = nil; want non-nil error")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(5 args) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", err)
	}
}

// TestValidateArgCountErrorContainsCount verifies that the error message includes
// the actual argument count provided, for debugging.
func TestValidateArgCountErrorContainsCount(t *testing.T) {
	tests := []struct {
		args          []string
		expectedCount string
	}{
		{[]string{}, "0"},
		{[]string{"a"}, "1"},
		{[]string{"a", "b"}, "2"},
		{[]string{"a", "b", "c", "d"}, "4"},
		{[]string{"a", "b", "c", "d", "e"}, "5"},
	}

	for _, tt := range tests {
		t.Run("count "+tt.expectedCount, func(t *testing.T) {
			err := parser.ValidateArgCount(tt.args)
			if err == nil {
				t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", tt.args)
			}
			if !strings.Contains(err.Error(), tt.expectedCount) {
				t.Errorf("ValidateArgCount(%v).Error() = %q; want it to contain count %q",
					tt.args, err.Error(), tt.expectedCount)
			}
		})
	}
}
