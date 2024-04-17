package evaluator

import "bolang/object"

// Bo-Lang built-in functions.
var builtinFunctions = map[string]*object.Builtin{
	"len":         object.NewBuiltin(builtinLen),
	"array_first": object.NewBuiltin(builtinArrayFirst),
	"array_last":  object.NewBuiltin(builtinArrayLast),
	"array_push":  object.NewBuiltin(builtinArrayPush),
	"array_pop":   object.NewBuiltin(builtinArrayPop),
}

// builtinLen returns the length of the given object.
func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}
	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}
	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

// builtinArrayFirst returns the first element of the given array.
func builtinArrayFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `array_first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NIL
}

// builtinArrayLast returns the last element of the given array.
func builtinArrayLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `array_last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NIL
}

// builtinArrayPush appends the given elements to the given array.
func builtinArrayPush(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, want=2..n", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `array_push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	arr.Elements = append(arr.Elements, args[1:]...)

	return arr
}

// builtinArrayPop removes the last element from the given array.
func builtinArrayPop(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `array_pop` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		popped := arr.Elements[length-1]
		arr.Elements = arr.Elements[:length-1]
		return popped
	}

	return NIL
}
