package ast

type Modifier func(Node) Node

func Modify(node Node, modifier Modifier) Node {
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *LetStatement:
		node.Value, _ = Modify(node.Value, modifier).(Expression)
	case *ReturnStatement:
		node.ReturnValue, _ = Modify(node.ReturnValue, modifier).(Expression)
	case *ExpressionStatement:
		node.Expression, _ = Modify(node.Expression, modifier).(Expression)
	case *Block:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *PrefixExpression:
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *InfixExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Right, _ = Modify(node.Right, modifier).(Expression)
	case *GroupedExpression:
		node.Expression, _ = Modify(node.Expression, modifier).(Expression)
	case *IfExpression:
		node.Condition, _ = Modify(node.Condition, modifier).(Expression)
		node.Consequence, _ = Modify(node.Consequence, modifier).(*Block)
		if node.Alternative != nil {
			node.Alternative, _ = Modify(node.Alternative, modifier).(*Block)
		}
	case *FunctionLiteral:
		for i, param := range node.Parameters {
			node.Parameters[i], _ = Modify(param, modifier).(*Identifier)
		}
		node.Body, _ = Modify(node.Body, modifier).(*Block)
	case *CallExpression:
		node.Function, _ = Modify(node.Function, modifier).(Expression)
		for i, arg := range node.Argument {
			node.Argument[i], _ = Modify(arg, modifier).(Expression)
		}
	case *ArrayLiteral:
		for i, elem := range node.Elements {
			node.Elements[i], _ = Modify(elem, modifier).(Expression)
		}
	case *IndexExpression:
		node.Left, _ = Modify(node.Left, modifier).(Expression)
		node.Index, _ = Modify(node.Index, modifier).(Expression)
	case *HashLiteral:
		newPairs := make(map[Expression]Expression)
		for key, val := range node.Pairs {
			newKey, _ := Modify(key, modifier).(Expression)
			newVal, _ := Modify(val, modifier).(Expression)
			newPairs[newKey] = newVal
		}
		node.Pairs = newPairs
	}

	return modifier(node)
}
