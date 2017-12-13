package runtime

import (
	"io"

	"parser"
	"types"
)

// Executor loads, checks, and runs the language.
type Executor struct {
}

// Check statically type-checks the source.
func (e *Executor) Check(fileName string, r io.RuneScanner) error {
	l := parser.NewLexer(fileName, r)
	p := parser.NewParser(l.Tokens())
	file, err := p.Do()
	if err != nil {
		return err
	}
	c := types.NewContext()
	return file.Check(c)
}
