package parser

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/lexer"
	"github.com/digital-codex/monkey/token"
	"strconv"
)

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Error string
type ErrorHandler func(error)

const (
	EXPECTED_EXPRESSION     Error = "expect expression"
	INVALID_INTEGER_LITERAL Error = "invalid integer literal"
	UNEXPECTED_TOKEN        Error = "unexpected token"
)

type (
	PrefixParseFn func() ast.Expression
	InfixParseFn  func(expression ast.Expression) ast.Expression
)

type Rule struct {
	prefix     PrefixParseFn
	infix      InfixParseFn
	precedence Precedence
}

type Parser struct {
	l *lexer.Lexer

	current token.Token
	peek    token.Token

	rules map[token.Type]Rule

	ErrorHandler ErrorHandler
	errors       []error
}

type Precedence int

const (
	_ Precedence = iota
	NONE
	EQUALITY   // ==
	COMPARISON // > or <
	TERM       // +
	FACTOR     // *
	UNARY      // -x or !x
	CALL       // myFunction(x)
	INDEX      // array[index]
)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l, l.Next(), l.Next(), make(map[token.Type]Rule), nil, []error{}}
	p.registerRule(token.EOF, nil, nil, NONE)

	p.registerRule(token.EQUAL, nil, nil, NONE)
	p.registerRule(token.EQUAL_EQUAL, nil, p.parseInfixExpression, EQUALITY)
	p.registerRule(token.BANG, p.parsePrefixExpression, nil, NONE)
	p.registerRule(token.BANG_EQUAL, nil, p.parseInfixExpression, EQUALITY)

	p.registerRule(token.PLUS, nil, p.parseInfixExpression, TERM)
	p.registerRule(token.MINUS, p.parsePrefixExpression, p.parseInfixExpression, TERM)
	p.registerRule(token.STAR, nil, p.parseInfixExpression, FACTOR)
	p.registerRule(token.SLASH, nil, p.parseInfixExpression, FACTOR)

	p.registerRule(token.LESS, nil, p.parseInfixExpression, COMPARISON)
	p.registerRule(token.MORE, nil, p.parseInfixExpression, COMPARISON)

	p.registerRule(token.COMMA, nil, nil, NONE)
	p.registerRule(token.COLON, nil, nil, NONE)
	p.registerRule(token.SEMICOLON, nil, nil, NONE)

	p.registerRule(token.LPAREN, p.parseGroupedExpression, p.parseCallExpression, CALL)
	p.registerRule(token.LBRACE, p.parseHashLiteral, nil, NONE)
	p.registerRule(token.LBRACKET, p.parseArrayLiteral, p.parseIndexExpression, INDEX)

	p.registerRule(token.IDENT, p.parseIdentifier, nil, NONE)
	p.registerRule(token.NUMBER, p.parseIntegerLiteral, nil, NONE)
	p.registerRule(token.STRING, p.parseStringLiteral, nil, NONE)

	p.registerRule(token.FN, p.parseFunctionLiteral, nil, NONE)
	p.registerRule(token.LET, nil, nil, NONE)
	p.registerRule(token.TRUE, p.parseBoolean, nil, NONE)
	p.registerRule(token.FALSE, p.parseBoolean, nil, NONE)
	p.registerRule(token.IF, p.parseIfExpression, nil, NONE)
	p.registerRule(token.ELSE, nil, nil, NONE)
	p.registerRule(token.RETURN, nil, nil, NONE)
	p.registerRule(token.MACRO, p.parseMacroLiteral, nil, NONE)

	p.registerRule(token.ILLEGAL, nil, nil, NONE)

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.current.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.next()
	}

	return program
}

func (p *Parser) Errors() []error {
	return p.errors
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func (p *Parser) parseStatement() ast.Statement {
	switch p.current.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.current}

	if !p.expect(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.current, Value: p.current.Lexeme}

	if !p.expect(token.EQUAL) {
		return nil
	}

	p.next()

	stmt.Value = p.parseExpression(NONE)

	if p.peekTokenIs(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.current}

	p.next()

	stmt.ReturnValue = p.parseExpression(NONE)

	if p.peekTokenIs(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.current}

	stmt.Expression = p.parseExpression(NONE)

	if p.peekTokenIs(token.SEMICOLON) {
		p.next()
	}

	return stmt
}

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Token: p.current}
	block.Statements = []ast.Statement{}

	p.next()

	for !p.currentTokenIs(token.RBRACE) && !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.next()
	}

	return block
}

func (p *Parser) parseExpression(precedence Precedence) ast.Expression {
	prefix := p.rule(p.current.Type).prefix
	if prefix == nil {
		p.error(EXPECTED_EXPRESSION)
		return nil
	}
	expr := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.rule(p.peek.Type).precedence {
		infix := p.rule(p.peek.Type).infix
		if infix == nil {
			return expr
		}

		p.next()

		expr = infix(expr)
	}

	return expr
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.current, Value: p.current.Lexeme}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	expr := &ast.IntegerLiteral{Token: p.current}

	value, err := strconv.ParseInt(p.current.Lexeme, 0, 64)
	if err != nil {
		p.error(INVALID_INTEGER_LITERAL)
		return nil
	}

	expr.Value = value
	return expr
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expr := &ast.PrefixExpression{Token: p.current, Operator: p.current.Lexeme}

	p.next()

	expr.Right = p.parseExpression(UNARY)

	return expr
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expr := &ast.InfixExpression{Token: p.current, Left: left, Operator: p.current.Lexeme}

	precedence := p.rule(p.current.Type).precedence
	p.next()
	expr.Right = p.parseExpression(precedence)

	return expr
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	expr := &ast.GroupedExpression{Token: p.current}
	p.next()

	expr.Expression = p.parseExpression(NONE)
	if !p.expect(token.RPAREN) {
		return nil
	}

	return expr
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.current, Value: p.currentTokenIs(token.TRUE)}
}

