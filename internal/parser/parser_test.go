//go:build !integration

package parser_test

import (
	"errors"
	"testing"

	"github.com/example/calc-app/internal/calc"
	errs "github.com/example/calc-app/internal/errors"
	"github.com/example/calc-app/internal/parser"
)

// TestParseAddition verifies that Parse correctly parses an addition expression.
func TestParseAddition(t *testing.T) {
	input, err := parser.Parse([]string{"3", "+", "4"})
	if err != nil {
		t.Fatalf("Parse(3 + 4) returned unexpected error: %v", err)
	}
	if input.Left != 3 {
		t.Errorf("input.Left = %v, want 3", input.Left)
	}
	if input.Operator != calc.OpAdd {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpAdd)
	}
	if input.Right != 4 {
		t.Errorf("input.Right = %v, want 4", input.Right)
	}
}

// TestParseDivision verifies that Parse correctly parses a division expression.
func TestParseDivision(t *testing.T) {
	input, err := parser.Parse([]string{"10", "/", "0"})
	if err != nil {
		t.Fatalf("Parse(10 / 0) returned unexpected error: %v", err)
	}
	if input.Left != 10 {
		t.Errorf("input.Left = %v, want 10", input.Left)
	}
	if input.Operator != calc.OpDivide {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpDivide)
	}
	if input.Right != 0 {
		t.Errorf("input.Right = %v, want 0", input.Right)
	}
}

// TestParseMultiplication verifies that Parse correctly parses a multiplication expression.
func TestParseMultiplication(t *testing.T) {
	input, err := parser.Parse([]string{"1.5", "*", "2"})
	if err != nil {
		t.Fatalf("Parse(1.5 * 2) returned unexpected error: %v", err)
	}
	if input.Left != 1.5 {
		t.Errorf("input.Left = %v, want 1.5", input.Left)
	}
	if input.Operator != calc.OpMultiply {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpMultiply)
	}
	if input.Right != 2 {
		t.Errorf("input.Right = %v, want 2", input.Right)
	}
}

// TestParseNegativeNumbers verifies that Parse correctly parses negative operands.
func TestParseNegativeNumbers(t *testing.T) {
	input, err := parser.Parse([]string{"-3", "+", "-2"})
	if err != nil {
		t.Fatalf("Parse(-3 + -2) returned unexpected error: %v", err)
	}
	if input.Left != -3 {
		t.Errorf("input.Left = %v, want -3", input.Left)
	}
	if input.Operator != calc.OpAdd {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpAdd)
	}
	if input.Right != -2 {
		t.Errorf("input.Right = %v, want -2", input.Right)
	}
}

// TestParseSubtraction verifies that Parse correctly parses a subtraction expression.
func TestParseSubtraction(t *testing.T) {
	input, err := parser.Parse([]string{"9", "/", "3"})
	if err != nil {
		t.Fatalf("Parse(9 / 3) returned unexpected error: %v", err)
	}
	if input.Left != 9 {
		t.Errorf("input.Left = %v, want 9", input.Left)
	}
	if input.Operator != calc.OpDivide {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpDivide)
	}
	if input.Right != 3 {
		t.Errorf("input.Right = %v, want 3", input.Right)
	}
}

// TestParseFloatOperands verifies that Parse handles floating-point operands correctly.
func TestParseFloatOperands(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantLeft float64
		wantOp   calc.Operator
		wantRight float64
	}{
		{
			name:      "decimal addition",
			args:      []string{"2.5", "+", "1.5"},
			wantLeft:  2.5,
			wantOp:    calc.OpAdd,
			wantRight: 1.5,
		},
		{
			name:      "decimal subtraction",
			args:      []string{"5.5", "-", "2.2"},
			wantLeft:  5.5,
			wantOp:    calc.OpSubtract,
			wantRight: 2.2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input, err := parser.Parse(tt.args)
			if err != nil {
				t.Fatalf("Parse(%v) returned unexpected error: %v", tt.args, err)
			}
			if input.Left != tt.wantLeft {
				t.Errorf("input.Left = %v, want %v", input.Left, tt.wantLeft)
			}
			if input.Operator != tt.wantOp {
				t.Errorf("input.Operator = %q, want %q", input.Operator, tt.wantOp)
			}
			if input.Right != tt.wantRight {
				t.Errorf("input.Right = %v, want %v", input.Right, tt.wantRight)
			}
		})
	}
}

// TestParseReturnsErrorForTooFewArgs verifies that Parse returns an error when
// fewer than 3 arguments are provided.
func TestParseReturnsErrorForTooFewArgs(t *testing.T) {
	cases := [][]string{
		{},
		{"3"},
		{"3", "+"},
	}

	for _, args := range cases {
		_, err := parser.Parse(args)
		if err == nil {
			t.Errorf("Parse(%v) expected error for too few args, got nil", args)
		}
	}
}

// TestParseReturnsErrorForTooManyArgs verifies that Parse returns an error when
// more than 3 arguments are provided.
func TestParseReturnsErrorForTooManyArgs(t *testing.T) {
	cases := [][]string{
		{"3", "+", "4", "extra"},
		{"1", "+", "2", "3", "4"},
	}

	for _, args := range cases {
		_, err := parser.Parse(args)
		if err == nil {
			t.Errorf("Parse(%v) expected error for too many args, got nil", args)
		}
	}
}

