package object

import (
	"bo/ast"
	"fmt"
)

// The Function struct represents a function object.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	return fmt.Sprintf("fn(%s) {\n%s\n}",
		f.Parameters, f.Body.String())
}
