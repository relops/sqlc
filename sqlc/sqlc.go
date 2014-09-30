package sqlc

import (
	"bytes"
	"database/sql"
	"io"
)

type PredicateType int
type JoinType int
type Dialect int
type Function int

const (
	None Function = iota
	Max
	Min
)

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

const (
	Sqlite Dialect = iota
	MySQL
)

type TableLike interface {
	Selectable
	Name() string
	Queryable
}

type Field interface {
	Name() string
	Min() Field
	Max() Field
	Function() Function
}

type TableField interface {
	Field
	Table() string
}

type FieldBinding struct {
	Field Field
	Value interface{}
}

type TableFieldBinding struct {
	Field TableField
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
	On(...JoinCondition) SelectWhereStep
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

type InsertSetMoreStep interface {
	Executable
	InsertSetStep
}

type UpdateSetMoreStep interface {
	Executable
	UpdateSetStep
	Where(conditions ...Condition) Executable
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

type Executable interface {
	Renderable
	Exec(db *sql.DB) (sql.Result, error)
}

type Selectable interface {
	IsSelectable()
}

type JoinCondition struct {
	Lhs, Rhs  TableField
	Predicate PredicateType
}

type join struct {
	target   TableLike
	joinType JoinType
	conds    []JoinCondition
}

type insert struct {
	table    TableLike
	bindings []TableFieldBinding
}

type update struct {
	table     TableLike
	bindings  []TableFieldBinding
	predicate []Condition
}

func InsertInto(t TableLike) InsertSetStep {
	return &insert{table: t}
}

func Update(t TableLike) UpdateSetStep {
	return &update{table: t}
}

func (u *update) Where(c ...Condition) Executable {
	u.predicate = c
	return u
}

func (u *update) set(f TableField, v interface{}) UpdateSetMoreStep {
	binding := TableFieldBinding{Field: f, Value: v}
	u.bindings = append(u.bindings, binding)
	return u
}

func (i *insert) set(f TableField, v interface{}) InsertSetMoreStep {
	binding := TableFieldBinding{Field: f, Value: v}
	i.bindings = append(i.bindings, binding)
	return i
}

func (s *insert) Exec(db *sql.DB) (sql.Result, error) {
	return exec(s, db)
}

func (u *update) Exec(db *sql.DB) (sql.Result, error) {
	return exec(u, db)
}

func exec(r Renderable, db *sql.DB) (sql.Result, error) {
	var buf bytes.Buffer
	args := r.Render(&buf)
	return db.Exec(buf.String(), args...)
}
