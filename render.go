package sqlc

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

var predicateTypes = map[PredicateType]string{
	EqPredicate: "=",
	GtPredicate: ">",
	GePredicate: ">=",
	LtPredicate: "<",
	LePredicate: "<=",
}

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

	if ctx.hasConditions() {
		fmt.Fprint(buf, " ")
		renderWhereClause(ctx, buf)
	}

	return nil
}

func columnClause(cols []Column) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		colFragments[i] = col.ColumnName()
	}
	return strings.Join(colFragments, ", ")
}

func renderWhereClause(ctx *Context, buf *bytes.Buffer) {
	fmt.Fprint(buf, "WHERE ")

	whereFragments := make([]string, len(ctx.Conditions))
	for i, condition := range ctx.Conditions {
		col := condition.Binding.Column.ColumnName()
		pred := condition.Predicate
		whereFragments[i] = fmt.Sprintf("%s %s ?", col, predicateTypes[pred])
	}

	whereClause := strings.Join(whereFragments, " AND ")
	fmt.Fprint(buf, whereClause)
}
