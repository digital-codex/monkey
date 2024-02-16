package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/digital-codex/monkey/token"
	"strings"
)

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Error string
type ErrorHandler func(error)

const (
	UNEXPECTED_CHARACTER Error = "unexpected character"
	UNTERMINATED_STRING  Error = "unterminated string"
)

type Predicate func(byte) bool

type Lexer struct {
	source string

	start   int // start position in source of Token under examination
	current int // current position in source of Token under examination

	line int

	ErrorHandler ErrorHandler
	errors       bool
}

var keywords = map[string]token.Type{
	"fn":     token.FN,
	"let":    token.LET,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
	"macro":  token.MACRO,
}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func New(input string) *Lexer {
	return &Lexer{input, 0, 0, 1, nil, false}
}

func (l *Lexer) Next() token.Token {
	for l.current < len(l.source) && !l.errors {
		l.start = l.current

		ch := l.peek(0)
		switch ch {
		case '=':
			if l.match('=') {
				return l.emit(token.EQUAL_EQUAL)
			} else {
				return l.emit(token.EQUAL)
			}
		case '!':
			if l.match('=') {
				return l.emit(token.BANG_EQUAL)
			} else {
				return l.emit(token.BANG)
			}
		case '+':
			return l.emit(token.PLUS)
		case '-':
			return l.emit(token.MINUS)
		case '/':
			return l.emit(token.SLASH)
		case '*':
			return l.emit(token.STAR)
		case '<':
			return l.emit(token.LESS)
		case '>':
			return l.emit(token.MORE)
		case ',':
			return l.emit(token.COMMA)
		case ':':
			return l.emit(token.COLON)
		case ';':
			return l.emit(token.SEMICOLON)
		case '(':
			return l.emit(token.LPAREN)
		case ')':
			return l.emit(token.RPAREN)
		case '{':
			return l.emit(token.LBRACE)
		case '}':
			return l.emit(token.RBRACE)
		case '[':
			return l.emit(token.LBRACKET)
		case ']':
			return l.emit(token.RBRACKET)
		case '"':
			return l.string()
		case ' ', '\t', '\n', '\r':
			l.skip(isWhiteSpace)
		default:
			if isAlpha(ch) {
				return l.ident()
			} else if isDigit(ch) {
				return l.number()
			} else {
				return l.unexpected()
			}
		}
	}

	return l.emit(token.EOF)
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func (l *Lexer) ident() token.Token {
	lit := l.read(isAlpha)
	var t token.Type = token.IDENT
	if tt, ok := keywords[lit]; ok {
		t = tt
	}
	return l.emitWithLexeme(t, lit)
}

func (l *Lexer) number() token.Token {
	return l.emitWithLexeme(token.NUMBER, l.read(isDigit))
}

func (l *Lexer) string() token.Token {
	// consume front double-quote
	l.consume()

	for ch := l.peek(0); ch != '"' && ch != '\n' && ch != 0; ch = l.peek(0) {
		l.consume()
	}

	if l.peek(0) != '"' {
		return l.emitWithLexeme(token.ILLEGAL, l.error(UNTERMINATED_STRING))
	}

	// consume back double-quote
	l.consume()
	return l.emitWithLexeme(token.STRING, l.source[l.start+1:l.current-1])
}

func (l *Lexer) unexpected() token.Token {
	tok := l.emit(token.ILLEGAL)
	tok.Lexeme = l.error(UNEXPECTED_CHARACTER)
	return tok
}

func (l *Lexer) skip(condition Predicate) {
	for ch := l.peek(0); condition(ch); ch = l.peek(0) {
		if ch == '\n' {
			l.line++
		}
		l.consume()
	}
}

func (l *Lexer) read(condition Predicate) string {
	for condition(l.peek(0)) {
		l.consume()
	}
	return l.source[l.start:l.current]
}

func (l *Lexer) peek(n int) byte {
	if l.current+n < len(l.source) {
		return l.source[l.current+n]
	} else {
		return 0
	}
}

func (l *Lexer) consume() {
	if l.current < len(l.source) {
		l.current++
	}
}

func (l *Lexer) match(ch byte) bool {
	if l.peek(1) == ch {
		l.consume()
		return true
	}
	return false
}

func (l *Lexer) emit(t token.Type) token.Token {
	l.consume()
	return token.Token{
		Type:   t,
		Start:  l.start,
		Length: l.current - l.start,
		Line:   l.line,
		Lexeme: t.String(),
	}
}

func (l *Lexer) emitWithLexeme(t token.Type, lexeme string) token.Token {
	return token.Token{
		Type:   t,
		Start:  l.start,
		Length: l.current - l.start,
		Line:   l.line,
		Lexeme: lexeme,
	}
}

func (l *Lexer) error(e Error) string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("Error: %s", e))
	out.WriteString("\n    ")
	start := l.start
	for l.source[start] != '\n' && 0 < start {
		start--
	}
	if l.source[start] == '\n' {
		start++
	}
	line := fmt.Sprintf("%d | %s\n", l.line, l.source[start:l.current])
	out.WriteString(line)
	off := len(line)
	out.WriteString(strings.Repeat(" ", off+2))
	out.WriteString("^--- Here")

	if l.ErrorHandler != nil {
		l.ErrorHandler(errors.New(out.String()))
	}
	l.errors = true

	return string(e)
}

/*****************************************************************************
 *                                 UTILITIES                                 *
 *****************************************************************************/

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isAlpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
