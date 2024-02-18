package evaluator

import (
	"github.com/digital-codex/assertions"
	"github.com/digital-codex/monkey/object"
	"github.com/digital-codex/monkey/parser"
	"reflect"
	"strconv"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		input    string
		expected any
	}{
		{`let a = 5; a;`, 5},
		{`let a = 5 * 5; a;`, 25},
		{`let a = 5; let b = a; b;`, 5},
		{`let a = 5; let b = a; let c = a + b + 5; c;`, 15},
		{`return 10;`, 10},
		{`return 10; 9;`, 10},
		{`return 2 * 5; 9;`, 10},
		{`9; return 2 * 5; 9;`, 10},
		{`if (10 > 1) { if (10 > 1) { return 10; } return 1; }`, 10},
		{`5`, 5},
		{`10`, 10},
		{`-5`, -5},
		{`-10`, -10},
		{`5 + 5 + 5 + 5 - 10`, 10},
		{`2 * 2 * 2 * 2 * 2`, 32},
		{`-50 + 100 + -50`, 0},
		{`5 * 2 + 10`, 20},
		{`5 + 2 * 10`, 25},
		{`20 + 2 * -10`, 0},
		{`50 / 2 * 2 + 10`, 60},
		{`2 * (5 + 10)`, 30},
		{`3 * 3 * 3 + 10`, 37},
		{`3 * (3 * 3) + 10`, 37},
		{`(5 + 10 * 2 + 15 / 3) * 2 + -10`, 50},
		{`true`, true},
		{`false`, false},
		{`1 < 2`, true},
		{`1 > 2`, false},
		{`1 < 1`, false},
		{`1 > 1`, false},
		{`1 == 1`, true},
		{`1 != 1`, false},
		{`1 == 2`, false},
		{`1 != 2`, true},
		{`true == true`, true},
		{`false == false`, true},
		{`true == false`, false},
		{`true != false`, true},
		{`false != true`, true},
		{`(1 < 2) == true`, true},
		{`(1 < 2) == false`, false},
		{`(1 > 2) == true`, false},
		{`(1 > 2) == false`, true},
		{`!true`, false},
		{`!false`, true},
		{`!5`, false},
		{`!!true`, true},
		{`!!false`, false},
		{`!!5`, true},
		{`if (true) { 10 }`, 10},
		{`if (false) { 10 }`, NULL},
		{`if (1) { 10 }`, 10},
		{`if (1 < 2) { 10 }`, 10},
		{`if (1 > 2) { 10 }`, NULL},
		{`if (1 > 2) { 10 } else { 20 }`, 20},
		{`if (1 < 2) { 10 } else { 20 }`, 10},
		{`let identity = fn(x) { x; }; identity(5);`, 5},
		{`let identity = fn(x) { return x; }; identity(5);`, 5},
		{`let double = fn(x) { x * 2; }; double(5);`, 10},
		{`let add = fn(x, y) { x + y; }; add(5, 5);`, 10},
		{`let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));`, 20},
		{`fn(x) { x + 2; }(2);`, 4},
		{`fn(x) { x; }(5)`, 5},
		{`let adder = fn(x) { fn(y) { x + y }; }; let addTwo = adder(2); addTwo(2);`, 4},
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`"Hello World!"`, "Hello World!"},
		{`"Hello" + " " + "World!"`, "Hello World!"},
		{`[1, 2 * 2, 3 + 3][0]`, 1},
		{`[1, 2 * 2, 3 + 3][1]`, 4},
		{`[1, 2 * 2, 3 + 3][2]`, 6},
		{`[1, 2, 3][0]`, 1},
		{`[1, 2, 3][1]`, 2},
		{`[1, 2, 3][2]`, 3},
		{`let i = 0; [1][i]`, 1},
		{`[1, 2, 3][1 + 1]`, 3},
		{`let array = [1, 2, 3]; array[2]`, 3},
		{`let array = [1, 2, 3]; array[0] + array[1] + array[2]`, 6},
		{`let array = [1, 2, 3]; let i = array[0]; array[i]`, 2},
		{`[1, 2, 3][3]`, NULL},
		{`[1, 2, 3][-1]`, NULL},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }["one"]`, 1},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }["two"]`, 2},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }["three"]`, 3},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }[4]`, 4},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }[true]`, 5},
		{`let two = "two"; { "one": 10 - 9, two: 1 + 1, "thr" + "ee": 6 / 2, 4: 4, true: 5, false: 6 }[false]`, 6},
		{`{"foo": 5}["foo"]`, 5},
		{`{"foo": 5}["bar"]`, NULL},
		{`let key = "foo"; {"foo": 5}[key]`, 5},
		{`{}["foo"]`, NULL},
		{`{5: 5}[5]`, 5},
		{`{true: 5}[true]`, 5},
		{`{false: 5}[false]`, 5},
	}

	for i, test := range tests {
		evaluated := eval(test.input)
		testObject(evaluated)(t, i, evaluated, test.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`5 + true;`, "type mismatch: NUMBER + BOOLEAN"},
		{`5 + true; 5;`, "type mismatch: NUMBER + BOOLEAN"},
		{`-true;`, "unknown operator: -BOOLEAN"},
		{`true + false;`, "unknown operator: BOOLEAN + BOOLEAN"},
		{`5; true + false; 5`, "unknown operator: BOOLEAN + BOOLEAN"},
		{`if (10 > 1) { true + false }`, "unknown operator: BOOLEAN + BOOLEAN"},
		{`if (10 > 1) { true + false }`, "unknown operator: BOOLEAN + BOOLEAN"},
		{`if (10 > 1) { if (10 > 1) { return true + false; } return 1; }`, "unknown operator: BOOLEAN + BOOLEAN"},
		{`len(1)`, "argument to `len` not supported, got NUMBER"},
		{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		{`foobar`, "identifier not found: foobar"},
		{`"Hello" - "World"`, "unknown operator: STRING - STRING"},
		{`{"name": "Monkey"}[fn(x) { x}];`, "unusable as hash key: FUNCTION"},
	}

	for i, test := range tests {
		evaluated := eval(test.input)
		assertions.AssertTypeOf(t, reflect.TypeOf(object.Error{}), evaluated, "test["+strconv.Itoa(i)+"] - unexpected type")
		assertions.AssertStringEquals(t, test.expected, evaluated.(*object.Error).Message, "test["+strconv.Itoa(i)+"] - err.Message wrong")
	}
}

func eval(input string) object.Object {
	p := parser.New(input)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return Eval(program, env)
}

func testObject(o object.Object) func(*testing.T, int, object.Object, any) {
	switch o.Type() {
	case object.NUMBER:
		return testNumberObject
	case object.BOOLEAN:
		return testBooleanObject
	case object.NULL:
		return testNullObject
	case object.STRING:
		return testStringObject
	default:
		panic("unsupported object type " + o.Type().String())
	}
}

func testNumberObject(t *testing.T, i int, o object.Object, expected any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(object.Number{}), o, "test["+strconv.Itoa(i)+"] - unexpected type")
	value, ok := expected.(int)
	if !ok {
		t.Fatalf("testNumberObject: expect unexpected type: expect=int, actual=%T", expected)
	}
	assertions.AssertInt64Equals(t, int64(value), o.(*object.Number).Value, "test["+strconv.Itoa(i)+"] - o.(*object.Number).Value wrong")
}

func testBooleanObject(t *testing.T, i int, o object.Object, expected any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(object.Boolean{}), o, "test["+strconv.Itoa(i)+"] - unexpected type")
	value, ok := expected.(bool)
	if !ok {
		t.Fatalf("testBooleanObject: expect unexpected type: expect=bool, actual=%T", expected)
	}
	assertions.AssertBoolEquals(t, value, o.(*object.Boolean).Value, "test["+strconv.Itoa(i)+"] - o.(*object.Boolean).Value wrong")
}

func testNullObject(t *testing.T, i int, o object.Object, expected any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(object.Null{}), o, "test["+strconv.Itoa(i)+"] - unexpected type")
	value, ok := expected.(*object.Null)
	if !ok {
		t.Fatalf("testNullObject: expect unexpected type: expect=object.Null, actual=%T", expected)
	}
	assertions.AssertEquals(t, value, o.(*object.Null), "test["+strconv.Itoa(i)+"] - o.(*object.Null) wrong")
}

func testStringObject(t *testing.T, i int, o object.Object, expected any) {
	assertions.AssertTypeOf(t, reflect.TypeOf(object.String{}), o, "test["+strconv.Itoa(i)+"] - unexpected type")
	value, ok := expected.(string)
	if !ok {
		t.Fatalf("testStringObject: expect unexpected type: expect=string, actual=%T", expected)
	}
	assertions.AssertStringEquals(t, value, o.(*object.String).Value, "test["+strconv.Itoa(i)+"] - o.(*object.String).Value wrong")
}
