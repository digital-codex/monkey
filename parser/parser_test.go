package parser

import (
	"fmt"
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/lexer"
	"reflect"
	"strconv"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			ident string
			value any
		}
	}{
		{
			input: `let x = 5;`,
			expected: struct {
				ident string
				value any
			}{
				"x",
				5,
			},
		},
		{
			input: `let y = true;`,
			expected: struct {
				ident string
				value any
			}{
				"y",
				true,
			},
		},
		{
			input: `let foobar = 838383;`,
			expected: struct {
				ident string
				value any
			}{
				"foobar",
				838383,
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.ident, test.expected.value)
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`return 5;`, 5},
		{`return 10;`, 10},
		{`return 993322;`, 993322},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected)
	}
}

func TestIdentifier(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			value   string
		}
	}{
		{
			input: `foobar;`,
			expected: struct {
				literal string
				value   string
			}{
				"foobar",
				"foobar",
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.value)
	}
}

func TestIntegerLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			value   int
		}
	}{
		{
			input: `5;`,
			expected: struct {
				literal string
				value   int
			}{
				"5",
				5,
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.value)
	}
}

func TestPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal  string
			operator string
			right    any
		}
	}{
		{
			input: `!5;`,
			expected: struct {
				literal  string
				operator string
				right    any
			}{
				"!",
				"!", 5,
			},
		},
		{
			input: `-15;`,
			expected: struct {
				literal  string
				operator string
				right    any
			}{
				"-",
				"-", 15,
			},
		},
		{
			input: `!true;`,
			expected: struct {
				literal  string
				operator string
				right    any
			}{
				"!",
				"!", true,
			},
		},
		{
			input: `!false;`,
			expected: struct {
				literal  string
				operator string
				right    any
			}{
				"!",
				"!", false,
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.operator, test.expected.right)
	}
}

func TestInfixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal  string
			left     any
			operator string
			right    any
		}
	}{
		{
			input: `5 + 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "+", 5,
			},
		},
		{
			input: `5 - 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "-", 5,
			},
		},
		{
			input: `5 * 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "*", 5,
			},
		},
		{
			input: `5 / 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "/", 5,
			},
		},
		{
			input: `5 > 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, ">", 5,
			},
		},
		{
			input: `5 < 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "<", 5,
			},
		},
		{
			input: `5 == 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "==", 5,
			},
		},
		{
			input: `5 != 5;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"5",
				5, "!=", 5,
			},
		},
		{
			input: `true == true;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"true",
				true, "==", true,
			},
		},
		{
			input: `true != false;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"true",
				true, "!=", false,
			},
		},
		{
			input: `false == false;`,
			expected: struct {
				literal  string
				left     any
				operator string
				right    any
			}{
				"false",
				false, "==", false,
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.left, test.expected.operator, test.expected.right)
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			value   bool
		}
	}{
		{
			input: `true;`,
			expected: struct {
				literal string
				value   bool
			}{
				"true",
				true,
			},
		},
		{
			input: `false;`,
			expected: struct {
				literal string
				value   bool
			}{
				"false",
				false,
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.value)
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal     string
			left        any
			operator    string
			right       any
			consequence struct {
				literal string
				value   any
			}
			alternative *struct {
				literal string
				value   any
			}
		}
	}{
		{
			input: `if (x < y) { x }`,
			expected: struct {
				literal     string
				left        any
				operator    string
				right       any
				consequence struct {
					literal string
					value   any
				}
				alternative *struct {
					literal string
					value   any
				}
			}{
				literal: "if",
				left:    "x", operator: "<", right: "y",
				consequence: struct {
					literal string
					value   any
				}{
					"x",
					"x",
				},
				alternative: nil,
			},
		},
		{
			input: `if (x < y) { x } else { y }`,
			expected: struct {
				literal     string
				left        any
				operator    string
				right       any
				consequence struct {
					literal string
					value   any
				}
				alternative *struct {
					literal string
					value   any
				}
			}{
				literal: "if",
				left:    "x", operator: "<", right: "y",
				consequence: struct {
					literal string
					value   any
				}{
					"x",
					"x",
				},
				alternative: &struct {
					literal string
					value   any
				}{
					"y",
					"y",
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.left, test.expected.operator, test.expected.right, test.expected.consequence, test.expected.alternative)
	}
}

func TestFunctionLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal    string
			parameters []string
			body       struct {
				literal  string
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `fn(x, y) { x + y; }`,
			expected: struct {
				literal    string
				parameters []string
				body       struct {
					literal  string
					left     any
					operator string
					right    any
				}
			}{
				literal:    "fn",
				parameters: []string{"x", "y"},
				body: struct {
					literal  string
					left     any
					operator string
					right    any
				}{
					"x",
					"x", "+", "y",
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.parameters, test.expected.body)
	}
}

func TestCallExpressionParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal   string
			ident     string
			arguments []struct {
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `add(1, 2 * 3, 4 + 5);`,
			expected: struct {
				literal   string
				ident     string
				arguments []struct {
					left     any
					operator string
					right    any
				}
			}{
				"add",
				"add",
				[]struct {
					left     any
					operator string
					right    any
				}{
					{nil, "", 1},
					{2, "*", 3},
					{4, "+", 5},
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.ident, test.expected.arguments)
	}
}

func TestStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			value   string
		}
	}{
		{
			input: `"hello, world";`,
			expected: struct {
				literal string
				value   string
			}{
				"hello, world",
				"hello, world",
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.value)
	}
}

func TestArrayLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal   string
			arguments []struct {
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `[1, 2 * 2, 3 + 3];`,
			expected: struct {
				literal   string
				arguments []struct {
					left     any
					operator string
					right    any
				}
			}{
				"[",
				[]struct {
					left     any
					operator string
					right    any
				}{
					{nil, "", 1},
					{2, "*", 2},
					{3, "+", 3},
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.arguments)
	}
}

func TestIndexExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			left    string
			index   struct {
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `array[1 + 1];`,
			expected: struct {
				literal string
				left    string
				index   struct {
					left     any
					operator string
					right    any
				}
			}{
				literal: "array",
				left:    "array",
				index: struct {
					left     any
					operator string
					right    any
				}{
					1, "+", 1,
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.left, test.expected.index)
	}
}

func TestHashLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			pairs   map[string]struct {
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `{"one": 1, "two": 2, "three": 3};`,
			expected: struct {
				literal string
				pairs   map[string]struct {
					left     any
					operator string
					right    any
				}
			}{
				literal: "{",
				pairs: map[string]struct {
					left     any
					operator string
					right    any
				}{
					"one":   {nil, "", 1},
					"two":   {nil, "", 2},
					"three": {nil, "", 3},
				},
			},
		},
		{
			input: `{};`,
			expected: struct {
				literal string
				pairs   map[string]struct {
					left     any
					operator string
					right    any
				}
			}{
				literal: "{",
				pairs: map[string]struct {
					left     any
					operator string
					right    any
				}{},
			},
		},
		{
			input: `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5};`,
			expected: struct {
				literal string
				pairs   map[string]struct {
					left     any
					operator string
					right    any
				}
			}{
				literal: "{",
				pairs: map[string]struct {
					left     any
					operator string
					right    any
				}{
					"one":   {0, "+", 1},
					"two":   {10, "-", 8},
					"three": {15, "/", 5},
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.pairs)
	}
}

func TestMacroLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal    string
			parameters []string
			body       struct {
				literal  string
				left     any
				operator string
				right    any
			}
		}
	}{
		{
			input: `macro(x, y) { x + y; }`,
			expected: struct {
				literal    string
				parameters []string
				body       struct {
					literal  string
					left     any
					operator string
					right    any
				}
			}{
				literal:    "macro",
				parameters: []string{"x", "y"},
				body: struct {
					literal  string
					left     any
					operator string
					right    any
				}{
					"x",
					"x", "+", "y",
				},
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.parameters, test.expected.body)
	}

}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
		{"a + add(b * c) + d", "((a + add((b * c))) + d)"},
		{"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))", "add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))"},
		{"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))"},
		{"a * [1, 2, 3, 4][b * c] * d", "((a * ([1, 2, 3, 4][(b * c)])) * d)"},
		{"add(a * b[2], b[1], 2 * [1, 2][1])", "add((a * (b[2])), (b[1]), (2 * ([1, 2][1])))"},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		assertions.AssertStringEquals(t, test.expected, program.String(), "test["+strconv.Itoa(i)+"] - program.String() wrong")
	}
}

