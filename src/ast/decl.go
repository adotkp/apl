package ast

import (
	"statement"
	"types"
)

type Decl interface {
	Name() string
}

type FnDecl struct {
	name   string
	Args   []types.Type
	Return types.Type

	Statements []statement.Statement
}

func (f *FnDecl) Name() string {
	return f.name
}

type TypeDecl struct {
	name string
	Typ  types.Type
}

func (t *TypeDecl) Name() string {
	return t.name
}
