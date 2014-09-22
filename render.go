package sqlc

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

func renderSelect(ctx *Context, buf *bytes.Buffer) error {

	fmt.Fprint(buf, "SELECT ")

	var colClause string
	if len(ctx.Columns) == 0 {
		colClause = columnClause(ctx.Table.ColumnDefinitions())

	} else {
		colClause = columnClause(ctx.Columns)
	}

	fmt.Fprint(buf, colClause)

	if ctx.Table == nil {
		return errors.New("No table supplied")
	}

	fmt.Fprintf(buf, " FROM %s", ctx.Table.TableName())

	return nil
}

func columnClause(cols []Column) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		colFragments[i] = col.ColumnName()
	}
	return strings.Join(colFragments, ", ")
}
