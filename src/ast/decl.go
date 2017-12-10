package ast

import (
	"fmt"
	"strings"

	"ast/statement"
	"types"
)

type Decl interface {
	Name() string
	String() string
}

type FnArg struct {
	Typ string
	Nam string
}

func (f *FnArg) String() string {
	return fmt.Sprintf("%s %s", f.Typ, f.Nam)
}

type FnReturn struct {
	Typ string
}

func (f *FnReturn) String() string {
	return f.Typ
}

type FnDecl struct {
	Nam    string
	Args   []*FnArg
	Return *FnReturn

	Statements []statement.Statement
}

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
	return fmt.Sprintf("Fn[%s](%s)->%v{%s}", f.Nam, strings.Join(argsStr, ","), f.Return, strings.Join(stmtStr, ","))
}

type TypeDecl struct {
	Nam string
	Typ types.Type
}

func (t *TypeDecl) Name() string {
	return t.Nam
}
