package ast

import (
	"fmt"
	"strings"

	"ast/source"
	"ast/statement"
)

type Decl interface {
	source.Source
	Name() string
	String() string
}

type FnArg struct {
	source.Source
	Typ string
	Nam string
}

func (f *FnArg) String() string {
	return fmt.Sprintf("%s %s %s", source.SourceString(f.Source), f.Typ, f.Nam)
}

type FnReturn struct {
	source.Source
	Typ string
}

func (f *FnReturn) String() string {
	return fmt.Sprintf("%s %s", source.SourceString(f.Source), f.Typ)
}

type FnDecl struct {
	source.Source
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
	return fmt.Sprintf("Fn(%s)[%s](%s)->%v{%s}", source.SourceString(f.Source), f.Nam, strings.Join(argsStr, ","), f.Return, strings.Join(stmtStr, ","))
}
