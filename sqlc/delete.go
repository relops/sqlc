package sqlc

import (
	"database/sql"
	"fmt"
	"io"
)

type DeleteWhereStep interface {
	Executable
	Where(...Condition) Executable
}

type deletion struct {
	table     TableLike
	predicate []Condition
}

func Delete(t TableLike) DeleteWhereStep {
	return &deletion{table: t}
}

func (d *deletion) Where(c ...Condition) Executable {
	d.predicate = c
	return d
}

func (d *deletion) Exec(dl Dialect, db *sql.DB) (sql.Result, error) {
	return exec(dl, d, db)
}

func (d *deletion) String(dl Dialect) string {
	return toString(dl, d)
}

func (d *deletion) Render(dl Dialect, w io.Writer) (placeholders []interface{}) {

	fmt.Fprintf(w, "DELETE FROM %s", d.table.Name())

	if len(d.predicate) > 0 {
		fmt.Fprint(w, " ")
		placeholders = renderWhereClause(d.table.Name(), d.predicate, dl, 0, w)
	}

	return placeholders
}
