package expr

import (
	"values"
)

type Expr interface {
	Eval() (values.Value, error)
	String() string
}

type Value struct {
	Text string
}

func (v *Value) Eval() (values.Value, error) {
	return nil, nil
}

func (v *Value) String() string {
	return v.Text
}
