package token

import "strconv"

type TokenType int

type Token struct {
	Type    TokenType
	Start   int
	Length  int
	Line    int
	Literal string
}

const (
	EOF = iota

	/*
	 * Operators
	 */
	ASSIGN
	PLUS
	MINUS
	BANG
	ASTERISK
	SLASH

	LT
	GT
	EQ
	NEQ

	/*
	 * Delimiters
	 */
	COMMA
	COLON
	SCOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACKET
	RBRACKET

	/*
	 * Identifiers + literals
	 */
	IDENT
	NUMBER
	STRING

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

var literals = [...]string{
	EOF: "",

	/*
	 * Operators
	 */
	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",

	LT:  "<",
	GT:  ">",
	EQ:  "==",
	NEQ: "!=",

	/*
	 * Delimiters
	 */
	COMMA:  ",",
	COLON:  ":",
	SCOLON: ";",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	/*
	 * Identifiers + literals
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

func (tt TokenType) String() string {
	s := ""
	if 0 <= tt && tt < TokenType(len(literals)) {
		s = literals[tt]
	}
	if tt != 0 && s == "" {
		s = "TokenType(" + strconv.Itoa(int(tt)) + ")"
	}
	return s
}

var keywords = map[string]TokenType{
	"fn":     FN,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func LookupIdent(literal string) TokenType {
	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}
