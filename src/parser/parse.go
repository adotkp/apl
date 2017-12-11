package parser

import (
	"errors"
	"fmt"

	"ast"
)

var (
	errEof = errors.New("unexpected eof")
)

type gettoken struct {
	tokens <-chan Token
	prev   Token
	undo   bool
}

func (gt *gettoken) get() (Token, error) {
	if gt.undo {
		gt.undo = false
		return gt.prev, nil
	}
	ret, isOpen := <-gt.tokens
	if !isOpen {
		return ret, errEof
	}
	if ret.Err != nil {
		return ret, ret.Err
	}
	gt.prev = ret
	return ret, nil
}

func (gt *gettoken) unread() {
	if gt.undo {
		panic("cannot undo token stream twice")
	}
	gt.undo = true
}

type P struct {
	tokens *gettoken
}

func NewParser(tokens <-chan Token) *P {
	return &P{
		tokens: &gettoken{tokens, Token{}, false},
	}
}

func (p *P) Do() (*ast.File, error) {
	tok, err := p.tokens.get()
	if err != nil {
		return nil, err
	}
	p.tokens.unread()
	imports, err := p.parseImports()
	if err != nil {
		return nil, err
	}
	decls, err := p.parseDecls()
	if err != nil {
		return nil, err
	}
	return &ast.File{
		Source:  TokenSource{tok},
		Imports: imports,
		Decls:   decls,
	}, nil
}

func (p *P) errf(t Token, format string, args ...interface{}) error {
	return fmt.Errorf("error at pos %d (%s): %s", t.Pos, string(t.Lit), fmt.Sprintf(format, args...))
}

func (p *P) consume(typs ...TokenType) (bool, Token, error) {
	tok, err := p.tokens.get()
	if err != nil {
		return false, tok, err
	}
	for _, typ := range typs {
		if typ == tok.Typ {
			return false, tok, nil
		}
	}
	return true, tok, p.errf(tok, "did not expect %v", tok.Typ)
}

func (p *P) consumeText() (string, Token, error) {
	tok, err := p.tokens.get()
	if err != nil {
		return "", tok, err
	}
	if tok.Typ != TokenText {
		return "", tok, p.errf(tok, "expected %v, got %v", TokenText, tok.Typ)
	}
	return string(tok.Lit), tok, nil
}

type TokenSource struct {
	Token
}

func (t TokenSource) Pos() int {
	return t.Token.Pos
}

func (t TokenSource) Line() int {
	return t.Token.Line
}

func (t TokenSource) LinePos() int {
	return t.Token.LinePos
}

func (t TokenSource) File() string {
	return t.Token.File
}
