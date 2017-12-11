package statement

import (
	"fmt"

	"ast/expr"
)

type Statement interface {
	TypeCheck() error
	String() string
}

type Import struct {
	Name string
}

func (i *Import) TypeCheck() error {
	return nil
}

func (i *Import) String() string {
	return "import " + i.Name
}

type Return struct {
	Expr expr.Expr
}

func (r *Return) Exec() error {
	return nil
}

func (r *Return) String() string {
	return fmt.Sprintf("return %v", r.Expr)
}

type FnCall struct {
	Nam    string
	Params []expr.Expr
}

func (f *FnCall) Exec() error {
	return nil
}

func (f *FnCall) String() string {
	return fmt.Sprintf("FnCall(%s:%v)", f.Nam, f.Params)
}

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
