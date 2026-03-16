package validation_test

import (
	"errors"
	"testing"

	"calculator/validation"
)

// TestCalcError_ImplementsErrorInterface verifies that CalcError satisfies the error interface.
func TestCalcError_ImplementsErrorInterface(t *testing.T) {
	var _ error = validation.CalcError("test message")
}

// TestCalcError_Error_ReturnsWrappedString verifies Error() returns the string content.
func TestCalcError_Error_ReturnsWrappedString(t *testing.T) {
	msg := "division by zero"
	e := validation.CalcError(msg)
	if e.Error() != msg {
		t.Errorf("CalcError(%q).Error() = %q, want %q", msg, e.Error(), msg)
	}
}

// TestCalcError_Error_EmptyString verifies Error() on empty CalcError returns empty string.
func TestCalcError_Error_EmptyString(t *testing.T) {
	e := validation.CalcError("")
	if e.Error() != "" {
		t.Errorf("CalcError(%q).Error() = %q, want %q", "", e.Error(), "")
	}
}

// TestSentinelErrors_NonEmpty verifies all four sentinel errors have non-empty Error() messages.
func TestSentinelErrors_NonEmpty(t *testing.T) {
	sentinels := []struct {
		name string
		err  error
	}{
		{"ErrInvalidOperator", validation.ErrInvalidOperator},
		{"ErrInvalidNumber", validation.ErrInvalidNumber},
		{"ErrDivisionByZero", validation.ErrDivisionByZero},
		{"ErrInvalidArgCount", validation.ErrInvalidArgCount},
	}

	for _, tc := range sentinels {
		if tc.err.Error() == "" {
			t.Errorf("%s.Error() should not be empty string", tc.name)
		}
	}
}

// TestSentinelErrors_Distinct verifies that all four sentinel errors are distinct
// (errors.Is does not confuse one for another).
func TestSentinelErrors_Distinct(t *testing.T) {
	type namedErr struct {
		name string
		err  error
	}
	sentinels := []namedErr{
		{"ErrInvalidOperator", validation.ErrInvalidOperator},
		{"ErrInvalidNumber", validation.ErrInvalidNumber},
		{"ErrDivisionByZero", validation.ErrDivisionByZero},
		{"ErrInvalidArgCount", validation.ErrInvalidArgCount},
	}

	for i := 0; i < len(sentinels); i++ {
		for j := i + 1; j < len(sentinels); j++ {
			if errors.Is(sentinels[i].err, sentinels[j].err) {
				t.Errorf("%s and %s should be distinct errors, but errors.Is reports them as equal",
					sentinels[i].name, sentinels[j].name)
			}
		}
	}
}

// TestSentinelErrors_ErrorsIs_MatchesItself verifies each sentinel matches itself via errors.Is.
func TestSentinelErrors_ErrorsIs_MatchesItself(t *testing.T) {
	sentinels := []struct {
		name string
		err  error
	}{
		{"ErrInvalidOperator", validation.ErrInvalidOperator},
		{"ErrInvalidNumber", validation.ErrInvalidNumber},
		{"ErrDivisionByZero", validation.ErrDivisionByZero},
		{"ErrInvalidArgCount", validation.ErrInvalidArgCount},
	}

	for _, tc := range sentinels {
		if !errors.Is(tc.err, tc.err) {
			t.Errorf("errors.Is(%s, %s) should return true", tc.name, tc.name)
		}
	}
}

// TestErrInvalidOperator_Message verifies ErrInvalidOperator has a meaningful message.
func TestErrInvalidOperator_Message(t *testing.T) {
	msg := validation.ErrInvalidOperator.Error()
	if msg == "" {
		t.Error("ErrInvalidOperator.Error() should not be empty")
	}
}

// TestErrInvalidNumber_Message verifies ErrInvalidNumber has a meaningful message.
func TestErrInvalidNumber_Message(t *testing.T) {
	msg := validation.ErrInvalidNumber.Error()
	if msg == "" {
		t.Error("ErrInvalidNumber.Error() should not be empty")
	}
}

// TestErrDivisionByZero_Message verifies ErrDivisionByZero has a meaningful message.
func TestErrDivisionByZero_Message(t *testing.T) {
	msg := validation.ErrDivisionByZero.Error()
	if msg == "" {
		t.Error("ErrDivisionByZero.Error() should not be empty")
	}
}

// TestErrInvalidArgCount_Message verifies ErrInvalidArgCount has a meaningful message.
func TestErrInvalidArgCount_Message(t *testing.T) {
	msg := validation.ErrInvalidArgCount.Error()
	if msg == "" {
		t.Error("ErrInvalidArgCount.Error() should not be empty")
	}
}
