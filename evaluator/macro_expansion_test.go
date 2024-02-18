package evaluator

import (
	"fmt"
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/object"
	"github.com/digital-codex/monkey/parser"
	"reflect"
	"strconv"
	"testing"
)

func TestDefineMacros(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			stmts  int
			env    []string
			ident  string
			params struct {
				count  int
				idents []string
			}
			body string
		}
	}{
		{
			input: `let number = 1; let function = fn(x, y) { x + y }; let add = macro(x, y) { x + y; }`,
			expected: struct {
				stmts  int
				env    []string
				ident  string
				params struct {
					count  int
					idents []string
				}
				body string
			}{
				stmts: 2,
				env: []string{
					"number",
					"function",
				},
				ident: "add",
				params: struct {
					count  int
					idents []string
				}{
					count:  2,
					idents: []string{"x", "y"},
				},
				body: "(x + y)",
			},
		},
	}

	for i, test := range tests {

		program := parse(test.input)
		env := object.NewEnvironment()
		DefineMacros(program, env)

		assertions.AssertIntEquals(t, test.expected.stmts, len(program.Statements), "test["+strconv.Itoa(i)+"] - len(program.Statements) wrong")
		obj, ok := env.Get(test.expected.ident)
		assertions.AssertBoolEquals(t, true, ok, fmt.Sprintf("test[%d] - %s should be defined", i, test.expected.ident))
		for _, ident := range test.expected.env {
			_, ok = env.Get(ident)
			assertions.AssertBoolEquals(t, false, ok, fmt.Sprintf("test[%d] - %s should not be defined", i, ident))
		}
		assertions.AssertTypeOf(t, reflect.TypeOf(object.Macro{}), obj, "test["+strconv.Itoa(i)+"] - unexpected type")

		macro, _ := obj.(*object.Macro)
		assertions.AssertIntEquals(t, test.expected.params.count, len(macro.Parameters), "test["+strconv.Itoa(i)+"] - len(macro.Parameters) wrong")

		for n, ident := range test.expected.params.idents {
			assertions.AssertStringEquals(t, ident, macro.Parameters[n].Value, fmt.Sprintf("test[%d] - macro.Parameters[%d] wrong", i, n))
		}
		assertions.AssertStringEquals(t, test.expected.body, macro.Body.String(), "test["+strconv.Itoa(i)+"] - macro.Body wrong")
	}
}

func TestExpandMacros(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`let print = macro(func, buf) { quote(unquote(func)(unquote(buf))); }; print(puts, "hello world");`, `puts("hello world")`},
		{`let print = macro(condition, buf) { quote(if (unquote(condition)) { puts(unquote(buf)); }); }; print(true, "hello world");`, `if (true) { puts("hello world") }`},
		{`let infix = macro() { quote(1 + 2); }; infix()`, `(1 + 2)`},
		{`let reverse = macro(a, b) { quote(unquote(b) - unquote(a)); }; reverse(2 + 2, 10 - 5)`, `(10 - 5) - (2 + 2)`},
		{`let unless = macro(condition, consequence, alternative) { quote(if (!(unquote(condition))) { unquote(consequence); } else { unquote(alternative); }); }; unless(10 > 5, puts("not greater"), puts("greater"));`, `if (!(10 > 5)) { puts("not greater") } else { puts("greater") }`},
	}

	for i, test := range tests {
		expected := parse(test.expected)

		program := parse(test.input)
		env := object.NewEnvironment()
		DefineMacros(program, env)
		actual := ExpandMacros(program, env)

		assertions.AssertStringEquals(t, expected.String(), actual.String(), "test["+strconv.Itoa(i)+"] - expanded.String() wrong")
	}
}

func parse(input string) *ast.Program {
	p := parser.New(input)
	return p.ParseProgram()
}
