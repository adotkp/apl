package source

import (
	"fmt"
)

// Source represents positional information of an AST node.
type Source interface {
	File() string
	Pos() int
	Line() int
	LinePos() int
}

// String returns a string representation of a given Source.
func String(s Source) string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("@<%s:%d:%d:%d>", s.File(), s.Line()+1, s.LinePos()+1, s.Pos())
}
