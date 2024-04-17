package object

import "bytes"

// The Array struct represents an array object.
type Array struct {
	Elements []Object
}

// Type returns the type of the object.
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect returns a string-representation of the object.
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	out.WriteString("[")
	for i, el := range ao.Elements {
		out.WriteString(el.Inspect())
		if i < len(ao.Elements)-1 {
			out.WriteString(", ")
		}
	}

	out.WriteString("]")
	return out.String()
}