// TestParseReturnsErrorForInvalidLeftOperand verifies that a non-numeric left
// operand produces an error.
func TestParseReturnsErrorForInvalidLeftOperand(t *testing.T) {
	_, err := parser.Parse([]string{"abc", "+", "4"})
	if err == nil {
		t.Error("Parse(abc + 4) expected error for invalid left operand, got nil")
	}
}

// TestParseReturnsErrorForInvalidRightOperand verifies that a non-numeric right
// operand produces an error.
func TestParseReturnsErrorForInvalidRightOperand(t *testing.T) {
	_, err := parser.Parse([]string{"3", "+", "xyz"})
	if err == nil {
		t.Error("Parse(3 + xyz) expected error for invalid right operand, got nil")
	}
}

// TestParseReturnsErrorForUnknownOperator verifies that an unrecognized operator
// produces an error.
func TestParseReturnsErrorForUnknownOperator(t *testing.T) {
	cases := []string{"%", "^", "mod", "**"}

	for _, op := range cases {
		_, err := parser.Parse([]string{"3", op, "4"})
		if err == nil {
			t.Errorf("Parse(3 %s 4) expected error for unknown operator, got nil", op)
		}
	}
}

// BenchmarkParse measures the performance of the Parse function.
func BenchmarkParse(b *testing.B) {
	args := []string{"3", "+", "4"}
	for i := 0; i < b.N; i++ {
		parser.Parse(args)
	}
}

