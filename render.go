package sqlc

import (
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

func (s *selection) Render(w io.Writer) []interface{} {
	fmt.Fprint(w, "SELECT ")

	if len(s.projection) == 0 {
		fmt.Fprint(w, "*")
	} else {
		colClause := columnClause(s.projection)
		fmt.Fprint(w, colClause)
	}

	fmt.Fprintf(w, " FROM ")

	switch sub := s.selection.(type) {
	case table:
		fmt.Fprint(w, sub.name)
	case *selection:
		fmt.Fprint(w, "(")
		sub.Render(w)
		fmt.Fprint(w, ")")
	}

	if len(s.predicate) > 0 {
		fmt.Fprint(w, " ")
		return renderWhereClause(s.predicate, w)
	} else {
		return []interface{}{}
	}
}

func columnClause(cols []Column) string {
	colFragments := make([]string, len(cols))
	for i, col := range cols {
		colFragments[i] = col.Name()
	}
	return strings.Join(colFragments, ", ")
}

func renderWhereClause(conds []Condition, w io.Writer) []interface{} {
	fmt.Fprint(w, "WHERE ")

	whereFragments := make([]string, len(conds))
	values := make([]interface{}, len(conds))

	for i, condition := range conds {
		col := condition.Binding.Column.Name()
		pred := condition.Predicate
		whereFragments[i] = fmt.Sprintf("%s %s ?", col, predicateTypes[pred])
		values[i] = condition.Binding.Value
	}

	whereClause := strings.Join(whereFragments, " AND ")
	fmt.Fprint(w, whereClause)

	return values
}
