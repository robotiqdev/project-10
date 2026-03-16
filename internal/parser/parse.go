package parser

import "fmt"

// Input holds the parsed values from command-line arguments.
type Input struct {
	Left     float64
	Right    float64
	Operator string
}

// Parse parses command-line arguments into an Input.
// Expected format: [leftNumber, operator, rightNumber].
func Parse(args []string) (Input, error) {
	return Input{}, fmt.Errorf("not implemented")
}
