package main

import (
	"fmt"
	"os"

	"github.com/example/calc-app/internal/calc"
	"github.com/example/calc-app/internal/parser"
)

func main() {
	args := os.Args[1:]

	input, err := parser.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	engine := calc.NewEngine()
	result, err := engine.Calculate(calc.Operation{
		Left:     input.Left,
		Operator: input.Operator,
		Right:    input.Right,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "%s %s %s = %s\n", args[0], args[1], args[2], result)
}
