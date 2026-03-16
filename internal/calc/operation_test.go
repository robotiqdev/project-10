package calc

import "testing"

// TestOperatorConstants verifies each Operator constant has the expected string value.
func TestOperatorConstants(t *testing.T) {
	tests := []struct {
		name     string
		got      Operator
		expected Operator
	}{
		{"OpAdd is '+'", OpAdd, "+"},
		{"OpSubtract is '-'", OpSubtract, "-"},
		{"OpMultiply is '*'", OpMultiply, "*"},
		{"OpDivide is '/'", OpDivide, "/"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("expected Operator %q, got %q", tt.expected, tt.got)
			}
		})
	}
}

// TestOperatorIsStringType verifies that Operator is based on the string type,
// allowing string conversions and comparisons.
func TestOperatorIsStringType(t *testing.T) {
	var op Operator = "+"
	if string(op) != "+" {
		t.Errorf("expected Operator to be convertible to string '+', got %q", string(op))
	}
}

// TestOperatorConstantsAreDistinct verifies all four constants have different values.
func TestOperatorConstantsAreDistinct(t *testing.T) {
	ops := []Operator{OpAdd, OpSubtract, OpMultiply, OpDivide}
	seen := make(map[Operator]bool)
	for _, op := range ops {
		if seen[op] {
			t.Errorf("duplicate Operator value: %q", op)
		}
		seen[op] = true
	}
}

// TestOperationStructZeroValues verifies that a zero-valued Operation has the
// expected zero values for each field: 0.0 for Left and Right, "" for Operator.
func TestOperationStructZeroValues(t *testing.T) {
	var op Operation

	if op.Left != 0.0 {
		t.Errorf("expected zero-value Left to be 0.0, got %v", op.Left)
	}

	if op.Right != 0.0 {
		t.Errorf("expected zero-value Right to be 0.0, got %v", op.Right)
	}

	if op.Operator != "" {
		t.Errorf("expected zero-value Operator to be empty string, got %q", op.Operator)
	}
}

// TestOperationStructFieldAssignment verifies all fields of Operation are settable
// and readable (i.e., they are exported and of the correct types).
func TestOperationStructFieldAssignment(t *testing.T) {
	op := Operation{
		Left:     10.5,
		Right:    3.2,
		Operator: OpAdd,
	}

	if op.Left != 10.5 {
		t.Errorf("expected Left 10.5, got %v", op.Left)
	}

	if op.Right != 3.2 {
		t.Errorf("expected Right 3.2, got %v", op.Right)
	}

	if op.Operator != OpAdd {
		t.Errorf("expected Operator %q, got %q", OpAdd, op.Operator)
	}
}

// TestOperationFieldTypes verifies Left and Right accept float64 values
// (including negative numbers and zero).
func TestOperationFieldTypes(t *testing.T) {
	cases := []struct {
		left     float64
		right    float64
		operator Operator
	}{
		{0, 0, OpAdd},
		{-1.5, 2.5, OpSubtract},
		{100, -0.001, OpMultiply},
		{1e10, 1e-10, OpDivide},
	}

	for _, c := range cases {
		op := Operation{Left: c.left, Right: c.right, Operator: c.operator}
		if op.Left != c.left {
			t.Errorf("Left: expected %v, got %v", c.left, op.Left)
		}
		if op.Right != c.right {
			t.Errorf("Right: expected %v, got %v", c.right, op.Right)
		}
		if op.Operator != c.operator {
			t.Errorf("Operator: expected %q, got %q", c.operator, op.Operator)
		}
	}
}
