package parser

import (
	"errors"
	"strings"
	"testing"
)

func tokensEqual(t1, t2 Token) bool {
	return t1.String() == t2.String()
}

func TestLexer(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output []Token
	}{
		{
			name: "normal",
			input: `
import foo;

func main() {
  print("hello world");
}
`,
			output: []Token{
				Token{
					Typ:     TokenImport,
					Lit:     []rune("import"),
					Pos:     1,
					Line:    1,
					LinePos: 0,
				},
				Token{
					Typ:     TokenText,
					Lit:     []rune("foo"),
					Pos:     8,
					Line:    1,
					LinePos: 7,
				},
				Token{
					Typ:     TokenSemicolon,
					Lit:     []rune(";"),
					Pos:     11,
					Line:    1,
					LinePos: 10,
				},
				Token{
					Typ:     TokenFunc,
					Lit:     []rune("func"),
					Pos:     14,
					Line:    3,
					LinePos: 0,
				},
				Token{
					Typ:     TokenText,
					Lit:     []rune("main"),
					Pos:     19,
					Line:    3,
					LinePos: 5,
				},
				Token{
					Typ:     TokenParensOpen,
					Lit:     []rune("("),
					Pos:     23,
					Line:    3,
					LinePos: 9,
				},
				Token{
					Typ:     TokenParensClose,
					Lit:     []rune(")"),
					Pos:     24,
					Line:    3,
					LinePos: 10,
				},
				Token{
					Typ:     TokenBraceOpen,
					Lit:     []rune("{"),
					Pos:     26,
					Line:    3,
					LinePos: 12,
				},
				Token{
					Typ:     TokenText,
					Lit:     []rune("print"),
					Pos:     30,
					Line:    4,
					LinePos: 2,
				},
				Token{
					Typ:     TokenParensOpen,
					Lit:     []rune("("),
					Pos:     35,
					Line:    4,
					LinePos: 7,
				},
				Token{
					Typ:     TokenString,
					Lit:     []rune("hello world"),
					Pos:     37,
					Line:    4,
					LinePos: 8,
				},
				Token{
					Typ:     TokenParensClose,
					Lit:     []rune(")"),
					Pos:     49,
					Line:    4,
					LinePos: 21,
				},
				Token{
					Typ:     TokenSemicolon,
					Lit:     []rune(";"),
					Pos:     50,
					Line:    4,
					LinePos: 22,
				},
				Token{
					Typ:     TokenBraceClose,
					Lit:     []rune("}"),
					Pos:     52,
					Line:    5,
					LinePos: 0,
				},
			},
		},
		{
			name:  "unterminated_string",
			input: "\"foo",
			output: []Token{
				{Typ: TokenError, Pos: 0, Err: errors.New("unexpected eof")},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLexer("test.apl", strings.NewReader(tc.input))
			var tokens []Token
			for token := range l.Tokens() {
				tokens = append(tokens, token)
			}
			if len(tokens) != len(tc.output) {
				t.Fatalf("expected %d tokens. got %d", len(tc.output), len(tokens))
			}
			for i := range tokens {
				if !tokensEqual(tokens[i], tc.output[i]) {
					t.Errorf("tokens[%d] expected %v, but got %v", i, tc.output[i], tokens[i])
				}
				start := tokens[i].Pos
				end := start + len(tokens[i].Lit)
				if tc.input[start:end] != string(tokens[i].Lit) {
					t.Errorf("token offset wrong. expected %q, but got %q", string(tokens[i].Lit), tc.input[start:end])
				}
			}
		})
	}
}
