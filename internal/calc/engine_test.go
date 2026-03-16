package calc_test

import (
	"errors"
	"strconv"
	"testing"

	"calculator/internal/calc"
	errs "calculator/internal/errors"
	"calculator/internal/parser"
)

// parseResultFloat is a test helper that parses a result string to float64.
func parseResultFloat(t *testing.T, result string) float64 {
	t.Helper()
	v, err := strconv.ParseFloat(result, 64)
	if err != nil {
		t.Fatalf("result %q is not a valid float: %v", result, err)
	}
	return v
}

// TestFullPipeline_Addition verifies the complete pipeline with addition:
// args → Parse → Operation → Engine.Calculate → formatted string.
func TestFullPipeline_Addition(t *testing.T) {
	args := []string{"3", "+", "4"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 7.0 {
		t.Errorf("3 + 4: expected 7, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_Subtraction verifies the complete pipeline with subtraction.
func TestFullPipeline_Subtraction(t *testing.T) {
	args := []string{"10", "-", "3"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 7.0 {
		t.Errorf("10 - 3: expected 7, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_Multiplication verifies the complete pipeline with multiplication.
func TestFullPipeline_Multiplication(t *testing.T) {
	args := []string{"3", "*", "4"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 12.0 {
		t.Errorf("3 * 4: expected 12, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_Division verifies the complete pipeline with division.
func TestFullPipeline_Division(t *testing.T) {
	args := []string{"10", "/", "2"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 5.0 {
		t.Errorf("10 / 2: expected 5, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_DivisionDecimalResult verifies that division producing
// a decimal result is represented correctly in the formatted string output.
func TestFullPipeline_DivisionDecimalResult(t *testing.T) {
	args := []string{"1", "/", "4"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 0.25 {
		t.Errorf("1 / 4: expected 0.25, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_DecimalNumbers verifies the pipeline handles decimal number inputs.
func TestFullPipeline_DecimalNumbers(t *testing.T) {
	args := []string{"1.5", "+", "2.5"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 4.0 {
		t.Errorf("1.5 + 2.5: expected 4, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_NegativeNumbers verifies the pipeline handles negative number inputs.
func TestFullPipeline_NegativeNumbers(t *testing.T) {
	args := []string{"-3", "+", "5"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(op)
	if err != nil {
		t.Fatalf("unexpected calculate error: %v", err)
	}

	got := parseResultFloat(t, result)
	if got != 2.0 {
		t.Errorf("-3 + 5: expected 2, got %v (result=%q)", got, result)
	}
}

// TestFullPipeline_DivisionByZero verifies the pipeline returns an error
// wrapping ErrDivisionByZero when dividing by zero.
func TestFullPipeline_DivisionByZero(t *testing.T) {
	args := []string{"5", "/", "0"}

	input, err := parser.Parse(args)
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	_, err = engine.Calculate(op)
	if err == nil {
		t.Fatal("5 / 0: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrDivisionByZero) {
		t.Errorf("5 / 0: expected error wrapping ErrDivisionByZero, got: %v", err)
	}
}

// TestFullPipeline_InvalidNumber_Left verifies that Parse returns an error
// wrapping ErrInvalidNumber when the left argument is not a valid number.
func TestFullPipeline_InvalidNumber_Left(t *testing.T) {
	args := []string{"abc", "+", "4"}

	_, err := parser.Parse(args)
	if err == nil {
		t.Fatal("Parse with invalid left number: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrInvalidNumber) {
		t.Errorf("Parse with invalid left number: expected error wrapping ErrInvalidNumber, got: %v", err)
	}
}

// TestFullPipeline_InvalidNumber_Right verifies that Parse returns an error
// wrapping ErrInvalidNumber when the right argument is not a valid number.
func TestFullPipeline_InvalidNumber_Right(t *testing.T) {
	args := []string{"3", "+", "xyz"}

	_, err := parser.Parse(args)
	if err == nil {
		t.Fatal("Parse with invalid right number: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrInvalidNumber) {
		t.Errorf("Parse with invalid right number: expected error wrapping ErrInvalidNumber, got: %v", err)
	}
}

// TestFullPipeline_InvalidArgCount_TooFew verifies that Parse returns an error
// wrapping ErrInvalidArgCount when too few arguments are provided.
func TestFullPipeline_InvalidArgCount_TooFew(t *testing.T) {
	args := []string{"3", "+"}

	_, err := parser.Parse(args)
	if err == nil {
		t.Fatal("Parse with too few args: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("Parse with too few args: expected error wrapping ErrInvalidArgCount, got: %v", err)
	}
}

// TestFullPipeline_InvalidArgCount_TooMany verifies that Parse returns an error
// wrapping ErrInvalidArgCount when too many arguments are provided.
func TestFullPipeline_InvalidArgCount_TooMany(t *testing.T) {
	args := []string{"3", "+", "4", "5"}

	_, err := parser.Parse(args)
	if err == nil {
		t.Fatal("Parse with too many args: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("Parse with too many args: expected error wrapping ErrInvalidArgCount, got: %v", err)
	}
}

// TestFullPipeline_InvalidArgCount_Empty verifies that Parse returns an error
// wrapping ErrInvalidArgCount when no arguments are provided.
func TestFullPipeline_InvalidArgCount_Empty(t *testing.T) {
	args := []string{}

	_, err := parser.Parse(args)
	if err == nil {
		t.Fatal("Parse with empty args: expected an error, got nil")
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("Parse with empty args: expected error wrapping ErrInvalidArgCount, got: %v", err)
	}
}

// TestFullPipeline_UnknownOperator verifies that the pipeline returns an error
// wrapping ErrUnknownOperator when an unrecognized operator is used.
// The error may be returned by Parse or by Engine.Calculate.
func TestFullPipeline_UnknownOperator(t *testing.T) {
	args := []string{"3", "%", "4"}

	input, err := parser.Parse(args)
	if err != nil {
		// Parser may validate the operator — ErrUnknownOperator is acceptable here.
		if errors.Is(err, errs.ErrUnknownOperator) {
			return
		}
		t.Fatalf("Parse with unknown operator: unexpected error: %v", err)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	engine := calc.NewEngine()
	_, err = engine.Calculate(op)
	if err == nil {
		t.Fatal("3 % 4: expected an error for unknown operator, got nil")
	}
	if !errors.Is(err, errs.ErrUnknownOperator) {
		t.Errorf("3 %% 4: expected error wrapping ErrUnknownOperator, got: %v", err)
	}
}

// TestDataFlowContract_OperatorCast verifies that casting input.Operator (string)
// to calc.Operator preserves the operator value — the core data-flow contract
// at the boundary between the parser and engine packages.
func TestDataFlowContract_OperatorCast(t *testing.T) {
	operators := []string{"+", "-", "*", "/"}

	for _, opStr := range operators {
		args := []string{"1", opStr, "1"}

		input, err := parser.Parse(args)
		if err != nil {
			t.Errorf("Parse(%v) unexpected error: %v", args, err)
			continue
		}

		// The cast from string to calc.Operator must preserve the value.
		op := calc.Operator(input.Operator)
		if string(op) != opStr {
			t.Errorf("operator cast: expected %q, got %q", opStr, string(op))
		}
	}
}

// TestFullPipeline_MultipleOperations uses table-driven tests to verify
// the complete pipeline across a variety of inputs.
func TestFullPipeline_MultipleOperations(t *testing.T) {
	cases := []struct {
		left     string
		operator string
		right    string
		want     float64
	}{
		{"5", "+", "3", 8},
		{"10", "-", "4", 6},
		{"6", "*", "7", 42},
		{"15", "/", "3", 5},
		{"0", "+", "0", 0},
		{"100", "-", "100", 0},
		{"2.5", "*", "4", 10},
		{"9", "/", "4", 2.25},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.left+" "+tc.operator+" "+tc.right, func(t *testing.T) {
			args := []string{tc.left, tc.operator, tc.right}

			input, err := parser.Parse(args)
			if err != nil {
				t.Fatalf("Parse unexpected error: %v", err)
			}

			op := calc.Operation{
				Left:     input.Left,
				Right:    input.Right,
				Operator: calc.Operator(input.Operator),
			}

			engine := calc.NewEngine()
			result, err := engine.Calculate(op)
			if err != nil {
				t.Fatalf("Calculate unexpected error: %v", err)
			}

			got := parseResultFloat(t, result)
			if got != tc.want {
				t.Errorf("expected %v, got %v (result=%q)", tc.want, got, result)
			}
		})
	}
}
