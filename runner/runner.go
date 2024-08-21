package runner

import (
	"github.com/antlr4-go/antlr/v4"
)

func RunProgram(input antlr.ParseTree) {
	if input == nil {
		return
	}

	visitor := NewBoVisitor()
	visitor.Visit(input)
}
