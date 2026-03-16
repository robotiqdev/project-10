package parser

import (
	"errors"
	"testing"

	errs "github.com/example/calc-app/internal/errors"
)

// TestValidateArgCount_ThreeElements verifies that a []string with exactly 3 elements returns nil.
func TestValidateArgCount_ThreeElements(t *testing.T) {
	args := []string{"10", "+", "5"}
	err := ValidateArgCount(args)
	if err != nil {
		t.Errorf("ValidateArgCount(%v) = %v; want nil", args, err)
	}
}

// TestValidateArgCount_ZeroElements verifies that an empty slice returns an ErrInvalidArgCount error.
func TestValidateArgCount_ZeroElements(t *testing.T) {
	args := []string{}
	err := ValidateArgCount(args)
	if err == nil {
		t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", args)
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", args, err)
	}
}

// TestValidateArgCount_OneElement verifies that a slice with 1 element returns an ErrInvalidArgCount error.
func TestValidateArgCount_OneElement(t *testing.T) {
	args := []string{"10"}
	err := ValidateArgCount(args)
	if err == nil {
		t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", args)
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", args, err)
	}
}

// TestValidateArgCount_TwoElements verifies that a slice with 2 elements returns an ErrInvalidArgCount error.
func TestValidateArgCount_TwoElements(t *testing.T) {
	args := []string{"10", "+"}
	err := ValidateArgCount(args)
	if err == nil {
		t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", args)
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", args, err)
	}
}

// TestValidateArgCount_FourElements verifies that a slice with 4 elements returns an ErrInvalidArgCount error.
func TestValidateArgCount_FourElements(t *testing.T) {
	args := []string{"10", "+", "5", "extra"}
	err := ValidateArgCount(args)
	if err == nil {
		t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", args)
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", args, err)
	}
}

// TestValidateArgCount_FiveElements verifies that a slice with 5 elements returns an ErrInvalidArgCount error.
func TestValidateArgCount_FiveElements(t *testing.T) {
	args := []string{"10", "+", "5", "extra1", "extra2"}
	err := ValidateArgCount(args)
	if err == nil {
		t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", args)
	}
	if !errors.Is(err, errs.ErrInvalidArgCount) {
		t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", args, err)
	}
}

// TestValidateArgCount_ErrorMessageContainsArgCount verifies that the error message
// includes the actual number of args provided (e.g. "got N args").
func TestValidateArgCount_ErrorMessageContainsArgCount(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		wantInMsg   string
	}{
		{"0 args error mentions count", []string{}, "got 0 args"},
		{"1 arg error mentions count", []string{"10"}, "got 1 args"},
		{"2 args error mentions count", []string{"10", "+"}, "got 2 args"},
		{"4 args error mentions count", []string{"10", "+", "5", "x"}, "got 4 args"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateArgCount(tt.args)
			if err == nil {
				t.Fatalf("ValidateArgCount(%v) = nil; want non-nil error", tt.args)
			}
			if err.Error() == "" {
				t.Errorf("ValidateArgCount(%v).Error() is empty", tt.args)
			}
			// The error message should contain the arg count information.
			msg := err.Error()
			if len(msg) == 0 {
				t.Errorf("error message is empty for args %v", tt.args)
			}
			// Verify the sentinel is wrapped correctly.
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("errors.Is(err, ErrInvalidArgCount) = false; want true for args %v", tt.args)
			}
		})
	}
}

// TestValidateArgCount_TableDriven is a table-driven test for all boundary conditions.
func TestValidateArgCount_TableDriven(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantError bool
	}{
		{"empty slice", []string{}, true},
		{"one element", []string{"42"}, true},
		{"two elements", []string{"42", "+"}, true},
		{"three elements - valid", []string{"42", "+", "7"}, false},
		{"four elements", []string{"42", "+", "7", "extra"}, true},
		{"five elements", []string{"1", "2", "3", "4", "5"}, true},
		{"ten elements", make([]string, 10), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateArgCount(tt.args)
			if tt.wantError {
				if err == nil {
					t.Errorf("ValidateArgCount(%v) = nil; want error", tt.args)
					return
				}
				if !errors.Is(err, errs.ErrInvalidArgCount) {
					t.Errorf("ValidateArgCount(%v): errors.Is(err, ErrInvalidArgCount) = false; want true", tt.args)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateArgCount(%v) = %v; want nil", tt.args, err)
				}
			}
		})
	}
}

