package parser

import (
	"strings"
	"testing"

	"ast"
	"ast/expr"
	"ast/statement"
	"values"
)

func fileEqual(f1, f2 *ast.File) bool {
	return f1.String() == f2.String()
}

func TestParser(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output *ast.File
		err    string
	}{
		{
			name: "normal",
			input: `
import foo;
import bar;

func hello(int i, string x) bool {
  do();
  return false;
}
`,
			output: &ast.File{
				Source: TokenSource{
					Token{Line: 1, LinePos: 0, Pos: 1, File: "test.apl"},
				},
				Imports: []*statement.Import{
					{
						Name: "foo",
						Source: TokenSource{
							Token{Line: 1, LinePos: 0, Pos: 1, File: "test.apl"},
						},
					},
					{
						Name: "bar",
						Source: TokenSource{
							Token{Line: 2, LinePos: 0, Pos: 13, File: "test.apl"},
						},
					},
				},
				Decls: []ast.Decl{
					&ast.FnDecl{
						Nam: "hello",
						Source: TokenSource{
							Token{Line: 4, LinePos: 0, Pos: 26, File: "test.apl"},
						},
						Args: []*ast.FnArg{
							{
								Typ: "int",
								Nam: "i",
								Source: TokenSource{
									Token{Line: 4, LinePos: 11, Pos: 37, File: "test.apl"},
								},
							},
							{
								Typ: "string",
								Nam: "x",
								Source: TokenSource{
									Token{Line: 4, LinePos: 18, Pos: 44, File: "test.apl"},
								},
							},
						},
						Return: &ast.FnReturn{
							Typ: "bool",
							Source: TokenSource{
								Token{Line: 4, LinePos: 28, Pos: 54, File: "test.apl"},
							},
						},
						Statements: []statement.Statement{
							&statement.FnCall{
								Nam:    "do",
								Params: nil,
								Source: TokenSource{
									Token{Line: 5, LinePos: 2, Pos: 63, File: "test.apl"},
								},
							},
							&statement.Return{
								Source: TokenSource{
									Token{Line: 6, LinePos: 2, Pos: 71, File: "test.apl"},
								},
								Expr: &expr.Value{
									V: &values.Bool{V: false},
									Source: TokenSource{
										Token{Line: 6, LinePos: 9, Pos: 78, File: "test.apl"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name:   "lex_error",
			input:  "\"foo",
			output: nil,
			err:    "unexpected eof",
		},
		{
			name:   "import_missing_name",
			input:  "import ;",
			output: nil,
			err:    "error at pos 7 (;): expected TokenText, got TokenSemicolon",
		},
		{
			name:   "import_missing_semicolon",
			input:  "import foo import bar;",
			output: nil,
			err:    "error at pos 11 (import): did not expect TokenImport",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := NewLexer("test.apl", strings.NewReader(tc.input))
			p := NewParser(l.Tokens())
			file, err := p.Do()
			if err != nil {
				if tc.err == "" {
					t.Fatal(err)
				}
				if err.Error() != tc.err {
					t.Fatalf("expected error with %q, got %q", tc.err, err.Error())
				}
			} else {
				if !fileEqual(tc.output, file) {
					t.Errorf("expected\n%q\n\ngot\n%q", tc.output, file)
				}
			}
		})
	}
}
