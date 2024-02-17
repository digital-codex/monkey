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
	IF
	LET
	ELSE
	TRUE
	FALSE
	MACRO
	RETURN

	EOF
)

var lexemes = [...]string{
	ILLEGAL: "ILLEGAL",

	/*
	 * Operators
	 */
	EQUAL:       "=",
	EQUAL_EQUAL: "==",
	BANG:        "!",
	BANG_EQUAL:  "!=",

	PLUS:  "+",
	MINUS: "-",
	STAR:  "*",
	SLASH: "/",

	LESS: "<",
	MORE: ">",

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
	STRING: "STRING",
	IDENT:  "IDENT",
	NUMBER: "NUMBER",

	/*
	 * Keywords
	 */
	FN:     "fn",
	IF:     "if",
	LET:    "let",
	ELSE:   "else",
	TRUE:   "true",
	FALSE:  "false",
	MACRO:  "macro",
	RETURN: "return",

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