func (p *Parser) parseIfExpression() ast.Expression {
	expr := &ast.IfExpression{Token: p.current}

	if !p.expect(token.LPAREN) {
		return nil
	}

	p.next()
	expr.Condition = p.parseExpression(NONE)

	if !p.expect(token.RPAREN) {
		return nil
	}

	if !p.expect(token.LBRACE) {
		return nil
	}

	expr.Consequence = p.parseBlock()

	if p.peekTokenIs(token.ELSE) {
		p.next()

		if !p.expect(token.LBRACE) {
			return nil
		}

		expr.Alternative = p.parseBlock()
	}

	return expr
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	expr := &ast.FunctionLiteral{Token: p.current}

	if !p.expect(token.LPAREN) {
		return nil
	}

	expr.Parameters = p.parseParameters()

	if !p.expect(token.LBRACE) {
		return nil
	}
	expr.Body = p.parseBlock()
	return expr
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	expr := &ast.CallExpression{Token: p.current, Function: function}
	expr.Argument = p.parseExpressions(token.RPAREN)
	return expr
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.current, Value: p.current.Lexeme}
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	expr := &ast.ArrayLiteral{Token: p.current}
	expr.Elements = p.parseExpressions(token.RBRACKET)
	return expr
}

func (p *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	expr := &ast.IndexExpression{Token: p.current, Left: left}

	p.next()
	expr.Index = p.parseExpression(NONE)

	if !p.expect(token.RBRACKET) {
		return nil
	}

	return expr
}

func (p *Parser) parseHashLiteral() ast.Expression {
	expr := &ast.HashLiteral{Token: p.current}
	expr.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(token.RBRACE) {
		p.next()
		key := p.parseExpression(NONE)

		if !p.expect(token.COLON) {
			return nil
		}

		p.next()
		value := p.parseExpression(NONE)

		expr.Pairs[key] = value

		if !p.peekTokenIs(token.RBRACE) && !p.expect(token.COMMA) {
			return nil
		}
	}

	if !p.expect(token.RBRACE) {
		return nil
	}

	return expr
}

func (p *Parser) parseMacroLiteral() ast.Expression {
	expr := &ast.MacroLiteral{Token: p.current}

	if !p.expect(token.LPAREN) {
		return nil
	}

	expr.Parameters = p.parseParameters()

	if !p.expect(token.LBRACE) {
		return nil
	}
	expr.Body = p.parseBlock()
	return expr
}

func (p *Parser) parseParameters() []*ast.Identifier {
	var idents []*ast.Identifier

	if p.peekTokenIs(token.RPAREN) {
		p.next()
		return idents
	}

	p.next()

	ident := &ast.Identifier{Token: p.current, Value: p.current.Lexeme}
	idents = append(idents, ident)

	for p.peekTokenIs(token.COMMA) {
		p.next()
		p.next()
		ident = &ast.Identifier{Token: p.current, Value: p.current.Lexeme}
		idents = append(idents, ident)
	}

	if !p.expect(token.RPAREN) {
		return nil
	}

	return idents
}

func (p *Parser) parseExpressions(end token.Type) []ast.Expression {
	var exprs []ast.Expression

	if p.peekTokenIs(end) {
		p.next()
		return exprs
	}

	p.next()
	exprs = append(exprs, p.parseExpression(NONE))

	for p.peekTokenIs(token.COMMA) {
		p.next()
		p.next()
		exprs = append(exprs, p.parseExpression(NONE))
	}

	if !p.expect(end) {
		return nil
	}

	return exprs
}

func (p *Parser) currentTokenIs(tok token.Type) bool {
	return p.current.Type == tok
}

func (p *Parser) peekTokenIs(tok token.Type) bool {
	return p.peek.Type == tok
}

func (p *Parser) next() {
	p.current = p.peek
	p.peek = p.l.Next()
}

func (p *Parser) expect(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.next()
		return true
	} else {
		p.error(UNEXPECTED_TOKEN)
		return false
	}
}

func (p *Parser) rule(t token.Type) Rule {
	return p.rules[t]
}

func (p *Parser) registerRule(t token.Type, prefix PrefixParseFn, infix InfixParseFn, precedence Precedence) {
	p.rules[t] = Rule{prefix, infix, precedence}
}

func (p *Parser) error(e Error) {
	var out bytes.Buffer

	switch e {
	case EXPECTED_EXPRESSION:
		out.WriteString(fmt.Sprintf("Error:%d:%d: %s got %q\n", p.current.Line, p.current.Start, e, p.current.Lexeme))
	case INVALID_INTEGER_LITERAL:
		out.WriteString(fmt.Sprintf("Error:%d:%d: %s %q\n", p.current.Line, p.current.Start, e, p.current.Lexeme))
	case UNEXPECTED_TOKEN:
		out.WriteString(fmt.Sprintf("Error:%d:%d: %s wanted %q\n", p.peek.Line, p.peek.Start, e, p.peek.Lexeme))
	}

	err := errors.New(out.String())
	if p.ErrorHandler != nil {
		p.ErrorHandler(err)
	}
	p.errors = append(p.errors, err)
}
