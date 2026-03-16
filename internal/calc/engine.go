package calc

// Engine orchestrates arithmetic operations and result formatting.
type Engine struct{}

// NewEngine constructs and returns a new Engine.
func NewEngine() *Engine {
	return &Engine{}
}

// Calculate performs the operation described by op and returns a formatted
// result string. It returns an error if the operation is invalid (e.g.
// division by zero or unknown operator).
func (e *Engine) Calculate(op Operation) (string, error) {
	panic("not implemented")
}
