package object

import "fmt"

type Type string

const (
	INTEGER = "INTEGER"
	BOOLEAN = "BOOLEAN"
	NULLOBJ = "NULL"
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

type NULL struct{}

func (n *NULL) Type() Type {
	return NULLOBJ
}

func (n *NULL) Inspect() string {
	return "null"
}
