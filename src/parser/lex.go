package parser

import (
	"errors"
	"fmt"
	"io"
	"unicode"
)

type TokenType int64

const (
	TokenError TokenType = 1 << iota
	TokenBraceOpen
	TokenBraceClose
	TokenParensOpen
	TokenParensClose
	TokenComma
	TokenSemicolon
	TokenAssign
	TokenString
	TokenNumber
	TokenIf
	TokenElse
	TokenFunc
	TokenTyp
	TokenImport
	TokenReturn
	TokenText
)

func (t TokenType) String() string {
	return tokens[t]
}

var (
	keywords map[string]TokenType
	tokens   map[TokenType]string
)

func init() {
	keywords = map[string]TokenType{
		"if":     TokenIf,
		"else":   TokenElse,
		"func":   TokenFunc,
		"type":   TokenTyp,
		"import": TokenImport,
		"return": TokenReturn,
	}
	tokens = map[TokenType]string{
		TokenError:       "TokenError",
		TokenBraceOpen:   "TokenBraceOpen",
		TokenBraceClose:  "TokenBraceClose",
		TokenParensOpen:  "TokenParensOpen",
		TokenParensClose: "TokenParensClose",
		TokenComma:       "TokenComma",
		TokenSemicolon:   "TokenSemicolon",
		TokenAssign:      "TokenAssign",
		TokenString:      "TokenString",
		TokenNumber:      "TokenNumber",
		TokenIf:          "TokenIf",
		TokenElse:        "TokenElse",
		TokenFunc:        "TokenFunc",
		TokenTyp:         "TokenType",
		TokenImport:      "TokenImport",
		TokenReturn:      "TokenReturn",
		TokenText:        "TokenText",
	}
}

type Token struct {
	Typ       TokenType
	Lit       []rune
	Pos       int
	Err       error
	IsKeyword bool
}

func (t Token) String() string {
	if t.Typ == TokenError {
		return fmt.Sprintf("TokenError(%s)", t.Err)
	}
	return fmt.Sprintf("Token(%v,%d,%s,@%d)", t.Typ, TokenError, string(t.Lit), t.Pos)
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
			token := l.nextIgnoreSpace()
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

func (l *Lexer) nextIgnoreSpace() Token {
	err := l.consumeWhitespace()
	if err != nil {
		return l.err(err)
	}
	pos := l.pos
	t := l.next()
	t.Pos = pos
	if t.Typ == TokenString {
		t.Pos++
	}
	return t
}

func (l *Lexer) next() Token {
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
	case '=':
		return l.emitSymbol(r, TokenAssign)
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
	if t.Err != nil {
		return t
	}
	t.Typ = TokenString
	return t
}

func (l *Lexer) emitAlphaNum() Token {
	t := l.emitUntil(func(b rune) bool {
		isAlphaNum := ('0' <= b && b <= '9') ||
			('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			(b == '_')
		return !isAlphaNum
	})
	if t.Err != nil {
		return t
	}
	err := l.unread()
	if err != nil {
		return l.err(err)
	}
	if typ, ok := keywords[string(t.Lit)]; ok {
		t.Typ = typ
		t.IsKeyword = true
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
