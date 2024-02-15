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

func TestIdentifierExpression(t *testing.T) {
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

func TestIntegerLiteralExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal string
			value   int64
		}
	}{
		{
			input: `5;`,
			expected: struct {
				literal string
				value   int64
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
				"!", true,
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
			input: `true == true`,
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
			input: `true != false`,
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
			input: `false == false`,
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
			consequence any
			alternative any
		}
	}{
		{
			input: `if (x < y) { x }`,
			expected: struct {
				literal     string
				left        any
				operator    string
				right       any
				consequence any
				alternative any
			}{
				"if",
				"x", "<", "y",
				"x",
				nil,
			},
		},
		{
			input: `if (x < y) { x } else { y }`,
			expected: struct {
				literal     string
				left        any
				operator    string
				right       any
				consequence any
				alternative any
			}{
				"if",
				"x", "<", "y",
				"x",
				"y",
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, i, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.left, test.expected.right, test.expected.consequence, test.expected.alternative)
	}
}

func TestFunctionLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected struct {
			literal    string
			parameters []string
			left       any
			operator   string
			right      any
		}
	}{
		{
			input: `fn(x, y) { x + y; }`,
			expected: struct {
				literal    string
				parameters []string
				left       any
				operator   string
				right      any
			}{
				"fn",
				[]string{"x", "y"},
				"x", "<", "y",
			},
		},
	}

	for i, test := range tests {
		l := lexer.New(test.input)
		p := New(l)
		program := p.ParseProgram()

		checkParserErrors(t, p)
		testProgram(t, 0, program)

		testStatement(program.Statements[0])(t, i, program.Statements[0], test.expected.literal, test.expected.parameters, test.expected.left, test.expected.operator, test.expected.right)
	}
}

