package lexer

import (
	"github.com/digital-codex/monkey/token"
)

type Predicate func() bool

type Lexer struct {
	source  string
	start   int // current position in source (points to current char)
	current int // current reading position in source (after current char)
	line    int
}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func New(input string) *Lexer {
	return &Lexer{input, 0, 0, 1}
}

func (l *Lexer) Next() token.Token {
	for l.current < len(l.source) {
		l.start = l.current

		ch := l.peek(0)
		switch ch {
		case '=':
			if l.match('=') {
				return l.emit(token.EQ)
			} else {
				return l.emit(token.ASSIGN)
			}
		case '!':
			if l.match('=') {
				return l.emit(token.NEQ)
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
			return l.emit(token.ASTERISK)
		case '<':
			return l.emit(token.LT)
		case '>':
			return l.emit(token.GT)
		case ',':
			return l.emit(token.COMMA)
		case ':':
			return l.emit(token.COLON)
		case ';':
			return l.emit(token.SEMICOLON)
		case '(':
			return l.emit(token.LPARENTHESIS)
		case ')':
			return l.emit(token.RPARENTHESIS)
		case '{':
			return l.emit(token.LBRACE)
		case '}':
			return l.emit(token.RBRACE)
		case '[':
			return l.emit(token.LBRACKET)
		case ']':
			return l.emit(token.RBRACKET)
		case '"':
			return l.emitWithLiteral(token.STRING, l.string())
		case ' ', '\t', '\n', '\r':
			l.skip(l.whitespace)
		default:
			if isLetter(ch) {
				lit := l.read(l.ident)
				return l.emitWithLiteral(token.LookupIdent(lit), lit)
			} else if isDigit(ch) {
				return l.emitWithLiteral(token.NUMBER, l.read(l.number))
			} else {
				l.consume()
				return l.emitWithLiteral(token.ILLEGAL, string(ch))
			}
		}
	}

	return l.emit(token.EOF)
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func (l *Lexer) string() string {
	// consume front double-quote
	l.consume()

	for ch := l.peek(0); ch != '"' && ch != 0; ch = l.peek(0) {
		l.consume()
	}

	// consume back double-quote
	l.consume()
	return l.source[l.start+1 : l.current-1]
}

func (l *Lexer) ident() bool {
	return isLetter(l.peek(0))
}

func (l *Lexer) number() bool {
	return isDigit(l.peek(0))
}

func (l *Lexer) whitespace() bool {
	return isWhiteSpace(l.peek(0))
}

func (l *Lexer) skip(condition Predicate) {
	for condition() {
		l.consume()
	}
}

func (l *Lexer) read(condition Predicate) string {
	for condition() {
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
		l.current += 1
	}
}

func (l *Lexer) match(ch byte) bool {
	if l.peek(1) == ch {
		l.consume()
		return true
	}
	return false
}

func (l *Lexer) emit(tokenType token.TokenType) token.Token {
	l.consume()
	return token.Token{
		Type:    tokenType,
		Start:   l.start,
		Length:  l.current - l.start,
		Line:    l.line,
		Literal: tokenType.String(),
	}
}

func (l *Lexer) emitWithLiteral(tokenType token.TokenType, literal string) token.Token {
	return token.Token{
		Type:    tokenType,
		Start:   l.start,
		Length:  l.current - l.start,
		Line:    l.line,
		Literal: literal,
	}
}

/*****************************************************************************
 *                                 UTILITIES                                 *
 *****************************************************************************/

func isWhiteSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
