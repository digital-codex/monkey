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

type Lexer struct {
	source string

	start   int // start position in source of Token under examination
	current int // current position in source of Token under examination

	line    int
	lineIdx int

	eh       ErrorHandler
	errorCnt int
}

var keywords = map[string]token.Type{
	"fn":     token.FN,
	"if":     token.IF,
	"let":    token.LET,
	"else":   token.ELSE,
	"true":   token.TRUE,
	"false":  token.FALSE,
	"macro":  token.MACRO,
	"return": token.RETURN,
}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func New(input string, eh ErrorHandler) *Lexer {
	return &Lexer{input, 0, 0, 1, 0, eh, 0}
}

func (l *Lexer) Next() token.Token {
	for l.current < len(l.source) {
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
		case '*':
			return l.emit(token.STAR)
		case '/':
			if l.match('/') {
				l.skip(isNotNLAndEOF)
			} else {
				return l.emit(token.SLASH)
			}
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
		case '\t', '\n', '\r', ' ':
			l.skip(isWhiteSpace)
		case '"':
			return l.string()
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
	lit := l.read(isAlphaNumeric)
	t := token.IDENT
	if tt, ok := keywords[lit]; ok {
		t = tt
	}
	return l.emitWithLexeme(t, lit)
}

func (l *Lexer) number() token.Token {
	for ch := l.peek(0); isDigit(ch) && ch != 0; ch = l.peek(0) {
		l.advance()
	}

	if l.peek(0) == '.' && isDigit(l.peek(1)) {
		// consume dot
		l.advance()

		for ch := l.peek(0); isDigit(ch) && ch != 0; ch = l.peek(0) {
			l.advance()
		}
	}

	return l.emitWithLexeme(token.NUMBER, l.source[l.start:l.current])
}

func (l *Lexer) string() token.Token {
	// consume leading double-quote
	l.advance()

	for ch := l.peek(0); ch != '"' && ch != '\n' && ch != 0; ch = l.peek(0) {
		l.advance()
	}

	if l.peek(0) != '"' {
		return l.emitWithLexeme(token.ILLEGAL, l.error(UNTERMINATED_STRING))
	}

	// consume trailing double-quote
	l.advance()
	return l.emitWithLexeme(token.STRING, l.source[l.start+1:l.current-1])
}

func (l *Lexer) unexpected() token.Token {
	tok := l.emit(token.ILLEGAL)
	tok.Lexeme = l.error(UNEXPECTED_CHARACTER)
	return tok
}

func (l *Lexer) skip(condition func(byte) bool) {
	for ch := l.peek(0); condition(ch); ch = l.peek(0) {
		if ch == '\n' {
			l.lineIdx = l.current
			l.line++
		}
		l.advance()
	}
}

func (l *Lexer) read(condition func(byte) bool) string {
	for condition(l.peek(0)) {
		l.advance()
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

func (l *Lexer) advance() {
	if l.current < len(l.source) {
		l.current++
	}
}

func (l *Lexer) match(ch byte) bool {
	if l.peek(1) == ch {
		l.advance()
		return true
	}
	return false
}

func (l *Lexer) emit(t token.Type) token.Token {
	l.advance()
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
	start := 0
	if l.lineIdx != 0 {
		start = l.lineIdx + 1
	}
	line := fmt.Sprintf("%d | %s\n", l.line, l.source[start:l.current])
	out.WriteString(line)
	off := len(line)
	out.WriteString(strings.Repeat(" ", off+2))
	out.WriteString("^--- Here")

	if l.eh != nil {
		l.eh(errors.New(out.String()))
	}
	l.errorCnt++

	return string(e)
}

/*****************************************************************************
 *                                 UTILITIES                                 *
 *****************************************************************************/

func isNotNLAndEOF(ch byte) bool {
	return ch != '\n' && ch != 0
}

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isAlpha(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isAlphaNumeric(ch byte) bool {
	return isAlpha(ch) || isDigit(ch)
}
