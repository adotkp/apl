package lex

import (
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer(strings.NewReader("foobar"))
	for token := range l.Tokens() {
		if string(token.Lit) != "test token" {
			t.Errorf("got unexpected token %v", token)
		}
	}
}
