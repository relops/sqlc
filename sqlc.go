package sqlc

import (
	"bytes"
)

type Context struct {
	Table   TableLike
	Columns []Column
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

type SelectFromStep interface {
	From(table TableLike) SelectWhereStep
}

type SelectWhereStep interface {
	X() string
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

func (c *Context) X() string {
	return ""
}

func (c *Context) Render() (string, error) {
	var buf bytes.Buffer
	if err := renderSelect(c, &buf); err != nil {
		return "", err
	} else {
		return buf.String(), nil
	}
}
