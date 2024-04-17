package builtins

import "bo/object"

// builtinArrayFirst returns the first element of the given array.
func builtinArrayFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return Error("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return Error("argument to `array_first` must be ARRAY, got %s", args[0].Type())
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
		return Error("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return Error("argument to `array_last` must be ARRAY, got %s", args[0].Type())
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
		return Error("wrong number of arguments. got=%d, want=2..n", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return Error("argument to `array_push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	arr.Elements = append(arr.Elements, args[1:]...)

	return arr
}

// builtinArrayPop removes the last element from the given array.
func builtinArrayPop(args ...object.Object) object.Object {
	if len(args) != 1 {
		return Error("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return Error("argument to `array_pop` must be ARRAY, got %s", args[0].Type())
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

// builtinArrayShift removes the first element from the given array.
func builtinArrayShift(args ...object.Object) object.Object {
	if len(args) != 1 {
		return Error("wrong number of arguments. got=%d, want=1", len(args))
	}

	if args[0].Type() != object.ARRAY_OBJ {
		return Error("argument to `array_shift` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		shifted := arr.Elements[0]
		arr.Elements = arr.Elements[1:]
		return shifted
	}

	return NIL
}
