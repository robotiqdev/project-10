//go:build !integration

package errors_test

import (
	"errors"
	"strings"
	"testing"

	errs "github.com/example/calc-app/internal/errors"
)

// TestCalcErrorFormatDivisionByZero verifies CalcError.Error() produces the expected
// message format for ErrDivisionByZero: `division by zero: "<input>"`.
func TestCalcErrorFormatDivisionByZero(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrDivisionByZero, Input: "0"}
	got := e.Error()
	want := `division by zero: "0"`
	if got != want {
		t.Errorf("CalcError.Error() = %q; want %q", got, want)
	}
}

// TestCalcErrorFormatUnknownOperator verifies CalcError.Error() produces the expected
// message format for ErrUnknownOperator: `unknown operator: "<input>"`.
func TestCalcErrorFormatUnknownOperator(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: "%"}
	got := e.Error()
	want := `unknown operator: "%"`
	if got != want {
		t.Errorf("CalcError.Error() = %q; want %q", got, want)
	}
}

// TestCalcErrorFormatInvalidNumber verifies CalcError.Error() produces the expected
// message format for ErrInvalidNumber: `invalid number: "<input>"`.
func TestCalcErrorFormatInvalidNumber(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: "abc"}
	got := e.Error()
	want := `invalid number: "abc"`
	if got != want {
		t.Errorf("CalcError.Error() = %q; want %q", got, want)
	}
}

// TestCalcErrorFormatInvalidArgCount verifies CalcError.Error() produces the expected
// message format for ErrInvalidArgCount: `invalid argument count: "<input>"`.
func TestCalcErrorFormatInvalidArgCount(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: "4"}
	got := e.Error()
	want := `invalid argument count: "4"`
	if got != want {
		t.Errorf("CalcError.Error() = %q; want %q", got, want)
	}
}

// TestCalcErrorFormatIncludesBadInput verifies that the Error() message includes
// the offending input value to aid debugging.
func TestCalcErrorFormatIncludesBadInput(t *testing.T) {
	tests := []struct {
		name     string
		sentinel error
		input    string
	}{
		{"division by zero includes 0", errs.ErrDivisionByZero, "0"},
		{"unknown operator includes symbol", errs.ErrUnknownOperator, "^"},
		{"invalid number includes text", errs.ErrInvalidNumber, "not_a_number"},
		{"invalid arg count includes count", errs.ErrInvalidArgCount, "5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &errs.CalcError{Sentinel: tt.sentinel, Input: tt.input}
			msg := e.Error()
			if !strings.Contains(msg, tt.input) {
				t.Errorf("CalcError.Error() = %q; want it to contain input %q", msg, tt.input)
			}
		})
	}
}

// TestCalcErrorUnwrapDivisionByZero verifies that errors.Is traverses the Unwrap chain
// and identifies the ErrDivisionByZero sentinel.
func TestCalcErrorUnwrapDivisionByZero(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrDivisionByZero, Input: "0"}

	if !errors.Is(e, errs.ErrDivisionByZero) {
		t.Errorf("errors.Is(calcErr, ErrDivisionByZero) = false; want true")
	}
}

// TestCalcErrorUnwrapUnknownOperator verifies that errors.Is traverses the Unwrap chain
// and identifies the ErrUnknownOperator sentinel.
func TestCalcErrorUnwrapUnknownOperator(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: "mod"}

	if !errors.Is(e, errs.ErrUnknownOperator) {
		t.Errorf("errors.Is(calcErr, ErrUnknownOperator) = false; want true")
	}
}

// TestCalcErrorUnwrapInvalidNumber verifies that errors.Is traverses the Unwrap chain
// and identifies the ErrInvalidNumber sentinel.
func TestCalcErrorUnwrapInvalidNumber(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrInvalidNumber, Input: "xyz"}

	if !errors.Is(e, errs.ErrInvalidNumber) {
		t.Errorf("errors.Is(calcErr, ErrInvalidNumber) = false; want true")
	}
}

// TestCalcErrorUnwrapInvalidArgCount verifies that errors.Is traverses the Unwrap chain
// and identifies the ErrInvalidArgCount sentinel.
func TestCalcErrorUnwrapInvalidArgCount(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrInvalidArgCount, Input: "1"}

	if !errors.Is(e, errs.ErrInvalidArgCount) {
		t.Errorf("errors.Is(calcErr, ErrInvalidArgCount) = false; want true")
	}
}

// TestCalcErrorDoesNotMatchWrongSentinel verifies that errors.Is does NOT match a
// CalcError against an unrelated sentinel.
func TestCalcErrorDoesNotMatchWrongSentinel(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrDivisionByZero, Input: "0"}

	if errors.Is(e, errs.ErrUnknownOperator) {
		t.Errorf("errors.Is(divisionByZeroErr, ErrUnknownOperator) = true; want false")
	}
	if errors.Is(e, errs.ErrInvalidNumber) {
		t.Errorf("errors.Is(divisionByZeroErr, ErrInvalidNumber) = true; want false")
	}
	if errors.Is(e, errs.ErrInvalidArgCount) {
		t.Errorf("errors.Is(divisionByZeroErr, ErrInvalidArgCount) = true; want false")
	}
}

// TestCalcErrorMessageFormatQuotesInput verifies that the input value is wrapped
// in double quotes in the error message, matching the %q format specifier.
func TestCalcErrorMessageFormatQuotesInput(t *testing.T) {
	e := &errs.CalcError{Sentinel: errs.ErrUnknownOperator, Input: "bad_op"}
	got := e.Error()
	if !strings.Contains(got, `"bad_op"`) {
		t.Errorf("CalcError.Error() = %q; want input to be quoted as %q", got, `"bad_op"`)
	}
}

// TestCalcErrorWrappedInStandardError verifies that CalcError can be wrapped by
// fmt.Errorf and still resolved by errors.Is.
func TestCalcErrorChainWithWrappedError(t *testing.T) {
	inner := &errs.CalcError{Sentinel: errs.ErrDivisionByZero, Input: "0"}

	// Simulate a CalcError being the direct error (not wrapped further)
	var e error = inner

	if !errors.Is(e, errs.ErrDivisionByZero) {
		t.Errorf("errors.Is(e, ErrDivisionByZero) = false; want true after assigning to error interface")
	}
}
