package runtime

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"ast"
	"parser"
	"types"
)

// Executor loads, checks, and runs the language.
type Executor struct {
  l Loader
	paths []string
	tc    *types.Context
	files map[string]*ast.File
}

// NewExecutor returns a new Executor.
func NewExecutor(l Loader, Paths []string) *Executor {
	return &Executor{
    l: l,
		paths: Paths,
		tc:    types.NewContext(),
		files: make(map[string]*ast.File),
	}
}

// Check statically checks an import path.
func (e *Executor) Check(importPath string) error {
	_, name := filepath.Split(importPath)
	path, err := e.abs(importPath)
	if err != nil {
		return err
	}
	if _, ok := e.files[path]; ok {
		// TODO(adi): Cycle-detect imports.
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	ast, err := e.CheckWithReader(name, bufio.NewReader(f))
	if err != nil {
		return err
	}
	e.files[path] = ast
	return nil
}

func (e *Executor) abs(filePath string) (string, error) {
	for _, path := range e.paths {
		absPath := filepath.Join(path, filePath)
		_, err := os.Stat(absPath)
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", nil
		}
		return absPath, nil
	}
	return "", fmt.Errorf("cannot load: %s", filePath)
}

// CheckWithReader statically checks the given byte stream.
func (e *Executor) CheckWithReader(fileName string, r io.RuneScanner) (*ast.File, error) {
	p := parser.NewParser(parser.NewLexer(fileName, r).Tokens())
	file, err := p.Do()
	if err != nil {
		return nil, err
	}
	for _, imp := range file.Imports {
		if err := e.Check(imp.Name); err != nil {
			return nil, err
		}
	}
	return file, file.Check(e.tc)
}
