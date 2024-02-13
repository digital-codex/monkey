package lexer

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/token"
	"strconv"
	"testing"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			`let five = 5;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let ten = 10;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let add = fn(x, y) { x + y; }`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.FUNCTION, Literal: "fn"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let result = add(five, ten);`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "result"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.IDENT, Literal: "add"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "five"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "ten"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let zero = 5 - 5 / 5 * 5;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "zero"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.INT, Literal: "5"},
				{Type: token.SLASH, Literal: "/"},
				{Type: token.INT, Literal: "5"},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let less = 5 < 10; let greater = 10 > 5;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "less"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "5"},
				{Type: token.LT, Literal: "<"},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "greater"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.GT, Literal: ">"},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let bool = if (!(5 < 10)) { return true; } else { return false; }`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "bool"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.IF, Literal: "if"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.BANG, Literal: "!"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.INT, Literal: "5"},
				{Type: token.LT, Literal: "<"},
				{Type: token.INT, Literal: "10"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RETURN, Literal: "return"},
				{Type: token.TRUE, Literal: "true"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.ELSE, Literal: "else"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.RETURN, Literal: "return"},
				{Type: token.FALSE, Literal: "false"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let equal = 10 == 10;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "equal"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.EQ, Literal: "=="},
				{Type: token.INT, Literal: "10"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let not_equal = 10 != 5;`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "not_equal"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INT, Literal: "10"},
				{Type: token.NEQ, Literal: "!="},
				{Type: token.INT, Literal: "5"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let foobar = "foobar"; let foo_bar = "foo bar";`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "foobar"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.STRING, Literal: "foobar"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "foo_bar"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.STRING, Literal: "foo bar"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let array = [5, 10];`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "array"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.LBRACKET, Literal: "["},
				{Type: token.INT, Literal: "5"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.INT, Literal: "10"},
				{Type: token.RBRACKET, Literal: "]"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let map = {"foo": "bar"};`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "map"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.STRING, Literal: "foo"},
				{Type: token.COLON, Literal: ":"},
				{Type: token.STRING, Literal: "bar"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			`let sub = macro(x, y) { quote(unquote(x) - unquote(y)); }`,
			[]token.Token{
				{Type: token.LET, Literal: "let"},
				{Type: token.IDENT, Literal: "sub"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.MACRO, Literal: "macro"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.COMMA, Literal: ","},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.LBRACE, Literal: "{"},
				{Type: token.IDENT, Literal: "quote"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "unquote"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.MINUS, Literal: "-"},
				{Type: token.IDENT, Literal: "unquote"},
				{Type: token.LPAREN, Literal: "("},
				{Type: token.IDENT, Literal: "y"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.RPAREN, Literal: ")"},
				{Type: token.SEMICOLON, Literal: ";"},
				{Type: token.RBRACE, Literal: "}"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for i, test := range tests {
		l := New(test.input)
		for _, expected := range test.expected {
			actual := l.NextToken()
			assertions.AssertStructEquals(t, expected.Type, actual.Type, "test["+strconv.Itoa(i)+"] - type wrong")
			assertions.AssertStringEquals(t, expected.Literal, actual.Literal, "test["+strconv.Itoa(i)+"] - literal wrong")
		}
	}
}
