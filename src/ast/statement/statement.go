package statement

import (
	"fmt"

	"ast/expr"
)

type Statement interface {
	Exec() error
	String() string
}

type Import struct {
	Name string
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
