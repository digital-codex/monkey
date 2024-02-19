package evaluator

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/object"
	"reflect"
	"strconv"
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(5)`, `5`},
		{`quote(5 + 8)`, `(5 + 8)`},
		{`quote(foobar)`, `foobar`},
		{`quote(hello + world)`, `(hello + world)`},
	}

	for i, test := range tests {
		evaluated := eval(test.input)
		assertions.AssertTypeOf(t, reflect.TypeOf(object.Quote{}), evaluated, "test["+strconv.Itoa(i)+"] - unexpected type")
		assertions.AssertNotNull(t, evaluated.(*object.Quote).Node, "test["+strconv.Itoa(i)+"] evaluated.(*object.Quote).Node wrong")
		assertions.AssertStringEquals(t, test.expected, evaluated.(*object.Quote).Node.String(), "test["+strconv.Itoa(i)+"] evaluated.(*object.Quote).Node.String() wrong")
	}
}

func TestUnquote(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`quote(unquote(4))`, `4.000000`},
		{`quote(unquote(4 + 4))`, `8.000000`},
		{`quote(8 + unquote(4 + 4))`, `(8 + 8.000000)`},
		{`quote(unquote(4 + 4) + 8)`, `(8.000000 + 8)`},
		{`let foobar = 8; quote(foobar)`, `foobar`},
		{`let foobar = 8; quote(unquote(foobar))`, `8.000000`},
		{`quote(unquote(true))`, `true`},
		{`quote(unquote(true == false))`, `false`},
		{`quote(unquote(quote(4 + 4)))`, `(4 + 4)`},
		{`let quotedInfixExpression = quote(4 + 4); quote(unquote(4 + 4) + unquote(quotedInfixExpression))`, `(8.000000 + (4 + 4))`},
	}

	for i, test := range tests {
		evaluated := eval(test.input)
		assertions.AssertTypeOf(t, reflect.TypeOf(object.Quote{}), evaluated, "test["+strconv.Itoa(i)+"] - unexpected type")
		assertions.AssertNotNull(t, evaluated.(*object.Quote).Node, "test["+strconv.Itoa(i)+"] evaluated.(*object.Quote).Node wrong")
		assertions.AssertStringEquals(t, test.expected, evaluated.(*object.Quote).Node.String(), "test["+strconv.Itoa(i)+"] evaluated.(*object.Quote).Node.String() wrong")
	}
}
