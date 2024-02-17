package object

import (
	"bytes"
	"fmt"
	"github.com/digital-codex/monkey/ast"
	"hash/fnv"
	"math/rand"
	"strings"
)

/*****************************************************************************
 *                                INTERFACES                                 *
 *****************************************************************************/

type Object interface {
	Type() Type
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

/*****************************************************************************
 *                                  TYPES                                    *
 *****************************************************************************/

type Type string

const (
	ANY          Type = "ANY"
	NUMBER       Type = "NUMBER"
	BOOLEAN      Type = "BOOLEAN"
	NULL         Type = "NULL"
	RETURN_VALUE Type = "RETURN_VALUE"
	ERROR        Type = "ERROR"
	FUNCTION     Type = "FUNCTION"
	BUILTIN      Type = "BUILTIN"
	STRING       Type = "STRING"
	ARRAY        Type = "ARRAY"
	HASH         Type = "HASH"
	QUOTE        Type = "QUOTE"
	MACRO        Type = "MACRO"
)

func (t Type) String() string {
	return string(t)
}

type Number struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type Null struct{}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.Block
	Env        *Environment
}

type BuiltinFunction func(args ...Object) Object
type Builtin struct {
	Fn BuiltinFunction
}

type String struct {
	Value string
}

type Array struct {
	Elements []Object
}

type HashKey struct {
	Type  Type
	Value uint64
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

type Quote struct {
	Node ast.Node
}

type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.Block
	Env        *Environment
}

/*****************************************************************************
 *                                OBJECTS                                    *
 *****************************************************************************/

func (i *Number) Type() Type {
	return NUMBER
}
func (b *Boolean) Type() Type {
	return BOOLEAN
}
func (n *Null) Type() Type {
	return NULL
}
func (rv *ReturnValue) Type() Type {
	return RETURN_VALUE
}
func (e *Error) Type() Type {
	return ERROR
}
func (f *Function) Type() Type {
	return FUNCTION
}
func (b *Builtin) Type() Type {
	return BUILTIN
}
func (s *String) Type() Type {
	return STRING
}
func (a *Array) Type() Type {
	return ARRAY
}
func (h *Hash) Type() Type {
	return HASH
}
func (q *Quote) Type() Type {
	return QUOTE
}
func (m *Macro) Type() Type {
	return MACRO
}

func (i *Number) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}
func (n *Null) Inspect() string {
	return "null"
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
func (e *Error) Inspect() string {
	return "Error: " + e.Message
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
func (s *String) Inspect() string {
	return s.Value
}
func (b *Builtin) Inspect() string {
	return "native function"
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
func (q *Quote) Inspect() string {
	return "QUOTE(" + q.Node.String() + ")"
}
func (m *Macro) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("macro")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString("{" + m.Body.String() + "}")

	return out.String()
}

/*****************************************************************************
 *                               HASHABLE                                    *
 *****************************************************************************/

func (i *Number) HashKey() HashKey {
	return HashKey{Type: NUMBER, Value: uint64(i.Value)}
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
