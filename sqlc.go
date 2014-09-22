package sqlc

import (
	"bytes"
)

type PredicateType int

const (
	EqPredicate PredicateType = iota
	GtPredicate
	GePredicate
	LtPredicate
	LePredicate
	InPredicate
)

type Context struct {
	Table      TableLike
	Columns    []Column
	Conditions []Condition
}

func NewContext() *Context {
	return &Context{}
}

type TableLike interface {
	TableName() string
	ColumnDefinitions() []Column
}

type Column interface {
	ColumnName() string
}

type ColumnBinding struct {
	Column Column
	Value  interface{}
}

type Condition struct {
	Binding   ColumnBinding
	Predicate PredicateType
}

type SelectFromStep interface {
	From(table TableLike) SelectWhereStep
}

type SelectWhereStep interface {
	Where(conditions ...Condition) Queryable
}

type Queryable interface {
}

func (c *Context) Select(cols ...Column) SelectFromStep {
	c.Columns = cols
	//c.Operation = ReadOperation
	return c
}

func (c *Context) From(t TableLike) SelectWhereStep {
	c.Table = t
	return c
}

func (c *Context) Where(cond ...Condition) Queryable {
	c.Conditions = cond
	return c
}

func (c *Context) Render() (string, error) {
	var buf bytes.Buffer
	if err := renderSelect(c, &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func (c *Context) hasConditions() bool {
	return len(c.Conditions) > 0
}
