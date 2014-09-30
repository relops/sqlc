package sqlc

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var predicateTypes = map[PredicateType]string{
	EqPredicate: "=",
	GtPredicate: ">",
	GePredicate: ">=",
	LtPredicate: "<",
	LePredicate: "<=",
}

func (u *update) String() string {
	return toString(u)
}

func (u *update) Render(w io.Writer) (placeholders []interface{}) {
	fmt.Fprintf(w, "UPDATE %s SET ", u.table.Name())

	setFragments := make([]string, len(u.bindings))
	setValues := make([]interface{}, len(u.bindings))

	for i, binding := range u.bindings {
		col := binding.Field.Name()
		setFragments[i] = fmt.Sprintf("%s = ?", col)
		setValues[i] = binding.Value
	}

	setClause := strings.Join(setFragments, ", ")
	fmt.Fprint(w, setClause)

	fmt.Fprint(w, " ")

	whereValues := renderWhereClause(u.table.Name(), u.predicate, w)

	placeholders = append(setValues, whereValues...)

	return placeholders
}

func (i *insert) String() string {
	return toString(i)
}

func (i *insert) Render(w io.Writer) (placeholders []interface{}) {
	fmt.Fprintf(w, "INSERT INTO %s (", i.table.Name())
	colFragments := make([]string, len(i.bindings))
	for i, binding := range i.bindings {
		col := binding.Field.Name()
		colFragments[i] = col
	}
	colClause := strings.Join(colFragments, ", ")
	fmt.Fprint(w, colClause)

	fmt.Fprint(w, ") VALUES (")

	placeHolderFragments := make([]string, len(i.bindings))
	values := make([]interface{}, len(i.bindings))
	for i, binding := range i.bindings {
		placeHolderFragments[i] = "?"
		values[i] = binding.Value
	}

	placeHolderClause := strings.Join(placeHolderFragments, ",")
	fmt.Fprint(w, placeHolderClause)
	fmt.Fprint(w, ")")

	return values
}

func columnClause(alias string, cols []Field) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		var f string
		switch col.Function() {
		case Max:
			f = fmt.Sprintf("MAX(%s.%s)", alias, col.Name())
		case Min:
			f = fmt.Sprintf("MIN(%s.%s)", alias, col.Name())
		default:
			f = fmt.Sprintf("%s.%s", alias, col.Name())
		}

		colFragments[i] = f

	}
	return strings.Join(colFragments, ", ")
}

func renderWhereClause(alias string, conds []Condition, w io.Writer) []interface{} {
	fmt.Fprint(w, "WHERE ")

	whereFragments := make([]string, len(conds))
	values := make([]interface{}, len(conds))

	for i, condition := range conds {
		col := condition.Binding.Field.Name()
		pred := condition.Predicate
		whereFragments[i] = fmt.Sprintf("%s.%s %s ?", alias, col, predicateTypes[pred])
		values[i] = condition.Binding.Value
	}

	whereClause := strings.Join(whereFragments, " AND ")
	fmt.Fprint(w, whereClause)

	return values
}

func toString(r Renderable) string {
	var buf bytes.Buffer
	r.Render(&buf)
	return buf.String()
}
