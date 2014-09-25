package sqlc

import (
	"bytes"
	"database/sql"
	"io"
)

type PredicateType int
type JoinType int

const (
	EqPredicate PredicateType = iota
	GtPredicate
	GePredicate
	LtPredicate
	LePredicate
	InPredicate
)

const (
	Join JoinType = iota
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

type SelectJoinStep interface {
	Join(TableLike) SelectOnStep
}

type SelectOnStep interface {
	On(...Condition) SelectWhereStep
	Query
}

type SelectWhereStep interface {
	Renderable
	Selectable
	SelectGroupByStep
	SelectJoinStep
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
	joins      []join
	joinTarget TableLike
}

type join struct {
	target   TableLike
	joinType JoinType
	conds    []Condition
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

func (s *selection) Join(t TableLike) SelectOnStep {
	s.joinTarget = t
	return s
}

func (s *selection) On(c ...Condition) SelectWhereStep {
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

func (s *selection) QueryRow(db *sql.DB) (*sql.Row, error) {
	var buf bytes.Buffer
	args := s.Render(&buf)
	return db.QueryRow(buf.String(), args...), nil
}
