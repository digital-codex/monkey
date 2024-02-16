package ast

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/token"
	"strconv"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    Node
		expected string
	}{
		{
			&Program{
				Statements: []Statement{
					&LetStatement{
						Token: token.Token{Type: token.LET, Lexeme: "let"},
						Name: &Identifier{
							Token: token.Token{Type: token.IDENT, Lexeme: "ident"},
							Value: "ident",
						},
						Value: &Identifier{
							Token: token.Token{Type: token.IDENT, Lexeme: "value"},
							Value: "value",
						},
					},
				},
			},
			`let ident = value;`,
		},
	}

	for i, test := range tests {
		assertions.AssertStringEquals(t, test.expected, test.input.String(), "test["+strconv.Itoa(i)+"] - program.String() wrong")
	}
}
