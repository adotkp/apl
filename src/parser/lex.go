package lex

import (
	"errors"
	"fmt"
	"io"
	"unicode"
)

type TokenType int64

const (
	TokenError TokenType = iota << 1
	TokenBraceOpen
	TokenBraceClose
	TokenParensOpen
	TokenParensClose
	TokenComma
	TokenSemicolon
	TokenString
	TokenNumber
	TokenIf
	TokenElse
	TokenFunc
	TokenImport
	TokenText
)

var keywords map[string]TokenType

func init() {
	keywords = map[string]TokenType{
		"if":     TokenIf,
		"else":   TokenElse,
		"func":   TokenFunc,
		"import": TokenImport,
	}
}

type Token struct {
	Typ TokenType
	Lit []rune
	Pos int
	End int
	Err error
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%d,%s)", t.Typ, t.Lit)
}

type Lexer struct {
	scanner io.RuneScanner
	pos     int
}

func NewLexer(scanner io.RuneScanner) *Lexer {
	return &Lexer{
		scanner: scanner,
	}
}

func (l *Lexer) Tokens() <-chan Token {
	tokens := make(chan Token)
	go func() {
		for {
			start := l.pos
			token := l.next()
			token.Pos = start
			token.End = l.pos
			if token.Err != nil {
				if token.Err != io.EOF {
					tokens <- token
				}
				close(tokens)
				return
			}
			tokens <- token
		}
		close(tokens)
	}()
	return tokens
}

func (l *Lexer) read() (rune, error) {
	r, _, err := l.scanner.ReadRune()
	if err != nil {
		return r, err
	}
	l.pos++
	return r, nil
}

func (l *Lexer) unread() error {
	err := l.scanner.UnreadRune()
	if err != nil {
		return err
	}
	l.pos--
	return nil
}

func (l *Lexer) consumeWhitespace() error {
	for {
		r, err := l.read()
		if err != nil {
			return err
		}
		if !unicode.IsSpace(r) {
			return l.unread()
		}
	}
}

func (l *Lexer) next() Token {
	err := l.consumeWhitespace()
	if err != nil {
		return l.err(err)
	}

	r, err := l.read()
	if err != nil {
		return l.err(err)
	}
	switch r {
	case '{':
		return l.emitSymbol(r, TokenBraceOpen)
	case '}':
		return l.emitSymbol(r, TokenBraceClose)
	case '(':
		return l.emitSymbol(r, TokenParensOpen)
	case ')':
		return l.emitSymbol(r, TokenParensClose)
	case ';':
		return l.emitSymbol(r, TokenSemicolon)
	case ',':
		return l.emitSymbol(r, TokenComma)
	case '"':
		return l.emitString()
	default:
		err = l.unread()
		if err != nil {
			return l.err(err)
		}
		return l.emitAlphaNum()
	}
}

func (l *Lexer) err(err error) Token {
	return Token{
		Err: err,
		Typ: TokenError,
	}
}

func (l *Lexer) emitSymbol(r rune, typ TokenType) Token {
	return Token{Typ: typ, Lit: []rune{r}}
}

func (l *Lexer) emitString() Token {
	t := l.emitUntil(func(b rune) bool {
		return b == '"'
	})
	t.Typ = TokenString
	return t
}

func (l *Lexer) emitAlphaNum() Token {
	t := l.emitUntil(func(b rune) bool {
		return unicode.IsSpace(b)
	})
	if typ, ok := keywords[string(t.Lit)]; ok {
		t.Typ = typ
	} else {
		t.Typ = TokenText
	}
	return t
}

func (l *Lexer) emitUntil(stop func(rune) bool) Token {
	var lit []rune
	for {
		r, err := l.read()
		if err != nil {
			if err == io.EOF {
				return l.err(errors.New("unexpected eof"))
			}
			return l.err(err)
		}
		if stop(r) {
			break
		}
		lit = append(lit, r)
	}
	return Token{Lit: lit}
}
