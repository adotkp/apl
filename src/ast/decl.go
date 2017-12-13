package ast

import (
	"fmt"
	"strings"

	"ast/source"
	"ast/statement"
	"types"
)

// Decl represents a type or function declaration.
type Decl interface {
	source.Source
	Name() string
	String() string
	Check(c *types.Context) error
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

// Check validates the arg and return types of the declared function, as well
// as the statements inside the function. The declared function is registered
// prior to checking the statements to support recursive calls.
func (f *FnDecl) Check(c *types.Context) error {
	var argTypes []types.Type
	for _, arg := range f.Args {
		typ, err := c.Get(arg.Typ)
		if err != nil {
			return arg.Errf(err.Error())
		}
		argTypes = append(argTypes, typ)
	}
	var retType types.Type
	var err error
	if f.Return != nil {
		retType, err = c.Get(f.Return.Typ)
		if err != nil {
			return f.Return.Errf(err.Error())
		}
	}
	// Register before evaluating statements to support recursion.
	err = c.Add(f.Nam, &types.Func{
		Args:   argTypes,
		Return: retType,
	})
	if err != nil {
		return f.Errf(err.Error())
	}
	for _, stmt := range f.Statements {
		_, err := stmt.Check(c)
		if err != nil {
			return err
		}
	}
	return nil
}
