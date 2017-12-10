package ast

type File struct {
	Imports []Import
	Decls   []Decl
}

type Import struct {
	Name string
}
