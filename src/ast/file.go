package ast

import (
	"fmt"
	"strings"

	"ast/source"
	"ast/statement"
)

// File is the root of the AST. It represents the result of parsing a single
// source file.
type File struct {
	source.Source
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
	return fmt.Sprintf(
		"File(%s) Imports(%s) Decls(%s)",
		source.String(f.Source),
		strings.Join(importStrs, ","),
		strings.Join(declStrs, ","))
}
