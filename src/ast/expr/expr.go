package expr

import (
	"fmt"

	"ast/source"
	"values"
)

type Expr interface {
	source.Source
	Eval() (values.Value, error)
	String() string
}

type Value struct {
	source.Source
	V values.Value
}

func (v *Value) Eval() (values.Value, error) {
	return v.V, nil
}

func (v *Value) String() string {
	return fmt.Sprintf("%v(%s)", v.V, source.SourceString(v.Source))
}
