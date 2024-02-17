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
	for i, statement := range program.Statements {
		if isMacroDefinition(statement) {
			defineMacro(statement, env)
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
		call, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro, ok := isMacroCall(call, env)
		if !ok {
			return node
		}

		args := quoteArgs(call)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)

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

func defineMacro(stmt ast.Statement, env *object.Environment) {
	let, _ := stmt.(*ast.LetStatement)
	literal, _ := let.Value.(*ast.MacroLiteral)

	macro := &object.Macro{
		Parameters: literal.Parameters,
		Env:        env,
		Body:       literal.Body,
	}

	env.Set(let.Name.Value, macro)
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

func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	var args []*object.Quote

	for _, a := range exp.Argument {
		args = append(args, &object.Quote{Node: a})
	}

	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)

	for idx, param := range macro.Parameters {
		extended.Set(param.Value, args[idx])
	}

	return extended
}
