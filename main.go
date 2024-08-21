package main

import (
	"bo/parser"
	"bo/runner"
)

func main() {
	input := `
	require <bo/fmt>
	require "path/to/file.bo"
	require 'path/to/file.bo'

	// Single line comment

	/*
	 * Multi-line comment
	 */

	int a = 1
	float b = 1.0
	bool c = true
	string b = "Hello, World!"

	// Built-in functions
	println(b)

	println("x:", x, "y:", y)
	`
	prog, err := parser.ParseString(input)
	if err != nil {
		panic(err)
	}

	runner.RunProgram(prog)
}
