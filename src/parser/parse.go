package parser

import (
	"errors"
	"fmt"

	"ast"
)

var (
	errEOF = errors.New("unexpected eof")
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
		return ret, errEOF
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

// P is the parser that converts a token stream to the AST.
type P struct {
	tokens *gettoken
}

// NewParser returns a new P.
func NewParser(tokens <-chan Token) *P {
	return &P{
		tokens: &gettoken{tokens, Token{}, false},
	}
}

// Do parses the token stream and returns the AST, or an error if parsing error
// occurred.
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

// TokenSource implements Source from the ast/source package from the Token.
type TokenSource struct {
	Token
}

// Pos returns the absolute rune-offset (0-indexed).
func (t TokenSource) Pos() int {
	return t.Token.Pos
}

// Line returns the line number (0-indexed).
func (t TokenSource) Line() int {
	return t.Token.Line
}

// LinePos returns the rune-offset within the line (0-indexed).
func (t TokenSource) LinePos() int {
	return t.Token.LinePos
}

// File returns the name of the source file.
func (t TokenSource) File() string {
	return t.Token.File
}

// Errf formats an error with the positional information prepended.
func (t TokenSource) Errf(format string, args ...interface{}) error {
	return fmt.Errorf("%s:%d:%d %s", t.File(), t.Line()+1, t.LinePos()+1, fmt.Sprintf(format, args...))
}
