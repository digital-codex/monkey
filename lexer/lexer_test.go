package lexer

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/token"
	"log"
	"strconv"
	"testing"
)

var logger = log.Default()

func LogError(e error) {
	logger.Print(e.Error())
}

func TestNext(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			`let five = 5;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "five"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let ten = 10;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "ten"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let add = fn(x, y) { x + y; }`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "add"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.FN, Lexeme: "fn"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "x"},
				{Type: token.COMMA, Lexeme: ","},
				{Type: token.IDENT, Lexeme: "y"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.LBRACE, Lexeme: "{"},
				{Type: token.IDENT, Lexeme: "x"},
				{Type: token.PLUS, Lexeme: "+"},
				{Type: token.IDENT, Lexeme: "y"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.RBRACE, Lexeme: "}"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let result = add(five, ten);`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "result"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.IDENT, Lexeme: "add"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "five"},
				{Type: token.COMMA, Lexeme: ","},
				{Type: token.IDENT, Lexeme: "ten"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let zero = 5 - 5 / 5 * 5;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "zero"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.MINUS, Lexeme: "-"},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.SLASH, Lexeme: "/"},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.STAR, Lexeme: "*"},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let less = 5 < 10; let greater = 10 > 5;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "less"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.LESS, Lexeme: "<"},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "greater"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.MORE, Lexeme: ">"},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let bool = if (!(5 < 10)) { return true; } else { return false; }`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "bool"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.IF, Lexeme: "if"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.BANG, Lexeme: "!"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.LESS, Lexeme: "<"},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.LBRACE, Lexeme: "{"},
				{Type: token.RETURN, Lexeme: "return"},
				{Type: token.TRUE, Lexeme: "true"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.RBRACE, Lexeme: "}"},
				{Type: token.ELSE, Lexeme: "else"},
				{Type: token.LBRACE, Lexeme: "{"},
				{Type: token.RETURN, Lexeme: "return"},
				{Type: token.FALSE, Lexeme: "false"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.RBRACE, Lexeme: "}"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let equal = 10 == 10;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "equal"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.EQUAL_EQUAL, Lexeme: "=="},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let not_equal = 10 != 5;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "not_equal"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.BANG_EQUAL, Lexeme: "!="},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let foobar = "foobar"; let foo_bar = "foo bar";`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "foobar"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.STRING, Lexeme: "foobar"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "foo_bar"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.STRING, Lexeme: "foo bar"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let array = [5, 10];`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "array"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.LBRACKET, Lexeme: "["},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.COMMA, Lexeme: ","},
				{Type: token.NUMBER, Lexeme: "10"},
				{Type: token.RBRACKET, Lexeme: "]"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let map = {"foo": "bar"};`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "map"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.LBRACE, Lexeme: "{"},
				{Type: token.STRING, Lexeme: "foo"},
				{Type: token.COLON, Lexeme: ":"},
				{Type: token.STRING, Lexeme: "bar"},
				{Type: token.RBRACE, Lexeme: "}"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let sub = macro(x, y) { quote(unquote(x) - unquote(y)); }`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "sub"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.MACRO, Lexeme: "macro"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "x"},
				{Type: token.COMMA, Lexeme: ","},
				{Type: token.IDENT, Lexeme: "y"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.LBRACE, Lexeme: "{"},
				{Type: token.IDENT, Lexeme: "quote"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "unquote"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "x"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.MINUS, Lexeme: "-"},
				{Type: token.IDENT, Lexeme: "unquote"},
				{Type: token.LPAREN, Lexeme: "("},
				{Type: token.IDENT, Lexeme: "y"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.RPAREN, Lexeme: ")"},
				{Type: token.SEMICOLON, Lexeme: ";"},
				{Type: token.RBRACE, Lexeme: "}"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
	}

	for i, test := range tests {
		l := New(test.input, LogError)
		for _, expected := range test.expected {
			actual := l.Next()
			assertions.AssertEquals(t, expected.Type, actual.Type, "test["+strconv.Itoa(i)+"] - Type wrong")
			assertions.AssertStringEquals(t, expected.Lexeme, actual.Lexeme, "test["+strconv.Itoa(i)+"] - Lexeme wrong")
		}
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected []token.Token
	}{
		{
			"\n\nlet five = 5.",
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "five"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.NUMBER, Lexeme: "5"},
				{Type: token.ILLEGAL, Lexeme: "unexpected character"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
		{
			`let foobar = "foobar;`,
			[]token.Token{
				{Type: token.LET, Lexeme: "let"},
				{Type: token.IDENT, Lexeme: "foobar"},
				{Type: token.EQUAL, Lexeme: "="},
				{Type: token.ILLEGAL, Lexeme: "unterminated string"},
				{Type: token.EOF, Lexeme: ""},
			},
		},
	}

	for i, test := range tests {
		l := New(test.input, LogError)
		for _, expected := range test.expected {
			actual := l.Next()
			assertions.AssertEquals(t, expected.Type, actual.Type, "test["+strconv.Itoa(i)+"] - Type wrong")
			assertions.AssertStringEquals(t, expected.Lexeme, actual.Lexeme, "test["+strconv.Itoa(i)+"] - Lexeme wrong")
		}
	}
}
