package evaluator

import (
	"fmt"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/object"
	"github.com/digital-codex/monkey/token"
)

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func quote(node ast.Node, env *object.Environment) object.Object {
	node = unquote(node, env)
	return &object.Quote{Node: node}
}

func unquote(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquotedCall(node) {
			return node
		}

		ce, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		if len(ce.Argument) != 1 {
			return node
		}

		unquoted := Eval(ce.Argument[0], env)
		return convertObjectToNode(unquoted)
	})
}

func isQuoteCall(node ast.Node) bool {
	ce, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return ce.Function.TokenLexeme() == "quote"
}

func isUnquotedCall(node ast.Node) bool {
	ce, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return ce.Function.TokenLexeme() == "unquote"
}

func convertObjectToNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Number:
		return &ast.NumberLiteral{Token: token.Token{Type: token.NUMBER, Lexeme: fmt.Sprintf("%f", obj.Value)}, Value: obj.Value}
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Lexeme: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Lexeme: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	case *object.String:
		return &ast.StringLiteral{Token: token.Token{Type: token.STRING, Lexeme: obj.Value}, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
