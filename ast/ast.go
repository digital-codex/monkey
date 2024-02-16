package ast

import (
	"bytes"
	"github.com/digital-codex/monkey/token"
	"strings"
)

/*****************************************************************************
 *                                INTERFACES                                 *
 *****************************************************************************/

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the 'let' token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

type Boolean struct {
	Token token.Token
	Value bool
}

type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

type FunctionLiteral struct {
	Token      token.Token // The 'fn' token
	Parameters []*Identifier
	Body       *BlockStatement
}

type CallExpression struct {
	Token    token.Token // The LPAREN token
	Function Expression  // Identifier or FunctionLiteral
	Argument []Expression
}

type StringLiteral struct {
	Token token.Token
	Value string
}

type ArrayLiteral struct {
	Token    token.Token // the '[' token
	Elements []Expression
}

type IndexExpression struct {
	Token token.Token // The '[' token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

type MacroLiteral struct {
	Token      token.Token // the 'macro' token
	Parameters []*Identifier
	Body       *BlockStatement
}

/*****************************************************************************
 *                                   NODES                                   *
 *****************************************************************************/

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*****************************************************************************
 *                                STATEMENTS                                 *
 *****************************************************************************/

func (ls *LetStatement) statementNode()        {}
func (rs *ReturnStatement) statementNode()     {}
func (es *ExpressionStatement) statementNode() {}
func (bs *BlockStatement) statementNode()      {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Lexeme
}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Lexeme
}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Lexeme
}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Lexeme
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral())

	if rs.ReturnValue != nil {
		out.WriteString(" " + rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*****************************************************************************
 *                               EXPRESSIONS                                 *
 *****************************************************************************/

func (i *Identifier) expressionNode()        {}
func (il *IntegerLiteral) expressionNode()   {}
func (pe *PrefixExpression) expressionNode() {}
func (ie *InfixExpression) expressionNode()  {}
func (b *Boolean) expressionNode()           {}
func (ie *IfExpression) expressionNode()     {}
func (fl *FunctionLiteral) expressionNode()  {}
func (ce *CallExpression) expressionNode()   {}
func (sl *StringLiteral) expressionNode()    {}
func (al *ArrayLiteral) expressionNode()     {}
func (ie *IndexExpression) expressionNode()  {}
func (hl *HashLiteral) expressionNode()      {}
func (ml *MacroLiteral) expressionNode()     {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Lexeme
}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Lexeme
}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Lexeme
}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Lexeme
}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Lexeme
}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Lexeme
}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Lexeme
}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Lexeme
}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Lexeme
}
func (al *ArrayLiteral) TokenLiteral() string {
	return al.Token.Lexeme
}
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Lexeme
}
func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Lexeme
}
func (ml *MacroLiteral) TokenLiteral() string {
	return ml.Token.Lexeme
}

func (i *Identifier) String() string {
	return i.Value
}
func (il *IntegerLiteral) String() string {
	return il.Token.Lexeme
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
func (b *Boolean) String() string {
	return b.Token.Lexeme
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString("(")
	out.WriteString(ie.Condition.String())
	out.WriteString(")")
	out.WriteString(" { " + ie.Consequence.String() + " }")
	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString(" { " + ie.Alternative.String() + " }")
	}

	return out.String()
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(" { " + fl.Body.String() + " }")

	return out.String()
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	var args []string
	for _, arg := range ce.Argument {
		args = append(args, arg.String())
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
func (sl *StringLiteral) String() string {
	return sl.Token.Lexeme
}
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer

	var elems []string
	for _, e := range al.Elements {
		elems = append(elems, e.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")

	return out.String()
}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	return out.String()
}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}
func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, param := range ml.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(ml.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString("{" + ml.Body.String() + "}")

	return out.String()
}
