package main

import (
	"fmt"
	"os"

	"calculator/internal/calc"
	"calculator/internal/parser"
)

func main() {
	input, err := parser.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	op := calc.Operation{
		Left:     input.Left,
		Right:    input.Right,
		Operator: calc.Operator(input.Operator),
	}

	result, err := calc.NewEngine().Calculate(op)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(result)
}
