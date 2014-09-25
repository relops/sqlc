package sqlc

import (
	"bytes"
	"database/sql"
	"io"
)

type PredicateType int

const (
	EqPredicate PredicateType = iota
	GtPredicate
	GePredicate
	LtPredicate
	LePredicate
	InPredicate
)

type TableLike interface {
	Selectable
	Name() string
	Queryable
}

type Column interface {
	Name() string
}

type ColumnBinding struct {
	Column Column
	Value  interface{}
}

type Condition struct {
	Binding   ColumnBinding
	Predicate PredicateType
}

type SelectFromStep interface {
	From(Selectable) SelectWhereStep
}

type SelectWhereStep interface {
	Renderable
	Selectable
	Where(conditions ...Condition) Query
}

type Renderable interface {
	Render(io.Writer) []interface{}
}

type Queryable interface {
	Columns() []Column
}

type Query interface {
	Renderable
	QueryRow(*sql.DB) (*sql.Row, error)
}

type Selectable interface {
	isSelectable()
}

type selection struct {
	selection  Selectable
	projection []Column
	predicate  []Condition
}

func (s *selection) isSelectable() {}

func (s *selection) Where(c ...Condition) Query {
	s.predicate = c
	return s
}

func Select(c ...Column) SelectFromStep {
	return &selection{projection: c}
}

func (sl *selection) From(s Selectable) SelectWhereStep {
	sl.selection = s
	return sl
}

func (s *selection) QueryRow(db *sql.DB) (*sql.Row, error) {
	var buf bytes.Buffer
	args := s.Render(&buf)
	return db.QueryRow(buf.String(), args...), nil
}
