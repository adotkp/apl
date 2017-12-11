package statement

import (
	"fmt"

	"ast/expr"
	"ast/source"
)

type Statement interface {
	// TypeCheck() error
	source.Source
	String() string
}

type Import struct {
	source.Source
	Name string
}

func (i *Import) TypeCheck() error {
	return nil
}

func (i *Import) String() string {
	return fmt.Sprintf("import(%s) %s", source.SourceString(i.Source), i.Name)
}

type Return struct {
	source.Source
	Expr expr.Expr
}

func (r *Return) Exec() error {
	return nil
}

func (r *Return) String() string {
	return fmt.Sprintf("return(%s) %v", source.SourceString(r.Source), r.Expr)
}

type FnCall struct {
	source.Source
	Nam    string
	Params []expr.Expr
}

func (f *FnCall) Exec() error {
	return nil
}

func (f *FnCall) String() string {
	return fmt.Sprintf("FnCall(%s:%s:%v)", source.SourceString(f.Source), f.Nam, f.Params)
}

/*
func (f *FnCall) TypeCheck(ctx types.Context) error {
	typ, err := ctx.Get(f.Nam)
	if err != nil {
		return err
	}
	fnTyp := typ.(*types.Function)
	if len(fnTyp.Args) != len(f.Params) {
		return
	}
}
*/
