package runtime

import (
	"path/filepath"

	"ast"
	"parser"
	"types"
)

// Executor loads, checks, and runs the language.
type Executor struct {
	loader Loader
	tc     *types.Context
	files  map[string]*ast.File
}

// NewExecutor returns a new Executor.
func NewExecutor(l Loader) *Executor {
	return &Executor{
		loader: l,
		tc:     types.NewContext(),
		files:  make(map[string]*ast.File),
	}
}

// Check statically checks an import path.
// TODO(adi): Check for import cycles.
func (e *Executor) Check(path string) error {
	if _, ok := e.files[path]; ok {
		return nil
	}
	r, err := e.loader.Load(path)
	if err != nil {
		return err
	}
	defer r.Close()
	_, name := filepath.Split(path)
	p := parser.NewParser(parser.NewLexer(name, r).Tokens())
	file, err := p.Do()
	if err != nil {
		return err
	}
	e.files[path] = file
	for _, imp := range file.Imports {
		if err := e.Check(imp.Name); err != nil {
			return err
		}
	}
	return file.Check(e.tc)
}
