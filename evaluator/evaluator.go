package evaluator

import (
	"fmt"
	"github.com/digital-codex/monkey/ast"
	"github.com/digital-codex/monkey/object"
)

/*****************************************************************************
 *                                INTERFACES                                 *
 *****************************************************************************/

type Operation interface {
	operation()
}

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

type (
	PrefixFunc func(object.Object) object.Object
	InfixFunc  func(object.Object, object.Object) object.Object
)

type PrefixOperation struct {
	right  object.Type
	result object.Type
	apply  PrefixFunc
}
type InfixOperation struct {
	left   object.Type
	right  object.Type
	result object.Type
	apply  InfixFunc
}

var operations = map[string][]Operation{
	"!": {
		&PrefixOperation{object.ANY, object.BOOLEAN, func(right object.Object) object.Object {
			return convertNativeBoolToBooleanObject(!isTruthy(right))
		}},
	},
	"==": {
		&InfixOperation{object.ANY, object.ANY, object.BOOLEAN, func(left, right object.Object) object.Object {
			if left.Type() == object.NUMBER && right.Type() == object.NUMBER {
				left := left.(*object.Number)
				right := right.(*object.Number)
				return convertNativeBoolToBooleanObject(left.Value == right.Value)
			}
			return convertNativeBoolToBooleanObject(left == right)
		}},
	},
	"!=": {
		&InfixOperation{object.ANY, object.ANY, object.BOOLEAN, func(left, right object.Object) object.Object {
			if left.Type() == object.NUMBER && right.Type() == object.NUMBER {
				left := left.(*object.Number)
				right := right.(*object.Number)
				return convertNativeBoolToBooleanObject(left.Value != right.Value)
			}
			return convertNativeBoolToBooleanObject(left != right)
		}},
	},
	"+": {
		&InfixOperation{object.NUMBER, object.NUMBER, object.NUMBER, func(left, right object.Object) object.Object {
			return &object.Number{Value: left.(*object.Number).Value + right.(*object.Number).Value}
		}},
		&InfixOperation{object.STRING, object.STRING, object.STRING, func(left, right object.Object) object.Object {
			return &object.String{Value: left.(*object.String).Value + right.(*object.String).Value}
		}},
	},
	"-": {
		&PrefixOperation{object.NUMBER, object.NUMBER, func(right object.Object) object.Object {
			return &object.Number{Value: -right.(*object.Number).Value}
		}},
		&InfixOperation{object.NUMBER, object.NUMBER, object.NUMBER, func(left, right object.Object) object.Object {
			return &object.Number{Value: left.(*object.Number).Value - right.(*object.Number).Value}
		}},
	},
	"*": {
		&InfixOperation{object.NUMBER, object.NUMBER, object.NUMBER, func(left, right object.Object) object.Object {
			return &object.Number{Value: left.(*object.Number).Value * right.(*object.Number).Value}
		}},
	},
	"/": {
		&InfixOperation{object.NUMBER, object.NUMBER, object.NUMBER, func(left, right object.Object) object.Object {
			return &object.Number{Value: left.(*object.Number).Value / right.(*object.Number).Value}
		}},
	},
	"<": {
		&InfixOperation{object.NUMBER, object.NUMBER, object.BOOLEAN, func(left, right object.Object) object.Object {
			return convertNativeBoolToBooleanObject(left.(*object.Number).Value < right.(*object.Number).Value)
		}},
	},
	">": {
		&InfixOperation{object.NUMBER, object.NUMBER, object.BOOLEAN, func(left, right object.Object) object.Object {
			return convertNativeBoolToBooleanObject(left.(*object.Number).Value > right.(*object.Number).Value)
		}},
	},
}

/*****************************************************************************
 *                                Operations                                 *
 *****************************************************************************/

func (po *PrefixOperation) operation() {}
func (io *InfixOperation) operation()  {}

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.LetStatement:
		if err, ok := evalLetStatement(node, env); !ok {
			return err
		}
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.ExpressionStatement:
		return evalExpressionStatement(node, env)
	case *ast.Block:
		return evalBlock(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.NumberLiteral:
		return evalNumberLiteral(node)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.GroupedExpression:
		return evalGroupedExpression(node, env)
	case *ast.Boolean:
		return evalBoolean(node)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.FunctionLiteral:
		return evalFunctionLiteral(node, env)
	case *ast.CallExpression:
		return evalCallExpression(node, env)
	case *ast.StringLiteral:
		return evalStringLiteral(node)
	case *ast.ArrayLiteral:
		return evalArrayLiteral(node, env)
	case *ast.IndexExpression:
		return evalIndexExpression(node, env)
	case *ast.HashLiteral:
		return evalHashLiteral(node, env)
	}

	return nil
}

