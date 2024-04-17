package object

import (
	"hash/fnv"
	"strings"
)

// Map is a Bo-Lang map object.
type Map struct {
	// Pairs is a map of keys and values.
	Pairs map[MapKey]MapPair
}

// MapPair is a key-value pair in a Bo-Lang map.
type MapPair struct {
	// Key is the key of the pair.
	Key Object
	// Value is the value of the pair.
	Value Object
}

// Type returns the type of the object.
func (m *Map) Type() ObjectType { return MAP_OBJ }

// Inspect returns a stringified version of the object.
func (m *Map) Inspect() string {
	pairs := []string{}

	for _, pair := range m.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+": "+pair.Value.Inspect())
	}

	return "{" + strings.Join(pairs, ", ") + "}"
}

// MapKey is a key in a Bo-Lang map.
type MapKey struct {
	Type  ObjectType
	Value uint64
}

// MapKey returns a new map key.
func (s *String) MapKey() MapKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return MapKey{Type: s.Type(), Value: h.Sum64()}
}

// MapKey returns a new map key.
func (i *Integer) MapKey() MapKey {
	return MapKey{Type: i.Type(), Value: uint64(i.Value)}
}

// MapKey returns a new map key.
func (b *Boolean) MapKey() MapKey {
	var value uint64
	if b.Value {
		value = 1
	}
	return MapKey{Type: b.Type(), Value: value}
}

// MapKey returns a new map key.
func (f *Float) MapKey() MapKey {
	return MapKey{Type: f.Type(), Value: uint64(f.Value)}
}

// Mapable is an interface for objects that can be used as map keys.
type Mapable interface {
	MapKey() MapKey
}
