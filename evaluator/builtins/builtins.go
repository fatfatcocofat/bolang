package builtins

import (
	"bolang/object"
	"fmt"
)

var (
	// NIL represents the absence of a value.
	NIL = &object.Nil{}
)

// Bo-Lang built-in functions.
var BuiltinsFn = map[string]*object.Builtin{
	"len":         object.NewBuiltin(builtinLen),
	"array_first": object.NewBuiltin(builtinArrayFirst),
	"array_last":  object.NewBuiltin(builtinArrayLast),
	"array_push":  object.NewBuiltin(builtinArrayPush),
	"array_pop":   object.NewBuiltin(builtinArrayPop),
	"array_shift": object.NewBuiltin(builtinArrayShift),
}

// builtinLen returns the length of the given object.
func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return Error("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return Error("argument to `len` not supported, got %s", args[0].Type())
	}
}

// Error returns a new object.Error with the given format and arguments.
func Error(format string, a ...interface{}) *object.Error {
	msg := fmt.Sprintf(format, a...)
	return &object.Error{Message: msg}
}
