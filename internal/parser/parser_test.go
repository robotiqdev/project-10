package parser

import (
	"errors"
	"testing"

	errs "github.com/example/calc-app/internal/errors"
)

// TestParse_ValidIntegerArgs verifies that valid integer args return the correct *Input and nil error.
func TestParse_ValidIntegerArgs(t *testing.T) {
	args := []string{"3", "+", "4"}
	got, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse(%v) returned error %v; want nil", args, err)
	}
	if got == nil {
		t.Fatal("Parse returned nil *Input; want non-nil")
	}
	if got.Left != 3 {
		t.Errorf("got.Left = %v; want 3", got.Left)
	}
	if got.Right != 4 {
		t.Errorf("got.Right = %v; want 4", got.Right)
	}
	if got.Operator != "+" {
		t.Errorf("got.Operator = %q; want %q", got.Operator, "+")
	}
}

// TestParse_AllOperators verifies that all supported operators are parsed correctly.
func TestParse_AllOperators(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantLeft float64
		wantRight float64
		wantOp   string
	}{
		{"addition", []string{"3", "+", "4"}, 3, 4, "+"},
		{"subtraction", []string{"10", "-", "5"}, 10, 5, "-"},
		{"multiplication", []string{"6", "*", "7"}, 6, 7, "*"},
		{"division", []string{"8", "/", "2"}, 8, 2, "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err != nil {
				t.Fatalf("Parse(%v) returned error %v; want nil", tt.args, err)
			}
			if got == nil {
				t.Fatalf("Parse(%v) returned nil *Input; want non-nil", tt.args)
			}
			if got.Left != tt.wantLeft {
				t.Errorf("got.Left = %v; want %v", got.Left, tt.wantLeft)
			}
			if got.Right != tt.wantRight {
				t.Errorf("got.Right = %v; want %v", got.Right, tt.wantRight)
			}
			if got.Operator != tt.wantOp {
				t.Errorf("got.Operator = %q; want %q", got.Operator, tt.wantOp)
			}
		})
	}
}

// TestParse_DecimalInputs verifies that decimal (float) args are parsed correctly.
func TestParse_DecimalInputs(t *testing.T) {
	args := []string{"1.5", "*", "2.5"}
	got, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse(%v) returned error %v; want nil", args, err)
	}
	if got == nil {
		t.Fatalf("Parse(%v) returned nil *Input; want non-nil", args)
	}
	if got.Left != 1.5 {
		t.Errorf("got.Left = %v; want 1.5", got.Left)
	}
	if got.Right != 2.5 {
		t.Errorf("got.Right = %v; want 2.5", got.Right)
	}
	if got.Operator != "*" {
		t.Errorf("got.Operator = %q; want %q", got.Operator, "*")
	}
}

// TestParse_NegativeLeftNumber verifies that a negative left number is parsed correctly.
func TestParse_NegativeLeftNumber(t *testing.T) {
	args := []string{"-3", "+", "4"}
	got, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse(%v) returned error %v; want nil", args, err)
	}
	if got == nil {
		t.Fatalf("Parse(%v) returned nil *Input; want non-nil", args)
	}
	if got.Left != -3 {
		t.Errorf("got.Left = %v; want -3", got.Left)
	}
	if got.Right != 4 {
		t.Errorf("got.Right = %v; want 4", got.Right)
	}
	if got.Operator != "+" {
		t.Errorf("got.Operator = %q; want %q", got.Operator, "+")
	}
}

// TestParse_NegativeRightNumber verifies that a negative right number is parsed correctly.
func TestParse_NegativeRightNumber(t *testing.T) {
	args := []string{"5", "-", "-2"}
	got, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse(%v) returned error %v; want nil", args, err)
	}
	if got == nil {
		t.Fatalf("Parse(%v) returned nil *Input; want non-nil", args)
	}
	if got.Left != 5 {
		t.Errorf("got.Left = %v; want 5", got.Left)
	}
	if got.Right != -2 {
		t.Errorf("got.Right = %v; want -2", got.Right)
	}
}

// TestParse_WrongArgCount_TooFew verifies that too few args return an ErrInvalidArgCount error.
func TestParse_WrongArgCount_TooFew(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"empty", []string{}},
		{"one arg", []string{"3"}},
		{"two args", []string{"3", "+"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err == nil {
				t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
			}
			if got != nil {
				t.Errorf("Parse(%v) returned non-nil *Input %+v; want nil", tt.args, got)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v, got err: %v", tt.args, err)
			}
		})
	}
}

// TestParse_WrongArgCount_TooMany verifies that too many args return an ErrInvalidArgCount error.
func TestParse_WrongArgCount_TooMany(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"four args", []string{"3", "+", "4", "extra"}},
		{"five args", []string{"3", "+", "4", "extra1", "extra2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err == nil {
				t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
			}
			if got != nil {
				t.Errorf("Parse(%v) returned non-nil *Input %+v; want nil", tt.args, got)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v, got err: %v", tt.args, err)
			}
		})
	}
}

