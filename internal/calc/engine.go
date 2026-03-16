package calc

import "fmt"

// Operator represents a supported mathematical operator.
type Operator string

// Operation represents a mathematical operation to perform.
type Operation struct {
	Left     float64
	Right    float64
	Operator Operator
}

// Engine performs mathematical calculations.
type Engine struct{}

// NewEngine creates a new Engine instance.
func NewEngine() *Engine {
	return &Engine{}
}

// Calculate performs the operation and returns a formatted string result.
func (e *Engine) Calculate(op Operation) (string, error) {
	return "", fmt.Errorf("not implemented")
}
