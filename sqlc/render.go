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

func (u *update) String(d Dialect) string {
	return toString(d, u)
}

func (u *update) Render(d Dialect, w io.Writer) (placeholders []interface{}) {
	fmt.Fprintf(w, "UPDATE %s SET ", u.table.Name())

	setFragments := make([]string, len(u.bindings))
	setValues := make([]interface{}, len(u.bindings))

	for i, binding := range u.bindings {
		col := binding.Field.Name()
		setFragments[i] = fmt.Sprintf("%s = %s", col, d.renderPlaceholder(i+1))
		setValues[i] = binding.Value
	}

	setClause := strings.Join(setFragments, ", ")
	fmt.Fprint(w, setClause)

	fmt.Fprint(w, " ")

	paramCount := len(u.bindings)
	whereValues := renderWhereClause(u.table.Name(), u.predicate, d, paramCount, w)

	placeholders = append(setValues, whereValues...)

	return placeholders
}

func (i *insert) String(d Dialect) string {
	return toString(d, i)
}

func (i *insert) Render(d Dialect, w io.Writer) (placeholders []interface{}) {
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
		placeHolderFragments[i] = d.renderPlaceholder(i + 1)
		values[i] = binding.Value
	}

	placeHolderClause := strings.Join(placeHolderFragments, ",")
	fmt.Fprint(w, placeHolderClause)
	fmt.Fprint(w, ")")

	return values
}

func resolveAlias(alias string, col Field) string {
	if alias == "" {
		if tabCol, ok := col.(TableField); ok {
			return tabCol.Table()
		}

		return ""
	} else {
		return alias
	}
}

func columnClause(alias string, cols []Field) string {

	colFragments := make([]string, len(cols))
	for i, col := range cols {
		al := resolveAlias(alias, col)
		var f string
		switch col.Function() {
		case Avg:
			f = fmt.Sprintf("AVG(%s.%s)", al, col.Name())
		case Max:
			f = fmt.Sprintf("MAX(%s.%s)", al, col.Name())
		case Min:
			f = fmt.Sprintf("MIN(%s.%s)", al, col.Name())
		default:
			f = fmt.Sprintf("%s.%s", al, col.Name())
		}

		colFragments[i] = f

	}
	return strings.Join(colFragments, ", ")
}

func renderWhereClause(alias string, conds []Condition, d Dialect, paramCount int, w io.Writer) []interface{} {
	fmt.Fprint(w, "WHERE ")

	whereFragments := make([]string, len(conds))
	values := make([]interface{}, len(conds))

	for i, condition := range conds {
		field := condition.Binding.Field
		al := resolveAlias(alias, field)
		col := field.Name()
		pred := condition.Predicate
		placeHolder := d.renderPlaceholder(i + paramCount + 1)
		whereFragments[i] = fmt.Sprintf("%s.%s %s %s", al, col, predicateTypes[pred], placeHolder)
		values[i] = condition.Binding.Value
	}

	whereClause := strings.Join(whereFragments, " AND ")
	fmt.Fprint(w, whereClause)

	return values
}

func (d Dialect) renderPlaceholder(n int) string {
	switch d {
	case Postgres:
		return fmt.Sprintf("$%d", n)
	default:
		return "?"
	}
}

func toString(d Dialect, r Renderable) string {
	var buf bytes.Buffer
	r.Render(d, &buf)
	return buf.String()
}
