package token

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
	DOT
	COLON
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	/*
	 * Identifiers + Literals
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

var tokens = [...]string{
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
	DOT:       ".",
	COLON:     ":",
	SEMICOLON: ";",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	/*
	 * Identifiers + Literals
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
	return tokens[t]
}
