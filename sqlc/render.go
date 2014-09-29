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

func (s *selection) String() string {
	return toString(s)
}

func (s *selection) Render(w io.Writer) (placeholders []interface{}) {

	var alias string

	// TODO This type switch is used twice, consider refactoring
	switch sub := s.selection.(type) {
	case TableLike:
		alias = sub.Name()
	}

	fmt.Fprint(w, "SELECT ")

	if len(s.projection) == 0 {
		fmt.Fprint(w, "*")
	} else {
		colClause := columnClause(alias, s.projection)
		fmt.Fprint(w, colClause)
	}

	fmt.Fprintf(w, " FROM ")

	switch sub := s.selection.(type) {
	case TableLike:
		fmt.Fprint(w, sub.Name())
	case *selection:
		fmt.Fprint(w, "(")
		sub.Render(w)
		fmt.Fprint(w, ")")
	}

	// TODO Support more than one join
	if len(s.joins) == 1 {
		join := s.joins[0]
		lhsAlias := join.conds[0].Lhs.Table()
		lhsField := join.conds[0].Lhs.Name()
		rhsAlias := join.conds[0].Rhs.Table()
		rhsField := join.conds[0].Rhs.Name()
		fmt.Fprintf(w, " JOIN %s ON %s.%s = %s.%s", join.target.Name(), lhsAlias, lhsField, rhsAlias, rhsField)
	}

	if len(s.predicate) > 0 {
		fmt.Fprint(w, " ")
		placeholders = renderWhereClause(alias, s.predicate, w)
	} else {
		placeholders = []interface{}{}
	}

	if (len(s.groups)) > 0 {
		fmt.Fprint(w, " GROUP BY ")
		colClause := columnClause(alias, s.groups)
		fmt.Fprint(w, colClause)
	}

	// TODO eliminate copy and paste
	if (len(s.ordering)) > 0 {
		fmt.Fprint(w, " ORDER BY ")
		colClause := columnClause(alias, s.ordering)
		fmt.Fprint(w, colClause)
	}

	return placeholders
}

func columnClause(alias string, cols []Field) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		colFragments[i] = fmt.Sprintf("%s.%s", alias, col.Name())
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
