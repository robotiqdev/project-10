package calc_test

import (
	"errors"
	"testing"

	"calculator/internal/calc"
	errs "calculator/internal/errors"
)

// TestDivide_ByZeroInteger verifies that Divide(1, 0) returns an error
// wrapping ErrDivisionByZero.
func TestDivide_ByZeroInteger(t *testing.T) {
	_, err := calc.Divide(1, 0)
	if err == nil {
		t.Fatal("Divide(1, 0) should return an error, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("errors.Is(err, ErrDivisionByZero) = false, want true; err = %v", err)
	}
}

// TestDivide_ZeroByZero verifies that Divide(0, 0) returns an error
// wrapping ErrDivisionByZero.
func TestDivide_ZeroByZero(t *testing.T) {
	_, err := calc.Divide(0, 0)
	if err == nil {
		t.Fatal("Divide(0, 0) should return an error, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("errors.Is(err, ErrDivisionByZero) = false, want true; err = %v", err)
	}
}

// TestDivide_ByZeroFloat verifies that Divide(5, 0.0) returns an error
// wrapping ErrDivisionByZero.
func TestDivide_ByZeroFloat(t *testing.T) {
	_, err := calc.Divide(5, 0.0)
	if err == nil {
		t.Fatal("Divide(5, 0.0) should return an error, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("errors.Is(err, ErrDivisionByZero) = false, want true; err = %v", err)
	}
}

// TestDivide_ByVerySmallNonZero verifies that Divide(5, 1e-300) does NOT return
// an error — a very small but non-zero denominator is valid.
func TestDivide_ByVerySmallNonZero(t *testing.T) {
	result, err := calc.Divide(5, 1e-300)
	if err != nil {
		t.Fatalf("Divide(5, 1e-300) should not return an error, got: %v", err)
	}
	if result == 0 {
		t.Error("Divide(5, 1e-300) result should be non-zero")
	}
}

// TestDivide_NormalDivision verifies that Divide returns the correct quotient
// for a normal (non-zero denominator) operation.
func TestDivide_NormalDivision(t *testing.T) {
	cases := []struct {
		a, b, want float64
	}{
		{10, 2, 5},
		{9, 3, 3},
		{1, 4, 0.25},
		{-6, 2, -3},
		{6, -2, -3},
		{0, 5, 0},
	}
	for _, tc := range cases {
		result, err := calc.Divide(tc.a, tc.b)
		if err != nil {
			t.Errorf("Divide(%g, %g) unexpected error: %v", tc.a, tc.b, err)
			continue
		}
		if result != tc.want {
			t.Errorf("Divide(%g, %g) = %g, want %g", tc.a, tc.b, result, tc.want)
		}
	}
}

// TestDivide_ByZero_ErrorIsCalcError verifies that the returned error is a
// *errs.CalcError so callers can inspect the offending input.
func TestDivide_ByZero_ErrorIsCalcError(t *testing.T) {
	_, err := calc.Divide(1, 0)
	if err == nil {
		t.Fatal("Divide(1, 0) should return an error, got nil")
	}
	var calcErr *errs.CalcError
	if !errors.As(err, &calcErr) {
		t.Errorf("error should be *errs.CalcError, got %T: %v", err, err)
	}
}

// TestDivide_ByZero_ErrorDoesNotWrapOtherSentinels verifies that the error
// returned for division by zero does NOT match other sentinel errors.
func TestDivide_ByZero_ErrorDoesNotWrapOtherSentinels(t *testing.T) {
	_, err := calc.Divide(1, 0)
	if err == nil {
		t.Fatal("Divide(1, 0) should return an error, got nil")
	}
	if errors.Is(err, errs.ErrInvalidNumber) {
		t.Error("division-by-zero error should not match ErrInvalidNumber")
	}
	if errors.Is(err, errs.ErrUnknownOperator) {
		t.Error("division-by-zero error should not match ErrUnknownOperator")
	}
	if errors.Is(err, errs.ErrInvalidArgCount) {
		t.Error("division-by-zero error should not match ErrInvalidArgCount")
	}
}
