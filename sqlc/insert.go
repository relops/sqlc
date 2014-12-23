package sqlc

import (
	"bytes"
	"database/sql"
)

type insert struct {
	table     TableLike
	bindings  []TableFieldBinding
	returning TableField
}

func InsertInto(t TableLike) InsertSetStep {
	return &insert{table: t}
}

func (i *insert) Exec(d Dialect, db *sql.DB) (sql.Result, error) {
	return exec(d, i, db)
}

func (i *insert) Returning(f TableField) InsertResultStep {
	i.returning = f
	return i
}

func (i *insert) Fetch(d Dialect, db *sql.DB) (*sql.Row, error) {
	var buf bytes.Buffer
	args := i.Render(d, &buf)
	return db.QueryRow(buf.String(), args...), nil
}

func (i *insert) set(f TableField, v interface{}) InsertSetMoreStep {
	binding := TableFieldBinding{Field: f, Value: v}
	i.bindings = append(i.bindings, binding)
	return i
}
