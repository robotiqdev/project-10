package calc

// Operator stub — constants intentionally left empty for TDD (implementation pending).
type Operator string

const (
	OpAdd      Operator = "+"
	OpSubtract Operator = "-"
	OpMultiply Operator = "*"
	OpDivide   Operator = "/"
)

// Operation stub — fields present for compilation but not fully wired.
type Operation struct {
	Left     float64
	Right    float64
	Operator Operator
}