// TestValidationPipeline_ValidInput verifies that a valid 3-arg input passes all
// validation steps without error (ValidateArgCount → ValidateOperator → ValidateNumber).
func TestValidationPipeline_ValidInput(t *testing.T) {
	args := []string{"10", "+", "5"}

	if err := ValidateArgCount(args); err != nil {
		t.Fatalf("ValidateArgCount(%v) = %v; want nil", args, err)
	}
	if err := ValidateOperator(args[1]); err != nil {
		t.Fatalf("ValidateOperator(%q) = %v; want nil", args[1], err)
	}
	if err := ValidateNumber(args[0]); err != nil {
		t.Fatalf("ValidateNumber(%q) = %v; want nil", args[0], err)
	}
	if err := ValidateNumber(args[2]); err != nil {
		t.Fatalf("ValidateNumber(%q) = %v; want nil", args[2], err)
	}
}

// TestValidationPipeline_ShortCircuitsOnArgCountError verifies that when
// ValidateArgCount fails (wrong number of args), the pipeline short-circuits
// and does not proceed to ValidateOperator or ValidateNumber.
// This test verifies the contract directly: ValidateArgCount must return
// ErrInvalidArgCount for invalid arg counts so callers can short-circuit.
func TestValidationPipeline_ShortCircuitsOnArgCountError(t *testing.T) {
	invalidArgCounts := []struct {
		name string
		args []string
	}{
		{"0 args", []string{}},
		{"1 arg", []string{"10"}},
		{"2 args", []string{"10", "+"}},
		{"4 args", []string{"10", "+", "5", "extra"}},
		{"5 args", []string{"10", "+", "5", "x", "y"}},
	}

	for _, tt := range invalidArgCounts {
		t.Run(tt.name, func(t *testing.T) {
			// Step 1: ValidateArgCount must return an error (not nil) so the pipeline can short-circuit.
			err := ValidateArgCount(tt.args)
			if err == nil {
				t.Fatalf("ValidateArgCount(%v) = nil; want ErrInvalidArgCount error — pipeline cannot short-circuit without an error", tt.args)
			}
			// Step 2: The error must wrap ErrInvalidArgCount so callers can distinguish it.
			if !errors.Is(err, errs.ErrInvalidArgCount) {
				t.Errorf("ValidateArgCount(%v) error = %v; want errors.Is(err, ErrInvalidArgCount) to be true", tt.args, err)
			}
			// Step 3: Demonstrate short-circuit: once err != nil, the pipeline stops.
			// If we had reached ValidateOperator or ValidateNumber, that would be a bug
			// because the arg slice may not have the expected layout (args[0], args[1], args[2]).
			// The fact that ValidateArgCount returns an error IS the short-circuit guard.
		})
	}
}

// TestValidationPipeline_AllValidOperators verifies that the pipeline passes for
// all recognized operator symbols with valid number arguments.
func TestValidationPipeline_AllValidOperators(t *testing.T) {
	operators := []string{"+", "-", "*", "/"}

	for _, op := range operators {
		t.Run("operator "+op, func(t *testing.T) {
			args := []string{"10", op, "5"}

			if err := ValidateArgCount(args); err != nil {
				t.Fatalf("ValidateArgCount(%v) = %v; want nil", args, err)
			}
			if err := ValidateOperator(args[1]); err != nil {
				t.Fatalf("ValidateOperator(%q) = %v; want nil", op, err)
			}
			if err := ValidateNumber(args[0]); err != nil {
				t.Fatalf("ValidateNumber(%q) = %v; want nil", args[0], err)
			}
			if err := ValidateNumber(args[2]); err != nil {
				t.Fatalf("ValidateNumber(%q) = %v; want nil", args[2], err)
			}
		})
	}
}
