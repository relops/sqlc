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
	Bindings   []ColumnBinding
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

func (c *Context) RenderSQL() (string, error) {
	var buf bytes.Buffer
	if err := renderSelect(c, &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}

func (c *Context) Build() (stmt string, placeHolders []interface{}, err error) {
	stmt, err = c.RenderSQL()
	if err != nil {
		return stmt, nil, err
	}

	bindings := len(c.Bindings) // TODO check whether this is nil
	conditions := 0

	if c.Conditions != nil {
		conditions = len(c.Conditions)
	}

	placeHolders = make([]interface{}, bindings+conditions)

	for i, bind := range c.Bindings {
		placeHolders[i] = bind.Value
	}

	if c.Conditions != nil {
		for i, cond := range c.Conditions {
			placeHolders[i+bindings] = cond.Binding.Value
		}
	}

	c.Dispose()

	return stmt, placeHolders, nil
}

func (c *Context) Dispose() {
	c.Columns = nil
	c.Table = nil
	c.Bindings = nil
	c.Conditions = nil
}

func (c *Context) hasConditions() bool {
	return len(c.Conditions) > 0
}