func testProgram(t *testing.T, i int, program *ast.Program) {
	assertions.AssertNotNull(t, program, "test["+strconv.Itoa(i)+"] - ParseProgram() returned nil")
	assertions.AssertIntEquals(t, 1, len(program.Statements), "test["+strconv.Itoa(i)+"] - program.Statements wrong")
}

func testStatement(stmt ast.Statement) func(*testing.T, int, ast.Statement, ...any) {
	switch stmt.(type) {
	case *ast.LetStatement:
		return testLetStatement
	case *ast.ReturnStatement:
		return testReturnStatement
	case *ast.ExpressionStatement:
		return testExpressionStatement
	case *ast.BlockStatement:
		return testBlockStatement
	default:
		panic(fmt.Sprintf("unsupported statement type: %T", stmt))
	}
}

func testExpression(exp ast.Expression) func(*testing.T, int, ast.Expression, ...any) {
	switch exp.(type) {
	case *ast.Identifier:
		return testIdentifier
	case *ast.IntegerLiteral:
		return testIntegerLiteral
	case *ast.PrefixExpression:
		return testPrefixExpression
	case *ast.InfixExpression:
		return testInfixExpression
	case *ast.Boolean:
		return testBoolean
	case *ast.IfExpression:
		return testIfExpression
	case *ast.FunctionLiteral:
		return testFunctionLiteral
	case *ast.CallExpression:
		return testCallExpression
	case *ast.StringLiteral:
		return testStringLiteral
	case *ast.ArrayLiteral:
		return testArrayLiteral
	case *ast.IndexExpression:
		return testIndexExpression
	case *ast.HashLiteral:
		return testHashLiteral
	case *ast.MacroLiteral:
		return testMacroLiteral
	default:
		panic(fmt.Sprintf("unsupported expression type: %T", exp))
	}
}

func testLetStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.LetStatement{}), stmt, "test["+strconv.Itoa(i)+"] - ast.Statement unexpected type")
	assertions.AssertStringEquals(t, "let", stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - ast.Statement.TokenLiteral() wrong")
	if 1 >= len(expected) {
		t.Fatalf("testLetStatement: len(expected) wrong: expected=>1, actual=%d", len(expected))
	}
	testIdentifier(t, i, stmt.(*ast.LetStatement).Name, expected[0])
	testExpression(stmt.(*ast.LetStatement).Value)(t, i, stmt.(*ast.LetStatement).Value, expected[1:]...)
}

func testReturnStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.ReturnStatement{}), stmt, "test["+strconv.Itoa(i)+"] - ast.Statement unexpected type")
	assertions.AssertStringEquals(t, "return", stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - ast.Statement.TokenLiteral() wrong")
	if 0 == len(expected) {
		t.Fatalf("testReturnStatement: len(expected) wrong: expected=>0, actual=%d", len(expected))
	}
	testExpression(stmt.(*ast.ReturnStatement).ReturnValue)(t, i, stmt.(*ast.ReturnStatement).ReturnValue, expected...)
}

func testExpressionStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.ExpressionStatement{}), stmt, "test["+strconv.Itoa(i)+"] - stmt unexpected type")
	if 1 >= len(expected) {
		t.Fatalf("testExpressionStatement: len(expected) wrong: expected=>1, actual=%d", len(expected))
	}
	literal, ok := expected[0].(string)
	if !ok {
		t.Fatalf("testExpressionStatement: expected[0] unexpected type: expected=string, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, literal, stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - stmt.TokenLiteral() wrong")
	testExpression(stmt.(*ast.ExpressionStatement).Expression)(t, i, stmt.(*ast.ExpressionStatement).Expression, expected[1:]...)
}

func testBlockStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertNotNull(t, stmt, "test["+strconv.Itoa(i)+"] - stmt is nil")
	assertions.AssertStringEquals(t, "{", stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - stmt.TokenLiteral() wrong")
	if 0 == len(expected) {
		t.Fatalf("testBlockStatement: len(expected) wrong: expected=>0, actual=%d", len(expected))
	}
	assertions.AssertIntEquals(t, 1, len((stmt.(*ast.BlockStatement)).Statements), "test["+strconv.Itoa(i)+"] - block.Statements wrong")
	testStatement(stmt.(*ast.BlockStatement).Statements[0])(t, i, stmt.(*ast.BlockStatement).Statements[0], expected...)
}

func testIdentifier(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(&ast.Identifier{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testIdentifier: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	value, ok := expected[0].(string)
	if !ok {
		t.Fatalf("testIdentifier: expected[0] unexpected type: expected=string, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, value, exp.(*ast.Identifier).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.Identifier).TokenLiteral() wrong")
	assertions.AssertStringEquals(t, value, exp.(*ast.Identifier).Value, "test["+strconv.Itoa(i)+"] exp.(*ast.Identifier).Value wrong")
}

func testIntegerLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.IntegerLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testIntegerLiteral: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	value, ok := expected[0].(int)
	if !ok {
		t.Fatalf("testIntegerLiteral: expected[0] unexpected type: expected=int64, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, fmt.Sprintf("%d", value), exp.(*ast.IntegerLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.IntegerLiteral).TokenLiteral() wrong")
	assertions.AssertInt64Equals(t, int64(value), exp.(*ast.IntegerLiteral).Value, "test["+strconv.Itoa(i)+"] exp.(*ast.IntegerLiteral).Value wrong")
}

func testPrefixExpression(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.PrefixExpression{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 2 != len(expected) {
		t.Fatalf("testPrefixExpression: len(expected) wrong: expected=2, actual=%d", len(expected))
	}
	operator, ok := expected[0].(string)
	if !ok {
		t.Fatalf("testPrefixExpression: expected[0] unexpected type: expected=string, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, operator, exp.(*ast.PrefixExpression).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.PrefixExpression).TokenLiteral() wrong")
	assertions.AssertStringEquals(t, operator, exp.(*ast.PrefixExpression).Operator, "test["+strconv.Itoa(i)+"] exp.(*ast.PrefixExpression).Operator wrong")
	testExpression(exp.(*ast.PrefixExpression).Right)(t, i, exp.(*ast.PrefixExpression).Right, expected[1])
}

func testInfixExpression(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.InfixExpression{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 3 != len(expected) {
		t.Fatalf("testInfixExpression: len(expected) wrong: expected=3, actual=%d", len(expected))
	}
	operator, ok := expected[1].(string)
	if !ok {
		t.Fatalf("testInfixExpression: expected[1] unexpected type: expected=string, actual=%T", expected[1])
	}
	assertions.AssertStringEquals(t, operator, exp.(*ast.InfixExpression).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.InfixExpression).TokenLiteral() wrong")
	assertions.AssertStringEquals(t, operator, exp.(*ast.InfixExpression).Operator, "test["+strconv.Itoa(i)+"] exp.(*ast.InfixExpression).Operator wrong")
	testExpression(exp.(*ast.InfixExpression).Left)(t, i, exp.(*ast.InfixExpression).Left, expected[0])
	testExpression(exp.(*ast.InfixExpression).Right)(t, i, exp.(*ast.InfixExpression).Right, expected[2])
}

func testBoolean(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.Boolean{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testBoolean: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	value, ok := expected[0].(bool)
	if !ok {
		t.Fatalf("testBoolean: expected[0] unexpected type: expected=bool, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, fmt.Sprintf("%t", value), exp.(*ast.Boolean).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.Boolean).TokenLiteral() wrong")
	assertions.AssertBoolEquals(t, value, exp.(*ast.Boolean).Value, "test["+strconv.Itoa(i)+"] exp.(*ast.Boolean).Value wrong")
}

func testIfExpression(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.IfExpression{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	assertions.AssertStringEquals(t, "if", exp.(*ast.IfExpression).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.IfExpression).TokenLiteral() wrong")
	if 5 != len(expected) {
		t.Fatalf("testIfExpression: len(expected) wrong: expected=5, actual=%d", len(expected))
	}
	consequence, ok := expected[3].(struct {
		literal string
		value   any
	})
	if !ok {
		t.Fatalf("testIfExpression: expected[3] unexpected type: expected=struct{literal string; value any}, actual=%T", expected[3])
	}
	testExpression(exp.(*ast.IfExpression).Condition)(t, i, exp.(*ast.IfExpression).Condition, expected[0:3]...)
	testBlockStatement(t, i, exp.(*ast.IfExpression).Consequence, consequence.literal, consequence.value)
	if exp.(*ast.IfExpression).Alternative != nil {
		alternative, ok := expected[4].(*struct {
			literal string
			value   any
		})
		if !ok {
			t.Fatalf("testIfExpression: expected[4] unexpected type: expected=struct{literal string; value any}, actual=%T", expected[4])
		}
		testBlockStatement(t, i, exp.(*ast.IfExpression).Alternative, alternative.literal, alternative.value)
	}
}

func testFunctionLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.FunctionLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	assertions.AssertStringEquals(t, "fn", exp.(*ast.FunctionLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.FunctionLiteral).TokenLiteral() wrong")
	if 2 != len(expected) {
		t.Fatalf("testFunctionLiteral: len(expected) wrong: expected=2, actual=%d", len(expected))
	}
	parameters, ok := expected[0].([]string)
	if !ok {
		t.Fatalf("testFunctionLiteral: expected[0] unexpected type: expected=[]string, actual=%T", expected[0])
	}
	for n, param := range parameters {
		testIdentifier(t, i, exp.(*ast.FunctionLiteral).Parameters[n], param)
	}
	body, ok := expected[1].(struct {
		literal  string
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testFunctionLiteral: expected[1] unexpected type: expected=[]struct{literal string; left any; operator string; right any}, actual=%T", expected[1])
	}
	testBlockStatement(t, i, exp.(*ast.FunctionLiteral).Body, body.literal, body.left, body.operator, body.right)
}

func testCallExpression(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.CallExpression{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	assertions.AssertStringEquals(t, "(", exp.(*ast.CallExpression).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.CallExpression).TokenLiteral() wrong")
	if 2 != len(expected) {
		t.Fatalf("testCallExpression: len(expected) wrong: expected=2, actual=%d", len(expected))
	}
	arguments, ok := expected[1].([]struct {
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testFunctionLiteral: expected[1] unexpected type: expected=[]struct{left any; operator string; right any}, actual=%T", expected[1])
	}
	testExpression(exp.(*ast.CallExpression).Function)(t, i, exp.(*ast.CallExpression).Function, expected[0])
	for n, arg := range arguments {
		switch {
		case arg.left == nil && arg.operator == "":
			testExpression(exp.(*ast.CallExpression).Argument[n])(t, i, exp.(*ast.CallExpression).Argument[n], arg.right)
		case arg.left == nil && arg.operator != "":
			testExpression(exp.(*ast.CallExpression).Argument[n])(t, i, exp.(*ast.CallExpression).Argument[n], arg.operator, arg.right)
		default:
			testExpression(exp.(*ast.CallExpression).Argument[n])(t, i, exp.(*ast.CallExpression).Argument[n], arg.left, arg.operator, arg.right)
		}
	}
}

func testStringLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.StringLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testStringLiteral: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	value, ok := expected[0].(string)
	if !ok {
		t.Fatalf("testStringLiteral: expected[0] unexpected type: expected=string, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, value, exp.(*ast.StringLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.StringLiteral).TokenLiteral() wrong")
	assertions.AssertStringEquals(t, value, exp.(*ast.StringLiteral).Value, "test["+strconv.Itoa(i)+"] exp.(*ast.StringLiteral).Value wrong")
}

func testArrayLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.ArrayLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testArrayLiteral: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	arguments, ok := expected[0].([]struct {
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testArrayLiteral: expected[0] unexpected type: expected=[]struct{left any; operator string; right any}, actual=%T", expected[0])
	}
	for n, arg := range arguments {
		switch {
		case arg.left == nil && arg.operator == "":
			testExpression(exp.(*ast.ArrayLiteral).Elements[n])(t, i, exp.(*ast.ArrayLiteral).Elements[n], arg.right)
		case arg.left == nil && arg.operator != "":
			testExpression(exp.(*ast.ArrayLiteral).Elements[n])(t, i, exp.(*ast.ArrayLiteral).Elements[n], arg.operator, arg.right)
		default:
			testExpression(exp.(*ast.ArrayLiteral).Elements[n])(t, i, exp.(*ast.ArrayLiteral).Elements[n], arg.left, arg.operator, arg.right)
		}
	}
}

func testIndexExpression(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.IndexExpression{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 2 != len(expected) {
		t.Fatalf("testIndexExpression: len(expected) wrong: expected=2, actual=%d", len(expected))
	}
	index, ok := expected[1].(struct {
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testIndexExpression: expected[1] unexpected type: expected=struct{left any; operator string; right any}, actual=%T", expected[1])
	}
	testExpression(exp.(*ast.IndexExpression).Left)(t, i, exp.(*ast.IndexExpression).Left, expected[0])
	testExpression(exp.(*ast.IndexExpression).Index)(t, i, exp.(*ast.IndexExpression).Index, index.left, index.operator, index.right)
}

func testHashLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.HashLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	if 1 != len(expected) {
		t.Fatalf("testHashLiteral: len(expected) wrong: expected=1, actual=%d", len(expected))
	}
	pairs, ok := expected[0].(map[string]struct {
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testHashLiteral: expected[0] unexpected type: expected=map[string]struct{left any; operator string; right any}, actual=%T", expected[0])
	}
	assertions.AssertIntEquals(t, len(pairs), len(exp.(*ast.HashLiteral).Pairs), "test["+strconv.Itoa(i)+"] - len(ast.(*ast.HashLiteral).Pairs) wrong")
	for key, val := range exp.(*ast.HashLiteral).Pairs {
		assertions.AssertTypeOf(t, reflect.TypeOf(&ast.StringLiteral{}), key, "test["+strconv.Itoa(i)+"] - key unexpected type")

		expect := pairs[key.String()]
		switch {
		case expect.left == nil && expect.operator == "":
			testExpression(val)(t, i, val, expect.right)
		case expect.left == nil && expect.operator != "":
			testExpression(val)(t, i, val, expect.operator, expect.right)
		default:
			testExpression(val)(t, i, val, expect.left, expect.operator, expect.right)
		}
	}
}

func testMacroLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.MacroLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	assertions.AssertStringEquals(t, "macro", exp.(*ast.MacroLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.MacroLiteral).TokenLiteral() wrong")
	if 2 != len(expected) {
		t.Fatalf("testMacroLiteral: len(expected) wrong: expected=2, actual=%d", len(expected))
	}
	parameters, ok := expected[0].([]string)
	if !ok {
		t.Fatalf("testMacroLiteral: expected[0] unexpected type: expected=[]string, actual=%T", expected[0])
	}
	for n, param := range parameters {
		testIdentifier(t, i, exp.(*ast.MacroLiteral).Parameters[n], param)
	}
	body, ok := expected[1].(struct {
		literal  string
		left     any
		operator string
		right    any
	})
	if !ok {
		t.Fatalf("testMacroLiteral: expected[1] unexpected type: expected=[]struct{literal string; left any; operator string; right any}, actual=%T", expected[1])
	}
	testBlockStatement(t, i, exp.(*ast.MacroLiteral).Body, body.literal, body.left, body.operator, body.right)
}

func checkParserErrors(t *testing.T, p *Parser) {
	if p.hadError {
		errors := p.Errors()
		t.Errorf("parser has %d errors", len(errors))
		for _, msg := range errors {
			t.Errorf("parser error: %q", msg)
		}
		t.FailNow()
	}
}
