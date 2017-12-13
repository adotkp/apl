package values

import (
	"fmt"

	"types"
)

// Value represents a value in the language.
type Value interface {
	Type() types.Type
	String() string
}

// Int is an integer value.
type Int struct {
	V int
}

// Type returns the Int type.
func (i *Int) Type() types.Type {
	return &types.Int{}
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.V)
}

// Bool is a boolean value.
type Bool struct {
	V bool
}

// Type returns the Bool type.
func (b *Bool) Type() types.Type {
	return &types.Bool{}
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.V)
}

// String is a string value.
type String struct {
	V string
}

// Type returns the String type.
func (s *String) Type() types.Type {
	return &types.String{}
}

func (s *String) String() string {
	return s.V
}
