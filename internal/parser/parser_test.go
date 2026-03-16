//go:build !integration

package parser_test

import (
	"testing"

	"github.com/example/calc-app/internal/calc"
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
