package evaluator

import (
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/object"
)

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func DefineMacros(program *ast.Program, env *object.Environment) {
	var definitions []int
	for i, stmt := range program.Statements {
		if isMacroDefinition(stmt) {
			let, _ := stmt.(*ast.LetStatement)
			literal, _ := let.Value.(*ast.MacroLiteral)

			macro := &object.Macro{
				Parameters: literal.Parameters,
				Env:        env,
				Body:       literal.Body,
			}

			env.Set(let.Name.Value, macro)
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i-- {
		idx := definitions[i]
		program.Statements = append(program.Statements[:idx], program.Statements[idx+1:]...)
	}
}

func ExpandMacros(program ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		ce, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(ce, env)
		if !ok {
			return node
		}

		var args []object.Object
		for _, arg := range ce.Argument {
			args = append(args, &object.Quote{Node: arg})
		}
		extendedEnv := object.ExtendEnvironment(macro, args)
		evaluated := Eval(macro.Body, extendedEnv)

		q, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return q.Node
	})
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func isMacroDefinition(node ast.Statement) bool {
	let, ok := node.(*ast.LetStatement)
	if !ok {
		return false
	}

	_, ok = let.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	return true
}

func isMacroCall(exp *ast.CallExpression, env *object.Environment) (*object.Macro, bool) {
	ident, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil, false
	}

	obj, ok := env.Get(ident.Value)
	if !ok {
		return nil, false
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil, false
	}

	return macro, true
}