//	func TestStringLiteralExpression(t *testing.T) {
//		input := `"hello, world";`
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//		testProgram(t, 0, program, 1)
//
//		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//		literal, ok := stmt.Expression.(*ast.StringLiteral)
//		if !ok {
//			t.Fatalf("exp is not *ast.StringLiteral. got=%T", stmt.Expression)
//		}
//
//		if literal.Value != "hello, world" {
//			t.Errorf("literal.Value not %q. got=%q", "hello world", literal.Value)
//		}
//	}
//
//	func TestArrayLiteralExpression(t *testing.T) {
//		input := "[1, 2 * 2, 3 + 3]"
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//		array, ok := stmt.Expression.(*ast.ArrayLiteral)
//		if !ok {
//			t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
//		}
//
//		if len(array.Elements) != 3 {
//			t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
//		}
//
//		testIntegerLiteral(t, array.Elements[0], 1)
//		testInfixExpression(t, array.Elements[1], 2, "*", 2)
//		testInfixExpression(t, array.Elements[2], 3, "+", 3)
//	}
//
//	func TestFunctionParameterParsing(t *testing.T) {
//		tests := []struct {
//			input    string
//			expected []string
//		}{
//			{"fn() {};", []string{}},
//			{"fn(x) {};", []string{"x"}},
//			{"fn(x, y, z) {};", []string{"x", "y", "z"}},
//		}
//
//		for i, tt := range tests {
//			l := lexer.New(tt.input)
//			p := New(l)
//			program := p.ParseProgram()
//			checkParserErrors(t, p)
//			testProgram(t, i, program, 1)
//
//			stmt := program.Statements[0].(*ast.ExpressionStatement)
//			function := stmt.Expression.(*ast.FunctionLiteral)
//			if len(function.Parameters) != len(tt.expected) {
//				t.Errorf("length parameters wrong. want %d, got=%d\n", len(tt.expected), len(function.Parameters))
//			}
//
//			for i, ident := range tt.expected {
//				testLiteralExpression(t, function.Parameters[i], ident)
//			}
//		}
//	}
//
//	func TestCallExpressionParsing(t *testing.T) {
//		input := "add(1, 2 * 3, 4 + 5);"
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//		testProgram(t, 0, program, 1)
//
//		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//		if !ok {
//			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
//		}
//
//		exp, ok := stmt.Expression.(*ast.CallExpression)
//		if !ok {
//			t.Fatalf("stmt.Expresssion is not ast.CallExpresssion. got=%T", stmt.Expression)
//		}
//
//		testIdentifier(t, 0, exp.Function, "add")
//
//		if len(exp.Argument) != 3 {
//			t.Errorf("length arguemnts wrong. want 3, got=%d\n", len(exp.Argument))
//		}
//
//		testLiteralExpression(t, exp.Argument[0], 1)
//		testInfixExpression(t, exp.Argument[1], 2, "*", 3)
//		testInfixExpression(t, exp.Argument[2], 4, "+", 5)
//	}
//
//	func TestParsingIndexExpression(t *testing.T) {
//		input := "myArray[1 + 1]"
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//		index, ok := stmt.Expression.(*ast.IndexExpression)
//		if !ok {
//			t.Fatalf("exp not *ast.IndexExpression. got=%T", stmt.Expression)
//		}
//
//		testIdentifier(t, 0, index.Left, "myArray")
//
//		if !testInfixExpression(t, index.Index, 1, "+", 1) {
//			return
//		}
//	}
//
//	func TestParsingHashLiteralStringKeys(t *testing.T) {
//		input := `{"one": 1, "two": 2, "three": 3}`
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt := program.Statements[0].(*ast.ExpressionStatement)
//		hash, ok := stmt.Expression.(*ast.HashLiteral)
//		if !ok {
//			t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
//		}
//
//		if len(hash.Paris) != 3 {
//			t.Errorf("hash.Paris has wrong length. got=%d", len(hash.Paris))
//		}
//
//		expected := map[string]int64{
//			"one":   1,
//			"two":   2,
//			"three": 3,
//		}
//
//		for key, value := range hash.Paris {
//			literal, ok := key.(*ast.StringLiteral)
//			if !ok {
//				t.Errorf("key is not ast.StringLiteral. got=%T", key)
//			}
//
//			expectedValue := expected[literal.String()]
//			testIntegerLiteral(t, value, expectedValue)
//		}
//	}
//
//	func TestParsingEmptyHashLiteral(t *testing.T) {
//		input := "{}"
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt := program.Statements[0].(*ast.ExpressionStatement)
//		hash, ok := stmt.Expression.(*ast.HashLiteral)
//		if !ok {
//			t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
//		}
//
//		if len(hash.Paris) != 0 {
//			t.Errorf("hash.Paris has wrong length. got=%d", len(hash.Paris))
//		}
//	}
//
//	func TestParsingHashLiteralWithExpressions(t *testing.T) {
//		input := `{"one": 0 + 1, "two": 10 - 8, "three": 15 / 5}`
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		stmt := program.Statements[0].(*ast.ExpressionStatement)
//		hash, ok := stmt.Expression.(*ast.HashLiteral)
//		if !ok {
//			t.Fatalf("exp is not ast.HashLiteral. got=%T", stmt.Expression)
//		}
//
//		if len(hash.Paris) != 3 {
//			t.Errorf("hash.Paris has wrong length. got=%d", len(hash.Paris))
//		}
//
//		tests := map[string]func(exp ast.Expression){
//			"one": func(exp ast.Expression) {
//				testInfixExpression(t, exp, 0, "+", 1)
//			},
//			"two": func(exp ast.Expression) {
//				testInfixExpression(t, exp, 10, "-", 8)
//			},
//			"three": func(exp ast.Expression) {
//				testInfixExpression(t, exp, 15, "/", 5)
//			},
//		}
//
//		for key, value := range hash.Paris {
//			literal, ok := key.(*ast.StringLiteral)
//			if !ok {
//				t.Errorf("key is not ast.StringLiteral. got=%T", key)
//			}
//
//			testFunc, ok := tests[literal.String()]
//			if !ok {
//				t.Errorf("No test function for key %q found", literal.String())
//			}
//
//			testFunc(value)
//		}
//	}
//
//	func TestMacroLiteralParsing(t *testing.T) {
//		input := `macro(x, y) { x + y; }`
//
//		l := lexer.New(input)
//		p := New(l)
//		program := p.ParseProgram()
//		checkParserErrors(t, p)
//
//		if len(program.Statements) != 1 {
//			t.Fatalf("program.Statements does not contain %d statements. got=%d", 1, len(program.Statements))
//		}
//
//		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
//		if !ok {
//			t.Fatalf("statement is not ast.ExpresssionStatement. got=%T", program.Statements[0])
//		}
//
//		macro, ok := stmt.Expression.(*ast.MacroLiteral)
//		if !ok {
//			t.Fatalf("stmt.Expression is not ast.MacroLiteral. got=%T", stmt.Expression)
//		}
//
//		if len(macro.Parameters) != 2 {
//			t.Fatalf("macro literal parameters wrong. want 2, got=%d\n", len(macro.Parameters))
//		}
//
//		testLiteralExpression(t, macro.Parameters[0], "x")
//		testLiteralExpression(t, macro.Parameters[1], "y")
//
//		if len(macro.Body.Statements) != 1 {
//			t.Fatalf("macro.Body.Statements has not 1 statement. got=%d\n", len(macro.Body.Statements))
//		}
//
//		bodyStmt, ok := macro.Body.Statements[0].(*ast.ExpressionStatement)
//		if !ok {
//			t.Fatalf("macro body stmt is not ast.ExpressionStatement. got=%T", macro.Body.Statements[0])
//		}
//
//		testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
//	}
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

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
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
	testExpression(stmt.(*ast.LetStatement).Value)(t, i, stmt.(*ast.LetStatement).Value, expected[1:])
}

func testReturnStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.ReturnStatement{}), stmt, "test["+strconv.Itoa(i)+"] - ast.Statement unexpected type")
	assertions.AssertStringEquals(t, "return", stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - ast.Statement.TokenLiteral() wrong")
	if 0 == len(expected) {
		t.Fatalf("testReturnStatement: len(expected) wrong: expected=>0, actual=%d", len(expected))
	}
	testExpression(stmt.(*ast.ReturnStatement).ReturnValue)(t, i, stmt.(*ast.ReturnStatement).ReturnValue, expected)
}

func testExpressionStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.ExpressionStatement{}), stmt, "test["+strconv.Itoa(i)+"] - ast.Statement unexpected type")
	if 1 >= len(expected) {
		t.Fatalf("testExpressionStatement: len(expected) wrong: expected=>1, actual=%d", len(expected))
	}
	literal, ok := expected[0].(string)
	if !ok {
		t.Fatalf("testExpressionStatement: expected[0] unexpected type: expected=string, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, literal, stmt.TokenLiteral(), "test["+strconv.Itoa(i)+"] - ast.Statement.TokenLiteral() wrong")
	testExpression(stmt.(*ast.ExpressionStatement).Expression)(t, i, stmt.(*ast.ExpressionStatement).Expression, expected[1:])
}

func testBlockStatement(t *testing.T, i int, stmt ast.Statement, expected ...any) {
	assertions.AssertNotNull(t, stmt, "test["+strconv.Itoa(i)+"] - stmt is nil")
	if 0 == len(expected) {
		t.Fatalf("testBlockStatement: len(expected) wrong: expected=>0, actual=%d", len(expected))
	}
	assertions.AssertIntEquals(t, 1, len((stmt.(*ast.BlockStatement)).Statements), "test["+strconv.Itoa(i)+"] - program.Statements wrong")
	testStatement(stmt.(*ast.BlockStatement).Statements[0])(t, i, stmt.(*ast.BlockStatement).Statements[0], expected)
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
	var value int64
	switch expected[0].(type) {
	case int:
		value = int64(expected[0].(int))
	case int64:
		value = expected[0].(int64)
	default:
		t.Fatalf("testIntegerLiteral: expected[0] unexpected type: expected=int64, actual=%T", expected[0])
	}
	assertions.AssertStringEquals(t, fmt.Sprintf("%d", value), exp.(*ast.IntegerLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.IntegerLiteral).TokenLiteral() wrong")
	assertions.AssertInt64Equals(t, value, exp.(*ast.IntegerLiteral).Value, "test["+strconv.Itoa(i)+"] exp.(*ast.IntegerLiteral).Value wrong")
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
	testExpression(exp.(*ast.InfixExpression).Right)(t, i, exp.(*ast.InfixExpression).Right, expected[0])
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
	testExpression(exp.(*ast.IfExpression).Condition)(t, i, exp.(*ast.IfExpression).Condition, expected[0:3])
	testBlockStatement(t, i, exp.(*ast.IfExpression).Consequence, expected[3])
	if exp.(*ast.IfExpression).Alternative != nil {
		testBlockStatement(t, i, exp.(*ast.IfExpression).Alternative, expected[4])
	}
}

func testFunctionLiteral(t *testing.T, i int, exp ast.Expression, expected ...any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(ast.FunctionLiteral{}), exp, "test["+strconv.Itoa(i)+"] - ast.Expression unexpected type")
	assertions.AssertStringEquals(t, "fn", exp.(*ast.FunctionLiteral).TokenLiteral(), "test["+strconv.Itoa(i)+"] exp.(*ast.IfExpression).TokenLiteral() wrong")
	if 4 != len(expected) {
		t.Fatalf("testFunctionLiteral: len(expected) wrong: expected=5, actual=%d", len(expected))
	}
	parameters, ok := expected[0].([]string)
	if !ok {
		t.Fatalf("testFunctionLiteral: expected[0] unexpected type: expected=[]string, actual=%T", expected[0])
	}
	for n, param := range parameters {
		testExpression(exp.(*ast.FunctionLiteral).Parameters[n])(t, i, exp.(*ast.FunctionLiteral).Parameters[n], param)
	}
	testBlockStatement(t, i, exp.(*ast.FunctionLiteral).Body, expected[1:])
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
