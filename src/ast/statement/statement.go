package statement

import (
	"fmt"

	"ast/expr"
	"ast/source"
	"types"
)

// Statement represents a statement in the language.
type Statement interface {
	source.Source
	String() string
	Check(*types.Context) (types.Type, error)
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

// Check always returns nil.
func (i *Import) Check(c *types.Context) (types.Type, error) {
	return nil, nil
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

// Check checks the expression to be returned.
func (r *Return) Check(c *types.Context) (types.Type, error) {
	return r.Expr.Check(c)
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

// Check validates the params for the function call and returns the return type
// of the function.
func (f *FnCall) Check(c *types.Context) (types.Type, error) {
	typ, err := c.Get(f.Nam)
	if err != nil {
		return nil, f.Errf(err.Error())
	}
	fnTyp, ok := typ.(*types.Func)
	if !ok {
		return nil, f.Errf("%s is %v, not func", f.Nam, typ)
	}
	if len(f.Params) != len(fnTyp.Args) {
		return nil, f.Errf("%s expects %d params, not %d", f.Nam, len(fnTyp.Args), len(f.Params))
	}
	for i, arg := range fnTyp.Args {
		paramTyp, err := f.Params[i].Check(c)
		if err != nil {
			return nil, err
		}
		if !paramTyp.Equals(arg) {
			return nil, f.Errf("%s param #%d expects %v, not %v", f.Nam, i+1, arg, paramTyp)
		}
	}
	return fnTyp.Return, nil
}
