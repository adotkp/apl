package ast

import (
	"fmt"
	"strings"

	"ast/source"
	"ast/statement"
)

// Decl represents a type or function declaration.
type Decl interface {
	source.Source
	Name() string
	String() string
}

// FnArg is an argument to a function.
type FnArg struct {
	source.Source
	Typ string
	Nam string
}

func (f *FnArg) String() string {
	return fmt.Sprintf("%s %s %s", source.String(f.Source), f.Typ, f.Nam)
}

// FnReturn is the return type of a function.
type FnReturn struct {
	source.Source
	Typ string
}

func (f *FnReturn) String() string {
	return fmt.Sprintf("%s %s", source.String(f.Source), f.Typ)
}

// FnDecl is a declaration for a function.
type FnDecl struct {
	source.Source
	Nam    string
	Args   []*FnArg
	Return *FnReturn // If null, does no return anything.

	Statements []statement.Statement
}

// Name returns the name of this function.
func (f *FnDecl) Name() string {
	return f.Nam
}

func (f *FnDecl) String() string {
	var argsStr []string
	for _, arg := range f.Args {
		argsStr = append(argsStr, arg.String())
	}
	var stmtStr []string
	for _, stmt := range f.Statements {
		stmtStr = append(stmtStr, stmt.String())
	}
	return fmt.Sprintf("Fn(%s)[%s](%s)->%v{%s}", source.String(f.Source), f.Nam, strings.Join(argsStr, ","), f.Return, strings.Join(stmtStr, ","))
}
