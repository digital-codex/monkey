package token

import "strconv"

type Type int

type Token struct {
	Type   Type
	Start  int
	Length int
	Line   int
	Lexeme string
}

const (
	EOF Type = iota

	/*
	 * Operators
	 */
	EQUAL
	EQUAL_EQUAL
	BANG
	BANG_EQUAL

	PLUS
	MINUS
	STAR
	SLASH

	LESS
	MORE

	/*
	 * Delimiters
	 */
	COMMA
	COLON
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	/*
	 * Identifiers + lexemes
	 */
	STRING
	IDENT
	NUMBER

	/*
	 * Keywords
	 */
	FN
	LET
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	MACRO

	ILLEGAL
)

var lexemes = [...]string{
	EOF: "",

	/*
	 * Operators
	 */
	EQUAL: "=",
	PLUS:  "+",
	MINUS: "-",
	BANG:  "!",
	STAR:  "*",
	SLASH: "/",

	LESS:        "<",
	MORE:        ">",
	EQUAL_EQUAL: "==",
	BANG_EQUAL:  "!=",

	/*
	 * Delimiters
	 */
	COMMA:     ",",
	COLON:     ":",
	SEMICOLON: ";",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	/*
	 * Identifiers + lexemes
	 */
	IDENT:  "IDENT",
	NUMBER: "NUMBER",
	STRING: "STRING",

	/*
	 * Keywords
	 */
	FN:     "fn",
	LET:    "let",
	TRUE:   "true",
	FALSE:  "false",
	IF:     "if",
	ELSE:   "else",
	RETURN: "return",
	MACRO:  "macro",

	ILLEGAL: "ILLEGAL",
}

func (t Type) String() string {
	s := ""
	if 0 <= t && t < Type(len(lexemes)) {
		s = lexemes[t]
	}
	if t != 0 && s == "" {
		s = "Type(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}
