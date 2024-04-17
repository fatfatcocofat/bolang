package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnvironment creates a new environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment enclosed in another environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get returns the object associated with the given key.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets the object associated with the given key.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