/*****************************************************************************
 *                             PRIVATE FUNCTIONS                             *
 *****************************************************************************/

func evalProgram(node *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalLetStatement(node *ast.LetStatement, env *object.Environment) (object.Object, bool) {
	val := Eval(node.Value, env)
	if isError(val) {
		return val, false
	}
	env.Set(node.Name.Value, val)
	return nil, true
}

func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) object.Object {
	val := Eval(node.ReturnValue, env)
	if isError(val) {
		return val
	}
	return &object.ReturnValue{Value: val}
}

func evalExpressionStatement(node *ast.ExpressionStatement, env *object.Environment) object.Object {
	return Eval(node.Expression, env)
}

func evalBlock(node *ast.Block, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range node.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.RETURN_VALUE || rt == object.ERROR {
				return result
			}
		}
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}
	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return errorf("identifier not found: " + node.Value)
}

func evalNumberLiteral(node *ast.NumberLiteral) object.Object {
	return &object.Number{Value: node.Value}
}

func evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) object.Object {
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}

	ops, ok := operations[node.Operator]
	if !ok {
		return errorf("unknown operator: %s%s", node.Operator, right.Type())
	}

	for _, op := range ops {
		op, ok := op.(*PrefixOperation)
		if ok {
			if right.Type() == op.right || op.right == object.ANY {
				return op.apply(right)
			}
		}
	}
	return errorf("unknown operator: %s%s", node.Operator, right.Type())
}

func evalInfixExpression(node *ast.InfixExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}

	if left.Type() != right.Type() {
		return errorf("type mismatch: %s + %s", left.Type(), right.Type())
	}

	ops, ok := operations[node.Operator]
	if !ok {
		return errorf("unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
	}

	for _, op := range ops {
		op, ok := op.(*InfixOperation)
		if ok {
			if (left.Type() == op.left || op.left == object.ANY) && (right.Type() == op.right || op.right == object.ANY) {
				return op.apply(left, right)
			}
		}
	}
	return errorf("unknown operator: %s %s %s", left.Type(), node.Operator, right.Type())
}

func evalGroupedExpression(node *ast.GroupedExpression, env *object.Environment) object.Object {
	return Eval(node.Expression, env)
}

func evalBoolean(node *ast.Boolean) object.Object {
	return convertNativeBoolToBooleanObject(node.Value)
}

func evalIfExpression(node *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(node.Condition, env)
	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(node.Consequence, env)
	} else if node.Alternative != nil {
		return Eval(node.Alternative, env)
	}

	return NULL
}

func evalFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	return &object.Function{Parameters: node.Parameters, Env: env, Body: node.Body}
}

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	if node.Function.TokenLexeme() == "quote" {
		return quote(node.Argument[0], env)
	}
	fn := Eval(node.Function, env)
	if isError(fn) {
		return fn
	}
	args := evalExpressions(node.Argument, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return call(fn, args)
}

func call(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := object.ExtendEnvironment(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)

		if returnValue, ok := evaluated.(*object.ReturnValue); ok {
			return returnValue.Value
		}
		return evaluated
	case *object.Builtin:
		return fn.Fn(args...)
	default:
		return errorf("not a function: %s", fn.Type())
	}
}

func evalStringLiteral(node *ast.StringLiteral) object.Object {
	return &object.String{Value: node.Value}
}

func evalArrayLiteral(node *ast.ArrayLiteral, env *object.Environment) object.Object {
	elements := evalExpressions(node.Elements, env)
	if len(elements) == 1 && isError(elements[0]) {
		return elements[0]
	}

	return &object.Array{Elements: elements}
}

func evalIndexExpression(node *ast.IndexExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	index := Eval(node.Index, env)
	if isError(index) {
		return index
	}

	switch {
	case left.Type() == object.ARRAY && index.Type() == object.NUMBER:
		array := left.(*object.Array)
		i := index.(*object.Number).Value

		if i < 0 || i > int64(len(array.Elements)-1) {
			return NULL
		}

		return array.Elements[i]
	case left.Type() == object.HASH:
		hash := left.(*object.Hash)

		key, ok := index.(object.Hashable)
		if !ok {
			return errorf("unusable as hash key: %s", index.Type())
		}

		pair, ok := hash.Pairs[key.HashKey()]
		if !ok {
			return NULL
		}

		return pair.Value
	default:
		return errorf("index operator not supported: %s", left.Type())
	}
}

func evalHashLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return errorf("unusable as hash key: %s", key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.HashPair{Key: key, Value: value}
	}

	return &object.Hash{Pairs: pairs}
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}

	return false
}

func convertNativeBoolToBooleanObject(b bool) object.Object {
	if b {
		return TRUE
	}
	return FALSE
}

func errorf(format string, a ...any) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
