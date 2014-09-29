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

func (d *deletion) Exec(db *sql.DB) (sql.Result, error) {
	return exec(d, db)
}

func (d *deletion) String() string {
	return toString(d)
}

func (d *deletion) Render(w io.Writer) (placeholders []interface{}) {

	fmt.Fprintf(w, "DELETE FROM %s", d.table.Name())

	if len(d.predicate) > 0 {
		fmt.Fprint(w, " ")
		placeholders = renderWhereClause(d.table.Name(), d.predicate, w)
	}

	return placeholders
}
