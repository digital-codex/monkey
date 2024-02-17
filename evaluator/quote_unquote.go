package evaluator

import (
	"fmt"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/object"
	"github.com/digital-codex/monkey/token"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = condense(node, env)
	return &object.Quote{Node: node}
}

func condense(quoted ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(quoted, func(node ast.Node) ast.Node {
		if !isUnquotedCall(node) {
			return node
		}

		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		if len(call.Argument) != 1 {
			return node
		}

		unquoted := Eval(call.Argument[0], env)
		return convert(unquoted)
	})
}

func isUnquotedCall(node ast.Node) bool {
	call, ok := node.(*ast.CallExpression)
	if !ok {
		return false
	}

	return call.Function.TokenLexeme() == "unquote"
}

func convert(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Integer:
		return &ast.IntegerLiteral{Token: token.Token{Type: token.NUMBER, Lexeme: fmt.Sprintf("%d", obj.Value)}, Value: obj.Value}
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
