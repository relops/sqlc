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

type Field interface {
	Name() string
}

type FieldBinding struct {
	Field Field
	Value interface{}
}

type Condition struct {
	Binding   FieldBinding
	Predicate PredicateType
}

type SelectFromStep interface {
	From(Selectable) SelectWhereStep
}

type SelectWhereStep interface {
	Renderable
	Selectable
	SelectGroupByStep
	Where(conditions ...Condition) Query
}

type SelectGroupByStep interface {
	GroupBy(...Field) SelectHavingStep
}

type SelectHavingStep interface {
	SelectOrderByStep
	Query
}

type SelectOrderByStep interface {
	OrderBy(...Field) SelectLimitStep
}

type SelectLimitStep interface {
	Query
}

type Renderable interface {
	Render(io.Writer) []interface{}
	String() string
}

type Queryable interface {
	Fields() []Field
}

type Query interface {
	Renderable
	Selectable
	QueryRow(*sql.DB) (*sql.Row, error)
}

type Selectable interface {
	isSelectable()
}

type selection struct {
	selection  Selectable
	projection []Field
	predicate  []Condition
	groups     []Field
	ordering   []Field
}

func (s *selection) isSelectable() {}

func (s *selection) Where(c ...Condition) Query {
	s.predicate = c
	return s
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}

func (sl *selection) From(s Selectable) SelectWhereStep {
	sl.selection = s
	return sl
}

func (sl *selection) GroupBy(f ...Field) SelectHavingStep {
	sl.groups = f
	return sl
}

func (sl *selection) OrderBy(f ...Field) SelectLimitStep {
	sl.ordering = f
	return sl
}

func (s *selection) QueryRow(db *sql.DB) (*sql.Row, error) {
	var buf bytes.Buffer
	args := s.Render(&buf)
	return db.QueryRow(buf.String(), args...), nil
}
