package expr

import (
	"fmt"

	"ast/source"
	"values"
)

// Expr represents an expression in the language.
type Expr interface {
	source.Source
	Eval() (values.Value, error)
	String() string
}

// Value is an expression that is a constant value.
type Value struct {
	source.Source
	V values.Value
}

// Eval returns the constant value. It always returns a non-nil error.
func (v *Value) Eval() (values.Value, error) {
	return v.V, nil
}

func (v *Value) String() string {
	return fmt.Sprintf("%v(%s)", v.V, source.String(v.Source))
}
