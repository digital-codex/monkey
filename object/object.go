package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"math/rand"
	"monkey/ast"
	"strings"
)

type Type string

const (
	INTEGER      = "INTEGER"
	BOOLEAN      = "BOOLEAN"
	NULL         = "NULL"
	RETURN_VALUE = "RETURN_VALUE"
	ERROR        = "ERROR"
	FUNCTION     = "FUNCTION"
	STRING       = "STRING"
	BUILTIN      = "BUILTIN"
	ARRAY        = "ARRAY"

	HASH  = "HASH"
	QUOTE = "QUOTE"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type {
	return INTEGER
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type {
	return BOOLEAN
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n *Null) Type() Type {
	return NULL
}
func (n *Null) Inspect() string {
	return "null"
}

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type {
	return RETURN_VALUE
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ERROR
}
func (e *Error) Inspect() string {
	return "Error: " + e.Message
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type {
	return FUNCTION
}
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() Type {
	return STRING
}
func (s *String) Inspect() string {
	return s.Value
}

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() Type {
	return BUILTIN
}
func (b *Builtin) Inspect() string {
	return "native function"
}

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type {
	return ARRAY
}
func (a *Array) Inspect() string {
	var out bytes.Buffer

	var elems []string
	for _, e := range a.Elements {
		elems = append(elems, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elems, ", "))
	out.WriteString("]")

	return out.String()
}

type Hashable interface {
	HashKey() HashKey
}

type HashKey struct {
	Type  Type
	Value uint64
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: INTEGER, Value: uint64(i.Value)}
}
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: BOOLEAN, Value: value}
}
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(s.Value))
	if err != nil {
		return HashKey{Type: STRING, Value: rand.Uint64()}
	}

	return HashKey{Type: STRING, Value: h.Sum64()}
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() Type {
	return HASH
}

func (h *Hash) Inspect() string {
	var out bytes.Buffer

	var pairs []string
	for _, pair := range h.Pairs {
		pairs = append(pairs, pair.Key.Inspect()+":"+pair.Value.Inspect())
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() Type {
	return QUOTE
}

func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
