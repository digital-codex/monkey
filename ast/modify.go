package ast

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Modifier func(Node) Node

/*****************************************************************************
 *                              PUBLIC FUNCTIONS                             *
 *****************************************************************************/

func Modify(node Node, modifier Modifier) Node {
	switch node := node.(type) {
	case *Program:
		for i, stmt := range node.Statements {
			node.Statements[i] = Modify(stmt, modifier).(Statement)
		}
	case *LetStatement:
		node.Value = Modify(node.Value, modifier).(Expression)
	case *ReturnStatement:
		node.ReturnValue = Modify(node.ReturnValue, modifier).(Expression)
	case *ExpressionStatement:
		node.Expression = Modify(node.Expression, modifier).(Expression)
	case *Block:
		for i, stmt := range node.Statements {
			node.Statements[i] = Modify(stmt, modifier).(Statement)
		}
	case *PrefixExpression:
		node.Right = Modify(node.Right, modifier).(Expression)
	case *InfixExpression:
		node.Left = Modify(node.Left, modifier).(Expression)
		node.Right = Modify(node.Right, modifier).(Expression)
	case *GroupedExpression:
		node.Expression = Modify(node.Expression, modifier).(Expression)
	case *IfExpression:
		node.Condition = Modify(node.Condition, modifier).(Expression)
		node.Consequence = Modify(node.Consequence, modifier).(*Block)
		if node.Alternative != nil {
			node.Alternative = Modify(node.Alternative, modifier).(*Block)
		}
	case *FunctionLiteral:
		for i, param := range node.Parameters {
			node.Parameters[i] = Modify(param, modifier).(*Identifier)
		}
		node.Body = Modify(node.Body, modifier).(*Block)
	case *CallExpression:
		node.Function = Modify(node.Function, modifier).(Expression)
		for i, arg := range node.Argument {
			node.Argument[i] = Modify(arg, modifier).(Expression)
		}
	case *ArrayLiteral:
		for i, elem := range node.Elements {
			node.Elements[i] = Modify(elem, modifier).(Expression)
		}
	case *IndexExpression:
		node.Left = Modify(node.Left, modifier).(Expression)
		node.Index = Modify(node.Index, modifier).(Expression)
	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey := Modify(key, modifier).(Expression)
			newVal := Modify(val, modifier).(Expression)
			newPairs[newKey] = newVal
		}
		node.Pairs = newPairs
	}

	return modifier(node)
}
