package evaluator

import "bolang/object"

// Bo-Lang built-in functions.
var builtinFunctions = map[string]*object.Builtin{
	"len": {Fn: builtinLen},
}

// builtinLen returns the length of the given object.
func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}
