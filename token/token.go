package token

import "strconv"

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Type int

type Token struct {
	Type   Type
	Start  int
	Length int
	Line   int
	Lexeme string
}

const (
	ILLEGAL Type = iota

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

	EOF
)

var lexemes = [...]string{
	ILLEGAL: "ILLEGAL",

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

	EOF: "",
}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func (t Type) String() string {
	s := ""
	if ILLEGAL <= t && t < EOF {
		s = lexemes[t]
	}
	if t != EOF && s == "" {
		s = "Type(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}
