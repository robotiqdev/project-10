//go:build !integration

package main_test

import (
	"errors"
	"testing"

	"github.com/example/calc-app/internal/calc"
	errs "github.com/example/calc-app/internal/errors"
	"github.com/example/calc-app/internal/parser"
)

// TestPipeline_Addition verifies the full pipeline: Parse "3 + 4" then Calculate produces "7".
func TestPipeline_Addition(t *testing.T) {
	input, err := parser.Parse([]string{"3", "+", "4"})
	if err != nil {
		t.Fatalf("Parse(3 + 4) failed: %v", err)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err != nil {
		t.Fatalf("Calculate(3 + 4) returned unexpected error: %v", err)
	}
	if result != "7" {
		t.Errorf("pipeline(3 + 4) = %q, want %q", result, "7")
	}
}

// TestPipeline_DivisionByZero verifies the full pipeline: Parse "10 / 0" then
// Calculate returns a division-by-zero error.
func TestPipeline_DivisionByZero(t *testing.T) {
	input, err := parser.Parse([]string{"10", "/", "0"})
	if err != nil {
		t.Fatalf("Parse(10 / 0) failed: %v", err)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err == nil {
		t.Fatal("pipeline(10 / 0) expected a division-by-zero error, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("pipeline(10 / 0) error = %v, want to wrap ErrDivisionByZero", err)
	}
	if result != "" {
		t.Errorf("pipeline(10 / 0) result = %q, want empty string on error", result)
	}
}

// TestPipeline_Multiplication verifies the full pipeline: Parse "1.5 * 2" then
// Calculate produces "3".
func TestPipeline_Multiplication(t *testing.T) {
	input, err := parser.Parse([]string{"1.5", "*", "2"})
	if err != nil {
		t.Fatalf("Parse(1.5 * 2) failed: %v", err)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err != nil {
		t.Fatalf("Calculate(1.5 * 2) returned unexpected error: %v", err)
	}
	if result != "3" {
		t.Errorf("pipeline(1.5 * 2) = %q, want %q", result, "3")
	}
}

// TestPipeline_NegativeNumbers verifies the full pipeline: Parse "-3 + -2" then
// Calculate produces "-5".
func TestPipeline_NegativeNumbers(t *testing.T) {
	input, err := parser.Parse([]string{"-3", "+", "-2"})
	if err != nil {
		t.Fatalf("Parse(-3 + -2) failed: %v", err)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err != nil {
		t.Fatalf("Calculate(-3 + -2) returned unexpected error: %v", err)
	}
	if result != "-5" {
		t.Errorf("pipeline(-3 + -2) = %q, want %q", result, "-5")
	}
}

// TestPipeline_DivisionWhole verifies the full pipeline: Parse "9 / 3" then
// Calculate produces "3".
func TestPipeline_DivisionWhole(t *testing.T) {
	input, err := parser.Parse([]string{"9", "/", "3"})
	if err != nil {
		t.Fatalf("Parse(9 / 3) failed: %v", err)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err != nil {
		t.Fatalf("Calculate(9 / 3) returned unexpected error: %v", err)
	}
	if result != "3" {
		t.Errorf("pipeline(9 / 3) = %q, want %q", result, "3")
	}
}

// TestPipeline_BadArgCount verifies that Parse propagates an error when given
// the wrong number of arguments, preventing Calculate from being reached.
func TestPipeline_BadArgCount(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"no args", []string{}},
		{"one arg", []string{"3"}},
		{"two args", []string{"3", "+"}},
		{"four args", []string{"3", "+", "4", "5"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.Parse(tc.args)
			if err == nil {
				t.Errorf("Parse(%v) expected error for bad arg count, got nil", tc.args)
			}
		})
	}
}

// TestPipeline_BadOperator verifies end-to-end error propagation when an
// unrecognized operator is provided. The error should surface either from
// Parse (if it validates operators) or from Engine.Calculate.
func TestPipeline_BadOperator(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"percent", []string{"3", "%", "4"}},
		{"caret", []string{"3", "^", "4"}},
		{"word mod", []string{"3", "mod", "4"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input, parseErr := parser.Parse(tc.args)
			if parseErr != nil {
				// Parser rejected the unknown operator — correct behavior.
				return
			}

			// If Parse did not reject the operator, Calculate must return an error.
			engine := calc.NewEngine()
			_, calcErr := engine.Calculate(calc.Operation{
				Left:     input.Left,
				Operator: input.Operator,
				Right:    input.Right,
			})
			if calcErr == nil {
				t.Errorf("pipeline(%v) expected an error for unknown operator, got nil", tc.args)
			}
			if !errors.Is(calcErr, errs.ErrUnknownOperator) {
				t.Errorf("pipeline(%v) error = %v, want to wrap ErrUnknownOperator", tc.args, calcErr)
			}
		})
	}
}

// TestPipeline_InvalidOperand verifies that Parse returns an error when an
// operand cannot be parsed as a number.
func TestPipeline_InvalidOperand(t *testing.T) {
	cases := []struct {
		name string
		args []string
	}{
		{"non-numeric left", []string{"abc", "+", "4"}},
		{"non-numeric right", []string{"3", "+", "xyz"}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := parser.Parse(tc.args)
			if err == nil {
				t.Errorf("Parse(%v) expected error for invalid operand, got nil", tc.args)
			}
		})
	}
}
