package main

import (
	"bo/repl"
	"os"
)

func main() {
	repl := repl.New(os.Stdin, os.Stdout, os.Stderr)
	repl.PrintBanner()
	repl.Start()
}
