package source

import (
	"fmt"
)

type Source interface {
	File() string
	Pos() int
	Line() int
	LinePos() int
}

func SourceString(s Source) string {
	if s == nil {
		return "nil"
	}
	return fmt.Sprintf("@<%s:%d:%d:%d>", s.File(), s.Line()+1, s.LinePos()+1, s.Pos())
}
