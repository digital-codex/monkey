package ast

import (
	"fmt"
	"github.com/digital-codex/assertions"
	"strconv"
	"testing"
)

func TestModify(t *testing.T) {
	tests := []struct {
		input struct {
			node     Node
			modifier Modifier
		}
		expected Node
	}{
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &IntegerLiteral{Value: 1},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &IntegerLiteral{Value: 2},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &Program{
					Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 1}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &Program{
				Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 2}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &InfixExpression{Left: &IntegerLiteral{Value: 1}, Operator: "+", Right: &IntegerLiteral{Value: 2}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &InfixExpression{Left: &IntegerLiteral{Value: 2}, Operator: "+", Right: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &InfixExpression{Left: &IntegerLiteral{Value: 2}, Operator: "+", Right: &IntegerLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &InfixExpression{Left: &IntegerLiteral{Value: 2}, Operator: "+", Right: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &PrefixExpression{Operator: "+", Right: &IntegerLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &PrefixExpression{Operator: "+", Right: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &IndexExpression{Left: &IntegerLiteral{Value: 1}, Index: &IntegerLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &IndexExpression{Left: &IntegerLiteral{Value: 2}, Index: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &IfExpression{
					Condition:   &IntegerLiteral{Value: 1},
					Consequence: &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 1}}}},
					Alternative: &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 1}}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &IfExpression{
				Condition:   &IntegerLiteral{Value: 2},
				Consequence: &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 2}}}},
				Alternative: &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 2}}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &ReturnStatement{ReturnValue: &IntegerLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &ReturnStatement{ReturnValue: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &LetStatement{Value: &IntegerLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &LetStatement{Value: &IntegerLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &FunctionLiteral{
					Parameters: []*Identifier{},
					Body:       &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 1}}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &FunctionLiteral{
				Parameters: []*Identifier{},
				Body:       &Block{Statements: []Statement{&ExpressionStatement{Expression: &IntegerLiteral{Value: 2}}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &ArrayLiteral{
					Elements: []Expression{&IntegerLiteral{Value: 1}, &IntegerLiteral{Value: 1}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &ArrayLiteral{
				Elements: []Expression{&IntegerLiteral{Value: 2}, &IntegerLiteral{Value: 2}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &HashLiteral{
					Pairs: map[Expression]Expression{
						&IntegerLiteral{Value: 1}: &IntegerLiteral{Value: 1},
						&IntegerLiteral{Value: 1}: &IntegerLiteral{Value: 1},
					},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*IntegerLiteral)
					if !ok {
						return node
					}

					if integer.Value != 1 {
						return node
					}

					integer.Value = 2
					return integer

				},
			},
			expected: &HashLiteral{
				Pairs: map[Expression]Expression{
					&IntegerLiteral{Value: 2}: &IntegerLiteral{Value: 2},
					&IntegerLiteral{Value: 2}: &IntegerLiteral{Value: 2},
				},
			},
		},
	}

	for i, test := range tests {
		switch test.expected.(type) {
		case *HashLiteral:
			modified := Modify(test.input.node, test.input.modifier)
			actual, ok := modified.(*HashLiteral)
			if !ok {
				t.Fatalf("TestModify: modified unexpected type: expect=&HashLiteral, actual=%T", modified)
			}
			assertions.AssertStringEquals(t, fmt.Sprint(test.expected.(*HashLiteral).Pairs), fmt.Sprint(actual.Pairs), "test["+strconv.Itoa(i)+"] - modified wrong")
		default:
			assertions.AssertDeepEquals(t, test.expected, Modify(test.input.node, test.input.modifier), "test["+strconv.Itoa(i)+"] - modified wrong")
		}
	}
}
