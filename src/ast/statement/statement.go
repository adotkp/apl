package statement

import (
	"fmt"

	"ast/expr"
	"ast/source"
)

// Statement represents a statement in the language.
type Statement interface {
	source.Source
	String() string
}

// Import is an import statement. Import statements load additional
// namespaced libs relative to the root of the project.
type Import struct {
	source.Source
	Name string
}

func (i *Import) String() string {
	return fmt.Sprintf("import(%s) %s", source.String(i.Source), i.Name)
}

// Return is a return. Return statements return a value from a function. Return
// statements must be the last statement in a function.
type Return struct {
	source.Source
	Expr expr.Expr
}

func (r *Return) String() string {
	return fmt.Sprintf("return(%s) %v", source.String(r.Source), r.Expr)
}

// FnCall is a statement to call a function.
type FnCall struct {
	source.Source
	Nam    string
	Params []expr.Expr
}

func (f *FnCall) String() string {
	return fmt.Sprintf("FnCall(%s:%s:%v)", source.String(f.Source), f.Nam, f.Params)
}
