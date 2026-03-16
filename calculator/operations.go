package calculator

// Operation holds the operands and operator for a single calculation.
type Operation struct {
	Left     float64
	Operator string
	Right    float64
}

// Result holds the numeric value and human-readable expression of a calculation.
type Result struct {
	Value      float64
	Expression string
}
