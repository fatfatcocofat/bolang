package object

// Builtin represents a built-in function.
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns the object's type.
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect returns a string-representation of the object.
func (b *Builtin) Inspect() string { return "<bo-builtin-fn>" }

// BuiltinFunction represents a built-in function.
type BuiltinFunction func(args ...Object) Object

// Create a new built-in function.
func NewBuiltin(fn BuiltinFunction) *Builtin {
	return &Builtin{Fn: fn}
}
