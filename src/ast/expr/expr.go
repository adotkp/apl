package expr

import (
	"fmt"

	"ast/source"
	"types"
	"values"
)

// Expr represents an expression in the language.
type Expr interface {
	source.Source
	Check(*types.Context) (types.Type, error)
	Eval() (values.Value, error)
	String() string
}

// Value is an expression that is a constant value.
type Value struct {
	source.Source
	V values.Value
}

// Check returns the type of the constant value. Always returns a non-nil
// error.
func (v *Value) Check(c *types.Context) (types.Type, error) {
	return v.V.Type(), nil
}

// Eval returns the constant value. It always returns a non-nil error.
func (v *Value) Eval() (values.Value, error) {
	return v.V, nil
}

func (v *Value) String() string {
	return fmt.Sprintf("%v(%s)", v.V, source.String(v.Source))
}
