package sqlc

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/0x6e6562/gosnow"
	"io"
	"strings"
)

var flake, _ = gosnow.Default()

type selection struct {
	selection  Selectable
	projection []Field
	predicate  []Condition
	groups     []Field
	ordering   []Field
	joins      []join
	joinTarget TableLike
	count      bool
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}

func SelectCount() SelectFromStep {
	return &selection{count: true}
}

func (s *selection) IsSelectable() {}

func (s *selection) Where(c ...Condition) Query {
	s.predicate = c
	return s
}

func (sl *selection) From(s Selectable) SelectWhereStep {
	sl.selection = s
	return sl
}

func (s *selection) Join(t TableLike) SelectOnStep {
	s.joinTarget = t
	return s
}

func (s *selection) On(c ...JoinCondition) SelectWhereStep {
	j := join{
		target:   s.joinTarget,
		joinType: Join,
		conds:    c,
	}
	s.joinTarget = nil
	s.joins = append(s.joins, j)
	return s
}

func (sl *selection) GroupBy(f ...Field) SelectHavingStep {
	sl.groups = f
	return sl
}

func (sl *selection) OrderBy(f ...Field) SelectLimitStep {
	sl.ordering = f
	return sl
}

func (s *selection) QueryRow(d Dialect, db *sql.DB) (*sql.Row, error) {
	var buf bytes.Buffer
	args := s.Render(d, &buf)
	return db.QueryRow(buf.String(), args...), nil
}

func (s *selection) String(d Dialect) string {
	return toString(d, s)
}

func (s *selection) Render(d Dialect, w io.Writer) (placeholders []interface{}) {

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
		sub.Render(d, w)
		// TODO Probably shouldn't swallow this error ......
		n, _ := flake.Next()
		fmt.Fprintf(w, ") AS alias_%d", n)
	}

	for _, join := range s.joins {
		fmt.Println("J...")
		conds := len(join.conds)
		switch conds {
		case 1:
			cond := join.conds[0]
			fmt.Fprintf(w, " JOIN %s ON %s", join.target.Name(), renderJoinFragment(cond))
		default:
			fragments := make([]string, conds)
			for i, cond := range join.conds {
				fragments[i] = renderJoinFragment(cond)
			}

			clause := strings.Join(fragments, " AND ")
			fmt.Fprintf(w, " JOIN %s ON (%s)", join.target.Name(), clause)
		}
	}

	if len(s.predicate) > 0 {
		fmt.Fprint(w, " ")
		placeholders = renderWhereClause(alias, s.predicate, d, 0, w)
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

func renderJoinFragment(cond JoinCondition) string {
	lhsAlias := cond.Lhs.Table()
	lhsField := cond.Lhs.Name()
	rhsAlias := cond.Rhs.Table()
	rhsField := cond.Rhs.Name()
	return fmt.Sprintf("%s.%s = %s.%s", lhsAlias, lhsField, rhsAlias, rhsField)
}