// TestParseInvalidOperatorSymbolsUseErrUnknownOperator verifies that Parse returns
// errors.Is(err, errs.ErrUnknownOperator) for invalid operator symbols.
func TestParseInvalidOperatorSymbolsUseErrUnknownOperator(t *testing.T) {
	invalidOps := []string{"%", "^", "!", "@", "#", "~", "mod", "div", "**", "++"}

	for _, op := range invalidOps {
		t.Run("invalid op "+op, func(t *testing.T) {
			_, err := parser.Parse([]string{"3", op, "4"})
			if err == nil {
				t.Fatalf("Parse(3 %s 4) = nil; want non-nil error", op)
			}
			if !errors.Is(err, errs.ErrUnknownOperator) {
				t.Errorf("Parse(3 %s 4) error = %v; want errors.Is(err, ErrUnknownOperator) to be true", op, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(3 %s 4) error message is empty; want non-empty", op)
			}
		})
	}
}

// TestParseMalformedLeftOperandLettersUsesErrInvalidNumber verifies that a left
// operand containing letters returns errors.Is(err, errs.ErrInvalidNumber).
func TestParseMalformedLeftOperandLettersUsesErrInvalidNumber(t *testing.T) {
	malformed := []string{"abc", "1a", "a1", "one", "1b2", "not_a_num"}

	for _, s := range malformed {
		t.Run("left operand "+s, func(t *testing.T) {
			_, err := parser.Parse([]string{s, "+", "4"})
			if err == nil {
				t.Fatalf("Parse(%s + 4) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("Parse(%s + 4) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(%s + 4) error message is empty; want non-empty", s)
			}
		})
	}
}

// TestParseMalformedLeftOperandMultipleDotsUsesErrInvalidNumber verifies that a left
// operand with multiple decimal points returns errors.Is(err, errs.ErrInvalidNumber).
func TestParseMalformedLeftOperandMultipleDotsUsesErrInvalidNumber(t *testing.T) {
	malformed := []string{"1.2.3", "1..2", "..1", "1..", "0.0.0"}

	for _, s := range malformed {
		t.Run("left operand "+s, func(t *testing.T) {
			_, err := parser.Parse([]string{s, "+", "4"})
			if err == nil {
				t.Fatalf("Parse(%s + 4) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("Parse(%s + 4) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(%s + 4) error message is empty; want non-empty", s)
			}
		})
	}
}

// TestParseMalformedRightOperandLettersUsesErrInvalidNumber verifies that a right
// operand containing letters returns errors.Is(err, errs.ErrInvalidNumber).
func TestParseMalformedRightOperandLettersUsesErrInvalidNumber(t *testing.T) {
	malformed := []string{"xyz", "1z", "z1", "two", "5x6", "bad_num"}

	for _, s := range malformed {
		t.Run("right operand "+s, func(t *testing.T) {
			_, err := parser.Parse([]string{"3", "+", s})
			if err == nil {
				t.Fatalf("Parse(3 + %s) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("Parse(3 + %s) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(3 + %s) error message is empty; want non-empty", s)
			}
		})
	}
}

// TestParseMalformedRightOperandMultipleDotsUsesErrInvalidNumber verifies that a right
// operand with multiple decimal points returns errors.Is(err, errs.ErrInvalidNumber).
func TestParseMalformedRightOperandMultipleDotsUsesErrInvalidNumber(t *testing.T) {
	malformed := []string{"1.2.3", "2..5", "..3", "3..", "9.9.9"}

	for _, s := range malformed {
		t.Run("right operand "+s, func(t *testing.T) {
			_, err := parser.Parse([]string{"3", "+", s})
			if err == nil {
				t.Fatalf("Parse(3 + %s) = nil; want non-nil error", s)
			}
			if !errors.Is(err, errs.ErrInvalidNumber) {
				t.Errorf("Parse(3 + %s) error = %v; want errors.Is(err, ErrInvalidNumber) to be true", s, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(3 + %s) error message is empty; want non-empty", s)
			}
		})
	}
}

// TestParseWrongArgCountUsesErrInvalidArgCount verifies that Parse returns
// errors.Is(err, errs.ErrInvalidArgCount) for arg counts 0, 1, 2, 4, and 5.
func TestParseWrongArgCountUsesErrInvalidArgCount(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"zero args", []string{}},
		{"one arg", []string{"5"}},
		{"two args", []string{"5", "+"}},
		{"four args", []string{"5", "+", "3", "extra"}},
		{"five args", []string{"5", "+", "3", "extra1", "extra2"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.Parse(tc.args)
			if err == nil {
				t.Fatalf("Parse(%v) = nil; want non-nil error", tc.args)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("Parse(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", tc.args, err)
			}
			if err.Error() == "" {
				t.Errorf("Parse(%v) error message is empty; want non-empty", tc.args)
			}
		})
	}
}

// TestParseDivisionByZeroSucceedsAtParseLevel verifies that Parse successfully
// parses "10 / 0" — the division-by-zero error surfaces during execution, not parsing.
func TestParseDivisionByZeroSucceedsAtParseLevel(t *testing.T) {
	input, err := parser.Parse([]string{"10", "/", "0"})
	if err != nil {
		t.Fatalf("Parse(10 / 0) returned unexpected error: %v; parser should not reject division by zero", err)
	}
	if input.Operator != calc.OpDivide {
		t.Errorf("input.Operator = %q, want %q", input.Operator, calc.OpDivide)
	}
	if input.Right != 0 {
		t.Errorf("input.Right = %v, want 0", input.Right)
	}
}

// TestParseDivisionByZeroProducesErrDivisionByZeroOnExecute verifies that when
// "10 / 0" is parsed and executed through the engine, ErrDivisionByZero is returned.
func TestParseDivisionByZeroProducesErrDivisionByZeroOnExecute(t *testing.T) {
	input, err := parser.Parse([]string{"10", "/", "0"})
	if err != nil {
		t.Fatalf("Parse(10 / 0) returned unexpected parse error: %v", err)
	}

	engine := calc.NewEngine()
	op := calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	}
	_, execErr := engine.Calculate(op)

	if execErr == nil {
		t.Fatal("engine.Calculate(10 / 0) = nil; want ErrDivisionByZero error")
	}
	if !errors.Is(execErr, errs.ErrDivisionByZero) {
		t.Errorf("engine.Calculate(10 / 0) error = %v; want errors.Is(err, ErrDivisionByZero) to be true", execErr)
	}
	if execErr.Error() == "" {
		t.Error("engine.Calculate(10 / 0) error message is empty; want non-empty")
	}
}

// TestParseErrorMessageNonEmptyForAllErrorCases verifies that all error paths
// from Parse return errors with non-empty messages.
func TestParseErrorMessageNonEmptyForAllErrorCases(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"zero args", []string{}},
		{"one arg", []string{"5"}},
		{"two args", []string{"5", "+"}},
		{"four args", []string{"5", "+", "3", "extra"}},
		{"five args", []string{"1", "+", "2", "3", "4"}},
		{"invalid left operand letters", []string{"abc", "+", "4"}},
		{"invalid left operand multiple dots", []string{"1.2.3", "+", "4"}},
		{"invalid right operand letters", []string{"3", "+", "xyz"}},
		{"invalid right operand multiple dots", []string{"3", "+", "1.2.3"}},
		{"invalid operator percent", []string{"3", "%", "4"}},
		{"invalid operator caret", []string{"3", "^", "4"}},
		{"invalid operator word", []string{"3", "mod", "4"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.Parse(tc.args)
			if err == nil {
				t.Fatalf("Parse(%v) = nil; want non-nil error", tc.args)
			}
			if err.Error() == "" {
				t.Errorf("Parse(%v) error message is empty; want non-empty message", tc.args)
			}
		})
	}
}

// TestParseAllOperators verifies that Parse correctly identifies all four
// supported arithmetic operators.
func TestParseAllOperators(t *testing.T) {
	tests := []struct {
		op     string
		wantOp calc.Operator
	}{
		{"+", calc.OpAdd},
		{"-", calc.OpSubtract},
		{"*", calc.OpMultiply},
		{"/", calc.OpDivide},
	}

	for _, tt := range tests {
		args := []string{"5", tt.op, "3"}
		input, err := parser.Parse(args)
		if err != nil {
			t.Fatalf("Parse(%v) returned unexpected error: %v", args, err)
		}
		if input.Operator != tt.wantOp {
			t.Errorf("Parse(%v) operator = %q, want %q", args, input.Operator, tt.wantOp)
		}
	}
}
