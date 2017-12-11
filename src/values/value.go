package values

import (
	"fmt"

	"types"
)

type Value interface {
	Type() types.Type
	String() string
}

type Int struct {
	V int
}

func (i *Int) Type() types.Type {
	return &types.Int{}
}

func (i *Int) String() string {
	return fmt.Sprintf("%d", i.V)
}

type Bool struct {
	V bool
}

func (b *Bool) Type() types.Type {
	return &types.Bool{}
}

func (b *Bool) String() string {
	return fmt.Sprintf("%t", b.V)
}

type String struct {
	V string
}

func (s *String) Type() types.Type {
	return &types.String{}
}

func (s *String) String() string {
	return s.V
}