// TestParse_UnknownOperator verifies that an unknown operator returns an error.
func TestParse_UnknownOperator(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"percent operator", []string{"3", "%", "4"}},
		{"caret operator", []string{"3", "^", "4"}},
		{"word operator", []string{"3", "plus", "4"}},
		{"empty operator", []string{"3", "", "4"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err == nil {
				t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
			}
			if got != nil {
				t.Errorf("Parse(%v) returned non-nil *Input %+v; want nil", tt.args, got)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v, got err: %v", tt.args, err)
			}
		})
	}
}

// TestParse_InvalidLeftNumber verifies that an invalid left number returns an error.
func TestParse_InvalidLeftNumber(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"left is word", []string{"abc", "+", "4"}},
		{"left is empty", []string{"", "+", "4"}},
		{"left has mixed chars", []string{"3x", "+", "4"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err == nil {
				t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
			}
			if got != nil {
				t.Errorf("Parse(%v) returned non-nil *Input %+v; want nil", tt.args, got)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v, got err: %v", tt.args, err)
			}
		})
	}
}

// TestParse_InvalidRightNumber verifies that an invalid right number returns an error.
func TestParse_InvalidRightNumber(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{"right is word", []string{"3", "+", "abc"}},
		{"right is empty", []string{"3", "+", ""}},
		{"right has mixed chars", []string{"3", "+", "4y"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if err == nil {
				t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
			}
			if got != nil {
				t.Errorf("Parse(%v) returned non-nil *Input %+v; want nil", tt.args, got)
			}
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v, got err: %v", tt.args, err)
			}
		})
	}
}

// TestParse_TableDriven is a comprehensive table-driven test for Parse().
func TestParse_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantLeft  float64
		wantRight float64
		wantOp    string
		wantErr   bool
		errSentinel error
	}{
		{
			name:      "integer addition",
			args:      []string{"3", "+", "4"},
			wantLeft:  3,
			wantRight: 4,
			wantOp:    "+",
			wantErr:   false,
		},
		{
			name:      "decimal multiplication",
			args:      []string{"1.5", "*", "2.5"},
			wantLeft:  1.5,
			wantRight: 2.5,
			wantOp:    "*",
			wantErr:   false,
		},
		{
			name:      "negative left number",
			args:      []string{"-3", "+", "4"},
			wantLeft:  -3,
			wantRight: 4,
			wantOp:    "+",
			wantErr:   false,
		},
		{
			name:        "wrong arg count - too few",
			args:        []string{"3", "+"},
			wantErr:     true,
			errSentinel: errs.ErrInvalidArgCount,
		},
		{
			name:        "wrong arg count - too many",
			args:        []string{"3", "+", "4", "extra"},
			wantErr:     true,
			errSentinel: errs.ErrInvalidArgCount,
		},
		{
			name:        "unknown operator",
			args:        []string{"3", "%", "4"},
			wantErr:     true,
			errSentinel: errs.ErrInvalidArgCount,
		},
		{
			name:        "invalid left number",
			args:        []string{"abc", "+", "4"},
			wantErr:     true,
			errSentinel: errs.ErrInvalidArgCount,
		},
		{
			name:        "invalid right number",
			args:        []string{"3", "+", "xyz"},
			wantErr:     true,
			errSentinel: errs.ErrInvalidArgCount,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("Parse(%v) returned nil error; want error", tt.args)
				}
				if got != nil {
					t.Errorf("Parse(%v) returned non-nil *Input; want nil on error", tt.args)
				}
				if tt.errSentinel != nil && !errors.Is(err, tt.errSentinel) {
					t.Errorf("errors.Is(err, %v) = false; want true for args %v, err: %v", tt.errSentinel, tt.args, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Parse(%v) returned error %v; want nil", tt.args, err)
				}
				if got == nil {
					t.Fatalf("Parse(%v) returned nil *Input; want non-nil", tt.args)
				}
				if got.Left != tt.wantLeft {
					t.Errorf("got.Left = %v; want %v", got.Left, tt.wantLeft)
				}
				if got.Right != tt.wantRight {
					t.Errorf("got.Right = %v; want %v", got.Right, tt.wantRight)
				}
				if got.Operator != tt.wantOp {
					t.Errorf("got.Operator = %q; want %q", got.Operator, tt.wantOp)
				}
			}
		})
	}
}

// TestParse_ReturnsPointerNotValue verifies that Parse returns a *Input pointer (not a value).
func TestParse_ReturnsPointerNotValue(t *testing.T) {
	args := []string{"3", "+", "4"}
	got, err := Parse(args)
	if err != nil {
		t.Fatalf("Parse(%v) returned error %v; want nil", args, err)
	}
	if got == nil {
		t.Fatal("Parse returned nil pointer; want non-nil *Input")
	}
}
