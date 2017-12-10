package ast

import (
	"fmt"
	"strings"

	"ast/statement"
)

type File struct {
	Imports []*statement.Import
	Decls   []Decl
}

func (f *File) String() string {
	var importStrs []string
	for _, imp := range f.Imports {
		importStrs = append(importStrs, imp.String())
	}
	var declStrs []string
	for _, decl := range f.Decls {
		declStrs = append(declStrs, decl.String())
	}
	return fmt.Sprintf("Imports(%s) Decls(%s)", strings.Join(importStrs, ","), strings.Join(declStrs, ","))
}
