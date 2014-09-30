package sqlc

import (
	"fmt"
	"io"
)

func SelectCount() SelectFromStep {
	return &selection{count: true}
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

	if s.count {
		fmt.Fprint(w, "COUNT(*)")
	} else {
		if len(s.projection) == 0 {
			fmt.Fprint(w, "*")
		} else {
			colClause := columnClause(alias, s.projection)
			fmt.Fprint(w, colClause)
		}
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
