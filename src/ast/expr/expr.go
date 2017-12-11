package expr

import (
	"values"
)

type Expr interface {
	Eval() (values.Value, error)
	String() string
}

type Value struct {
	V values.Value
}

func (v *Value) Eval() (values.Value, error) {
	return v.V, nil
}

func (v *Value) String() string {
	return v.V.String()
}
