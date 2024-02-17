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
				node: &NumberLiteral{Value: 1},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &NumberLiteral{Value: 2},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &Program{
					Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 1}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
				Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 2}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &InfixExpression{Left: &NumberLiteral{Value: 1}, Operator: "+", Right: &NumberLiteral{Value: 2}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &InfixExpression{Left: &NumberLiteral{Value: 2}, Operator: "+", Right: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &InfixExpression{Left: &NumberLiteral{Value: 2}, Operator: "+", Right: &NumberLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &InfixExpression{Left: &NumberLiteral{Value: 2}, Operator: "+", Right: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &PrefixExpression{Operator: "+", Right: &NumberLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &PrefixExpression{Operator: "+", Right: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &IndexExpression{Left: &NumberLiteral{Value: 1}, Index: &NumberLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &IndexExpression{Left: &NumberLiteral{Value: 2}, Index: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &IfExpression{
					Condition:   &NumberLiteral{Value: 1},
					Consequence: &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 1}}}},
					Alternative: &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 1}}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
				Condition:   &NumberLiteral{Value: 2},
				Consequence: &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 2}}}},
				Alternative: &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 2}}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &ReturnStatement{ReturnValue: &NumberLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &ReturnStatement{ReturnValue: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &LetStatement{Value: &NumberLiteral{Value: 1}},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
			expected: &LetStatement{Value: &NumberLiteral{Value: 2}},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &FunctionLiteral{
					Parameters: []*Identifier{},
					Body:       &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 1}}}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
				Body:       &Block{Statements: []Statement{&ExpressionStatement{Expression: &NumberLiteral{Value: 2}}}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &ArrayLiteral{
					Elements: []Expression{&NumberLiteral{Value: 1}, &NumberLiteral{Value: 1}},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
				Elements: []Expression{&NumberLiteral{Value: 2}, &NumberLiteral{Value: 2}},
			},
		},
		{
			input: struct {
				node     Node
				modifier Modifier
			}{
				node: &HashLiteral{
					Pairs: map[Expression]Expression{
						&NumberLiteral{Value: 1}: &NumberLiteral{Value: 1},
						&NumberLiteral{Value: 1}: &NumberLiteral{Value: 1},
					},
				},
				modifier: func(node Node) Node {
					integer, ok := node.(*NumberLiteral)
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
					&NumberLiteral{Value: 2}: &NumberLiteral{Value: 2},
					&NumberLiteral{Value: 2}: &NumberLiteral{Value: 2},
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
