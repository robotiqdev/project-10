package errors_test

import (
	"errors"
	"strings"
	"testing"

	calcerrors "calculator/internal/errors"
)

// TestSentinelErrors_Defined verifies all four sentinel errors are non-nil package-level vars.
func TestSentinelErrors_Defined(t *testing.T) {
	if calcerrors.ErrUnknownOperator == nil {
		t.Error("ErrUnknownOperator should not be nil")
	}
	if calcerrors.ErrInvalidNumber == nil {
		t.Error("ErrInvalidNumber should not be nil")
	}
	if calcerrors.ErrDivisionByZero == nil {
		t.Error("ErrDivisionByZero should not be nil")
	}
	if calcerrors.ErrInvalidArgCount == nil {
		t.Error("ErrInvalidArgCount should not be nil")
	}
}

// TestSentinelErrors_Distinct verifies that the four sentinels are distinct errors.
func TestSentinelErrors_Distinct(t *testing.T) {
	sentinels := []error{
		calcerrors.ErrUnknownOperator,
		calcerrors.ErrInvalidNumber,
		calcerrors.ErrDivisionByZero,
		calcerrors.ErrInvalidArgCount,
	}
	for i := 0; i < len(sentinels); i++ {
		for j := i + 1; j < len(sentinels); j++ {
			if errors.Is(sentinels[i], sentinels[j]) {
				t.Errorf("sentinel[%d] and sentinel[%d] should be distinct", i, j)
			}
		}
	}
}

// TestCalcError_ErrorsIs_ErrUnknownOperator verifies errors.Is works when CalcError wraps ErrUnknownOperator.
func TestCalcError_ErrorsIs_ErrUnknownOperator(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrUnknownOperator,
		Input:    "^",
	}
	if !errors.Is(err, calcerrors.ErrUnknownOperator) {
		t.Error("errors.Is should return true for CalcError wrapping ErrUnknownOperator")
	}
}

// TestCalcError_ErrorsIs_ErrInvalidNumber verifies errors.Is works when CalcError wraps ErrInvalidNumber.
func TestCalcError_ErrorsIs_ErrInvalidNumber(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrInvalidNumber,
		Input:    "abc",
	}
	if !errors.Is(err, calcerrors.ErrInvalidNumber) {
		t.Error("errors.Is should return true for CalcError wrapping ErrInvalidNumber")
	}
}

// TestCalcError_ErrorsIs_ErrDivisionByZero verifies errors.Is works when CalcError wraps ErrDivisionByZero.
func TestCalcError_ErrorsIs_ErrDivisionByZero(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrDivisionByZero,
		Input:    "0",
	}
	if !errors.Is(err, calcerrors.ErrDivisionByZero) {
		t.Error("errors.Is should return true for CalcError wrapping ErrDivisionByZero")
	}
}

// TestCalcError_ErrorsIs_ErrInvalidArgCount verifies errors.Is works when CalcError wraps ErrInvalidArgCount.
func TestCalcError_ErrorsIs_ErrInvalidArgCount(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrInvalidArgCount,
		Input:    "1 2 3",
	}
	if !errors.Is(err, calcerrors.ErrInvalidArgCount) {
		t.Error("errors.Is should return true for CalcError wrapping ErrInvalidArgCount")
	}
}

// TestCalcError_ErrorsIs_DoesNotMatchOtherSentinels verifies CalcError only matches its own sentinel.
func TestCalcError_ErrorsIs_DoesNotMatchOtherSentinels(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrUnknownOperator,
		Input:    "^",
	}
	if errors.Is(err, calcerrors.ErrInvalidNumber) {
		t.Error("CalcError wrapping ErrUnknownOperator should not match ErrInvalidNumber")
	}
	if errors.Is(err, calcerrors.ErrDivisionByZero) {
		t.Error("CalcError wrapping ErrUnknownOperator should not match ErrDivisionByZero")
	}
	if errors.Is(err, calcerrors.ErrInvalidArgCount) {
		t.Error("CalcError wrapping ErrUnknownOperator should not match ErrInvalidArgCount")
	}
}

// TestCalcError_Error_IncludesSentinelMessage verifies Error() contains the sentinel's message.
func TestCalcError_Error_IncludesSentinelMessage(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrUnknownOperator,
		Input:    "^",
	}
	msg := err.Error()
	if !strings.Contains(msg, calcerrors.ErrUnknownOperator.Error()) {
		t.Errorf("Error() %q should contain sentinel message %q", msg, calcerrors.ErrUnknownOperator.Error())
	}
}

// TestCalcError_Error_IncludesInputValue verifies Error() contains the bad input value.
func TestCalcError_Error_IncludesInputValue(t *testing.T) {
	input := "bad_input_xyz"
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrInvalidNumber,
		Input:    input,
	}
	msg := err.Error()
	if !strings.Contains(msg, input) {
		t.Errorf("Error() %q should contain input value %q", msg, input)
	}
}

// TestCalcError_Error_IncludesBothSentinelAndInput verifies Error() includes both sentinel message and input.
func TestCalcError_Error_IncludesBothSentinelAndInput(t *testing.T) {
	cases := []struct {
		sentinel error
		input    string
	}{
		{calcerrors.ErrUnknownOperator, "^"},
		{calcerrors.ErrInvalidNumber, "not-a-number"},
		{calcerrors.ErrDivisionByZero, "0"},
		{calcerrors.ErrInvalidArgCount, "1 2 3 4"},
	}
	for _, tc := range cases {
		err := &calcerrors.CalcError{
			Sentinel: tc.sentinel,
			Input:    tc.input,
		}
		msg := err.Error()
		if !strings.Contains(msg, tc.sentinel.Error()) {
			t.Errorf("Error() %q should contain sentinel message %q", msg, tc.sentinel.Error())
		}
		if !strings.Contains(msg, tc.input) {
			t.Errorf("Error() %q should contain input %q", msg, tc.input)
		}
	}
}

// TestCalcError_Unwrap_ReturnsSentinel verifies Unwrap() returns the sentinel error.
func TestCalcError_Unwrap_ReturnsSentinel(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrDivisionByZero,
		Input:    "0",
	}
	unwrapped := errors.Unwrap(err)
	if unwrapped != calcerrors.ErrDivisionByZero {
		t.Errorf("Unwrap() = %v, want %v", unwrapped, calcerrors.ErrDivisionByZero)
	}
}

// TestCalcError_ImplementsErrorInterface verifies *CalcError satisfies the error interface.
func TestCalcError_ImplementsErrorInterface(t *testing.T) {
	var _ error = &calcerrors.CalcError{
		Sentinel: calcerrors.ErrUnknownOperator,
		Input:    "test",
	}
}

// TestCalcError_EmptyInput verifies CalcError works with an empty input string.
func TestCalcError_EmptyInput(t *testing.T) {
	err := &calcerrors.CalcError{
		Sentinel: calcerrors.ErrInvalidNumber,
		Input:    "",
	}
	if !errors.Is(err, calcerrors.ErrInvalidNumber) {
		t.Error("errors.Is should return true even with empty input")
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() should not return empty string")
	}
}
