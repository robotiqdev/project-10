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

func Add(a, b float64) float64      { return a + b }
func Subtract(a, b float64) float64 { return a - b }
func Multiply(a, b float64) float64 { return a * b }
func Divide(a, b float64) float64   { return a / b }
