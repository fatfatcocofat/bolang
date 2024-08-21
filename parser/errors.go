package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type syntaxError struct {
	line   int
	column int
	msg    string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("Syntax error at line %d:%d: %s", e.line, e.column, e.msg)
}

type errorListener struct {
	*antlr.DefaultErrorListener
}

func newErrorListener() *errorListener {
	return &errorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (l *errorListener) SyntaxError(
	_ antlr.Recognizer, // recognizer
	_ interface{}, // offendingSymbol
	line, column int, // line, column
	msg string, // message
	_ antlr.RecognitionException, // exception
) {
	panic(&syntaxError{line: line, column: column, msg: msg})
}
