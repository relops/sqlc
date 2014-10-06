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
	joinTarget Selectable
	joinType   JoinType
	count      bool
	aliases    map[string]string
	alias      string
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}

func SelectCount() SelectFromStep {
	return &selection{count: true}
}

func (s *selection) IsSelectable() {}

func (sl *selection) maybeCacheAlias(s Selectable) {
	if t, ok := s.(TableLike); ok {
		sl.cacheAlias(t)
	}
}

func (s *selection) cacheAlias(t TableLike) {
	if len(s.aliases) == 0 {
		s.aliases = make(map[string]string)
	}
	s.aliases[t.Name()] = t.Alias()
}

func (s *selection) Alias() string {
	return s.alias
}

func (s *selection) As(a string) Selectable {
	s.alias = a
	return s
}

func (s *selection) Where(c ...Condition) Query {
	s.predicate = c
	return s
}

func (sl *selection) From(s Selectable) SelectWhereStep {
	sl.selection = s
	sl.maybeCacheAlias(s)
	return sl
}

func (s *selection) Join(t Selectable) SelectOnStep {
	s.joinTarget = t
	s.joinType = Join
	s.maybeCacheAlias(t)
	return s
}

func (s *selection) LeftOuterJoin(t Selectable) SelectOnStep {
	// TODO copy and paste from Join(.)
	s.joinTarget = t
	s.joinType = LeftOuterJoin
	s.maybeCacheAlias(t)
	return s
}

func (s *selection) On(c ...JoinCondition) SelectWhereStep {
	j := join{
		target:   s.joinTarget,
		joinType: s.joinType,
		conds:    c,
	}
	s.joinTarget = nil
	s.joinType = NotJoined
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

func (s *selection) Query(d Dialect, db *sql.DB) (*sql.Rows, error) {
	var buf bytes.Buffer
	args := s.Render(d, &buf)
	return db.Query(buf.String(), args...)
}

func (s *selection) String(d Dialect) string {
	return toString(d, s)
}

func (s *selection) Render(d Dialect, w io.Writer) (placeholders []interface{}) {

	alias := ""
	if al, ok := s.selection.(Aliasable); ok {
		if al.Alias() != "" {
			alias = al.Alias()
		}
	}

	fmt.Fprint(w, "SELECT ")

	if s.count {
		fmt.Fprint(w, "COUNT(*)")
	} else {
		if len(s.projection) == 0 {
			fmt.Fprint(w, "*")
		} else {
			//colAlias := ""
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
		fmt.Fprint(w, ")")
		if alias == "" {
			// TODO Probably shouldn't swallow this error ......
			n, _ := flake.Next()
			alias = fmt.Sprintf("alias_%d", n)
		}
	}

	if alias != "" {
		fmt.Fprintf(w, " AS %s", alias)
	}

	for _, join := range s.joins {

		var joinString string
		switch join.joinType {
		case LeftOuterJoin:
			joinString = "LEFT OUTER JOIN"
		case Join:
			joinString = "JOIN"
		}

		conds := len(join.conds)
		switch conds {
		case 1:
			cond := join.conds[0]
			var al string
			var aliased bool
			if t, ok := join.target.(TableLike); ok {
				al, aliased = renderTableAlias(t)
			} else {
				al = join.target.Alias()
				aliased = false
			}

			if aliased {
				fmt.Fprintf(w, " %s %s ON %s", joinString, al, s.renderJoinFragment(cond))
			} else {
				fmt.Fprintf(w, " %s %s ON %s", joinString, al, s.renderJoinFragment(cond))
			}
		default:
			// TODO copy and paste
			var al string
			var aliased bool
			if t, ok := join.target.(TableLike); ok {
				al, aliased = renderTableAlias(t)
			} else {
				al = join.target.Alias()
				aliased = false
			}

			fragments := make([]string, conds)
			for i, cond := range join.conds {
				if aliased {
					fragments[i] = s.renderJoinFragment(cond)
				} else {
					fragments[i] = s.renderJoinFragment(cond)
				}

			}

			clause := strings.Join(fragments, " AND ")

			fmt.Fprintf(w, " %s %s ON (%s)", joinString, al, clause)
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

func renderTableAlias(t TableLike) (string, bool) {
	if t.Alias() != "" {
		return fmt.Sprintf("%s AS %s", t.Name(), t.Alias()), true
	} else {
		return t.Name(), false
	}
}

func renderFieldAlias(alias string, f TableField) (string, bool) {
	if alias != "" {
		return fmt.Sprintf("%s.%s", alias, f.Name()), true
	} else if f.Alias() != "" {
		return fmt.Sprintf("%s.%s", f.Alias(), f.Name()), true
	} else {
		return fmt.Sprintf("%s.%s", f.Table(), f.Name()), false
	}
}

func (s *selection) renderJoinFragment(cond JoinCondition) string {
	lhsAlias, _ := renderFieldAlias(s.aliases[cond.Lhs.Table()], cond.Lhs)
	rhsAlias, _ := renderFieldAlias(s.aliases[cond.Rhs.Table()], cond.Rhs)
	return fmt.Sprintf("%s = %s", lhsAlias, rhsAlias)
}
