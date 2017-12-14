package runtime

import (
	"testing"
)

func TestCheck(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		err   string
	}{
		{
			name: "normal",
			input: `
func main(int x) {
  main(12345);
}
`,
			err: "",
		},
		{
			name: "function_declare_param_unknown_type",
			input: `
func main(foo x) {
  main("hello world");
}
`,
			err: "test:2:11 unknown type: foo",
		},
		{
			name: "function_declare_return_unknown_type",
			input: `
func main(int x) foo {
  main("hello world");
}
`,
			err: "test:2:18 unknown type: foo",
		},
		{
			name: "function_declare_conflict",
			input: `
func main(int x) {
}
func main(int x) {
}
`,
			err: "test:4:1 type main already declared as type<func>",
		},
		{
			name: "function_call_param_type_mismatch",
			input: `
func main(int x) {
  main("hello world");
}
`,
			err: "test:3:3 main param #1 expects type<int>, not type<string>",
		},
		{
			name: "function_call_param_count_mismatch",
			input: `
func main(int x, int y) {
  main(1);
}
`,
			err: "test:3:3 main expects 2 params, not 1",
		},
		{
			name: "function_call_unknown_name",
			input: `
func main() {
  foo();
}
`,
			err: "test:3:3 unknown type: foo",
		},
		{
			name: "function_call_wrong_type",
			input: `
func main() {
  int();
}
`,
			err: "test:3:3 int is type<int>, not func",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loader := &StringLoader{
				m: map[string]string{"test": tc.input},
			}
			e := NewExecutor(loader)
			err := e.Check("test")
			if tc.err == "" && err != nil {
				t.Errorf("unexpected error: %s", err)
			} else if tc.err != "" && err.Error() != tc.err {
				t.Errorf("expected %q but got %q", tc.err, err.Error())
			}
		})
	}
}

func TestImport(t *testing.T) {
	testCases := []struct {
		name  string
		input map[string]string
		err   string
	}{
		{
			name: "normal",
			input: map[string]string{
				"test": `
import foo;      
func main(int x) {
  lib(true);
}
`, "foo": `func lib(bool b) {}`,
			},
			err: "",
		},
		{
			name: "unknown_import",
			input: map[string]string{
				"test": `
import foo;      
func main(int x) {
  lib(true);
}`},
			err: "unknown import: foo",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			loader := &StringLoader{
				m: tc.input,
			}
			e := NewExecutor(loader)
			err := e.Check("test")
			if tc.err == "" && err != nil {
				t.Errorf("unexpected error: %s", err)
			} else if tc.err != "" && err.Error() != tc.err {
				t.Errorf("expected %q but got %q", tc.err, err.Error())
			}
		})
	}
}
