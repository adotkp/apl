package parser

import (
	"strings"
	"testing"

	"ast"
	"ast/expr"
	"ast/statement"
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
				Imports: []*statement.Import{
					{Name: "foo"},
					{Name: "bar"},
				},
				Decls: []ast.Decl{
					&ast.FnDecl{
						Nam: "hello",
						Args: []*ast.FnArg{
							{Typ: "int", Nam: "i"},
							{Typ: "string", Nam: "x"},
						},
						Return: &ast.FnReturn{Typ: "bool"},
						Statements: []statement.Statement{
							&statement.FnCall{
								Nam:    "do",
								Params: nil,
							},
							&statement.Return{
								Expr: &expr.Value{
									Text: "false",
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
			l := NewLexer(strings.NewReader(tc.input))
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
					t.Errorf("expected %q, got %q", tc.output, file)
				}
			}
		})
	}
}
