package ast

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "assign"},
					Value: "assign",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "value"},
					Value: "value",
				},
			},
		},
	}

	assertions.AssertStringEquals(t, "let assign = value;", program.String(), "program.String() wrong")
}
