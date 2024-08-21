package parser

import (
	"github.com/antlr4-go/antlr/v4"
)

func Parse(input *antlr.InputStream) (tree IProgramContext, err error) {
	defer func() {
		if r := recover(); r != nil {
			if rErr, ok := r.(*syntaxError); ok {
				tree, err = nil, rErr
			} else {
				panic(r)
			}
		}
	}()

	lexer := NewBoLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(newErrorListener())

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewBoParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(newErrorListener())

	return parser.Program(), nil
}

func ParseString(input string) (antlr.ParseTree, error) {
	return Parse(antlr.NewInputStream(input))
}

func ParseFile(filename string) (antlr.ParseTree, error) {
	fs, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
	}
	return Parse(&fs.InputStream)
}
