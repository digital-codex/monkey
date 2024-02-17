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
	TokenLexeme() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Declaration interface {
	Statement
	declarationNode()
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

type LetDeclaration struct {
	Token token.Token // The token.LET token
	Name  *Identifier
	Value Expression
}

type ReturnStatement struct {
	Token       token.Token // The token.RETURN token
	ReturnValue Expression
}

type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

type Block struct {
	Token      token.Token // The token.LBRACE token
	Statements []Statement
}

type Identifier struct {
	Token token.Token // The token.IDENT token
	Value string
}

type NumberLiteral struct {
	Token token.Token // The token.NUMBER token
	Value int64
}

type PrefixExpression struct {
	Token    token.Token // The operator token, e.g. !
	Operator string
	Right    Expression
}

type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

type GroupedExpression struct {
	Token      token.Token // The token.LPAREN token
	Expression Expression
}

type Boolean struct {
	Token token.Token // The token.TRUE or token.FALSE token
	Value bool
}

type IfExpression struct {
	Token       token.Token // The token.IF token
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

type FunctionLiteral struct {
	Token      token.Token // The token.FN token
	Parameters []*Identifier
	Body       *Block
}

type CallExpression struct {
	Token    token.Token // The token.LPAREN token
	Function Expression  // Identifier or FunctionLiteral
	Argument []Expression
}

type StringLiteral struct {
	Token token.Token // The token.STRING token
	Value string
}

type ArrayLiteral struct {
	Token    token.Token // The token.LBRACKET token
	Elements []Expression
}

type IndexExpression struct {
	Token token.Token // The token.LBRACKET token
	Left  Expression
	Index Expression
}

type HashLiteral struct {
	Token token.Token // The token.LBRACE token
	Pairs map[Expression]Expression
}

type MacroLiteral struct {
	Token      token.Token // The token.MACRO token
	Parameters []*Identifier
	Body       *Block
}

/*****************************************************************************
 *                                   NODES                                   *
 *****************************************************************************/

func (p *Program) TokenLexeme() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLexeme()
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
 *                               DECLARATION                                 *
 *****************************************************************************/

func (ld *LetDeclaration) statementNode() {}

func (ld *LetDeclaration) TokenLexeme() string {
	return ld.Token.Lexeme
}

func (ld *LetDeclaration) String() string {
	var out bytes.Buffer

	out.WriteString(ld.TokenLexeme() + " ")
	out.WriteString(ld.Name.String())
	out.WriteString(" = ")

	if ld.Value != nil {
		out.WriteString(ld.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

/*****************************************************************************
 *                                STATEMENTS                                 *
 *****************************************************************************/

func (rs *ReturnStatement) statementNode()     {}
func (es *ExpressionStatement) statementNode() {}
func (bs *Block) statementNode()               {}

func (rs *ReturnStatement) TokenLexeme() string {
	return rs.Token.Lexeme
}
func (es *ExpressionStatement) TokenLexeme() string {
	return es.Token.Lexeme
}
func (bs *Block) TokenLexeme() string {
	return bs.Token.Lexeme
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLexeme())

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
func (bs *Block) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*****************************************************************************
 *                               EXPRESSIONS                                 *
 *****************************************************************************/

func (i *Identifier) expressionNode()         {}
func (il *NumberLiteral) expressionNode()     {}
func (pe *PrefixExpression) expressionNode()  {}
func (ie *InfixExpression) expressionNode()   {}
func (ge *GroupedExpression) expressionNode() {}
func (b *Boolean) expressionNode()            {}
func (ie *IfExpression) expressionNode()      {}
func (fl *FunctionLiteral) expressionNode()   {}
func (ce *CallExpression) expressionNode()    {}
func (sl *StringLiteral) expressionNode()     {}
func (al *ArrayLiteral) expressionNode()      {}
func (ie *IndexExpression) expressionNode()   {}
func (hl *HashLiteral) expressionNode()       {}
func (ml *MacroLiteral) expressionNode()      {}

func (i *Identifier) TokenLexeme() string {
	return i.Token.Lexeme
}
func (il *NumberLiteral) TokenLexeme() string {
	return il.Token.Lexeme
}
func (pe *PrefixExpression) TokenLexeme() string {
	return pe.Token.Lexeme
}
func (ie *InfixExpression) TokenLexeme() string {
	return ie.Token.Lexeme
}
func (ge *GroupedExpression) TokenLexeme() string {
	return ge.Token.Lexeme
}
func (b *Boolean) TokenLexeme() string {
	return b.Token.Lexeme
}
func (ie *IfExpression) TokenLexeme() string {
	return ie.Token.Lexeme
}
func (fl *FunctionLiteral) TokenLexeme() string {
	return fl.Token.Lexeme
}
func (ce *CallExpression) TokenLexeme() string {
	return ce.Token.Lexeme
}
func (sl *StringLiteral) TokenLexeme() string {
	return sl.Token.Lexeme
}
func (al *ArrayLiteral) TokenLexeme() string {
	return al.Token.Lexeme
}
func (ie *IndexExpression) TokenLexeme() string {
	return ie.Token.Lexeme
}
func (hl *HashLiteral) TokenLexeme() string {
	return hl.Token.Lexeme
}
func (ml *MacroLiteral) TokenLexeme() string {
	return ml.Token.Lexeme
}

func (i *Identifier) String() string {
	return i.Value
}
func (il *NumberLiteral) String() string {
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
func (ge *GroupedExpression) String() string {
	return ge.Expression.String()
}
func (b *Boolean) String() string {
	return b.Token.Lexeme
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString("(" + ie.Condition.String() + ")")
	out.WriteString("{" + ie.Consequence.String() + "}")
	if ie.Alternative != nil {
		out.WriteString("else")
		out.WriteString("{" + ie.Alternative.String() + "}")
	}

	return out.String()
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLexeme())
	out.WriteString("(" + strings.Join(params, ", ") + ")")
	out.WriteString("{" + fl.Body.String() + "}")

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

	out.WriteString("[" + strings.Join(elems, ", ") + "]")

	return out.String()
}
func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[" + ie.Index.String() + "]")
	out.WriteString(")")

	return out.String()
}
func (hl *HashLiteral) String() string {
	var out bytes.Buffer

	var pairs []string
	for key, value := range hl.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}

	out.WriteString("{" + strings.Join(pairs, ", ") + "}")

	return out.String()
}
func (ml *MacroLiteral) String() string {
	var out bytes.Buffer

	var params []string
	for _, param := range ml.Parameters {
		params = append(params, param.String())
	}

	out.WriteString(ml.TokenLexeme())
	out.WriteString("(" + strings.Join(params, ", ") + ")")
	out.WriteString("{" + ml.Body.String() + "}")

	return out.String()
}
