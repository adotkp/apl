package runtime

import (
	"io"

	"ast"
)

type Executor struct {
}

func (e *Executor) TypeCheck(r io.RuneScanner) error {
	l := parser.NewLexer(r)
	p := parser.NewParser(l.Tokens())
	file, err := p.Do()
	if err != nil {
		return err
	}
	return file.TypeCheck()
}
