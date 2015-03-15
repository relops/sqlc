package sqlc

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"io"
	"reflect"
	"strings"
	"time"
)

type PredicateType int
type JoinType int
type Dialect int

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
	LeftOuterJoin
	NotJoined
)

const (
	Sqlite Dialect = iota
	MySQL
	Postgres
)

type Aliasable interface {
	Alias() string
	MaybeAlias() string
}

type TableLike interface {
	Selectable
	Name() string
	As(string) Selectable
	Queryable
}

type FieldFunction struct {
	Child *FieldFunction
	Name  string
	Expr  string
	Args  []interface{}
}

type Field interface {
	Aliasable
	Functional
	Name() string
	Type() reflect.Type
	As(string) Field
	Function() FieldFunction
}

type TableField interface {
	Field
	Parent() Selectable
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
	Join(Selectable) SelectOnStep
	LeftOuterJoin(Selectable) SelectOnStep
}

type SelectOnStep interface {
	On(...JoinCondition) SelectWhereStep
	Query
}

type SelectWhereStep interface {
	Query
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

type InsertResultStep interface {
	Renderable
	Fetch(Dialect, *sql.DB) (*sql.Row, error)
}

type InsertSetMoreStep interface {
	Executable
	InsertSetStep
	Returning(TableField) InsertResultStep
}

type UpdateSetMoreStep interface {
	Executable
	UpdateSetStep
	Where(conditions ...Condition) Executable
}

type Renderable interface {
	Render(Dialect, io.Writer) []interface{}
	String(Dialect) string
}

type Queryable interface {
	Fields() []Field
}

type Query interface {
	Renderable
	Selectable
	Query(Dialect, *sql.DB) (*sql.Rows, error)
	QueryRow(Dialect, *sql.DB) (*sql.Row, error)
}

type Executable interface {
	Renderable
	Exec(Dialect, *sql.DB) (sql.Result, error)
}

type Selectable interface {
	Aliasable
	Reflectable
	IsSelectable()
}

type JoinCondition struct {
	Lhs, Rhs  TableField
	Predicate PredicateType
}

type join struct {
	target   Selectable
	joinType JoinType
	conds    []JoinCondition
}

type update struct {
	table     TableLike
	bindings  []TableFieldBinding
	predicate []Condition
}

func Update(t TableLike) UpdateSetStep {
	return &update{table: t}
}

func (u *update) Where(c ...Condition) Executable {
	u.predicate = c
	return u
}

func (u *update) Set(f TableField, v interface{}) UpdateSetMoreStep {
	binding := TableFieldBinding{Field: f, Value: v}
	u.bindings = append(u.bindings, binding)
	return u
}

func (u *update) Exec(d Dialect, db *sql.DB) (sql.Result, error) {
	return exec(d, u, db)
}

func exec(d Dialect, r Renderable, db *sql.DB) (sql.Result, error) {
	var buf bytes.Buffer
	args := r.Render(d, &buf)
	return db.Exec(buf.String(), args...)
}

func Qualified(parts ...string) string {
	tmp := []string{}
	for _, part := range parts {
		if part != "" {
			tmp = append(tmp, part)
		}
	}
	return strings.Join(tmp, ".")
}

type NullableBlob struct {
	Inet  []byte
	Valid bool // Valid is true if Inet is not NULL
}

// Scan implements the Scanner interface.
func (self *NullableBlob) Scan(value interface{}) error {
	self.Inet, self.Valid = value.([]byte)
	return nil
}

// Value implements the driver Valuer interface.
func (self NullableBlob) Value() (driver.Value, error) {
	if !self.Valid {
		return nil, nil
	}
	return self.Inet, nil
}

type NullableDate struct {
	Date  time.Time
	Valid bool // Valid is true if Date is not NULL
}

// Scan implements the Scanner interface.
func (self *NullableDate) Scan(value interface{}) error {
	self.Date, self.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (self NullableDate) Value() (driver.Value, error) {
	if !self.Valid {
		return nil, nil
	}
	return self.Date, nil
}

type NullableTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (self *NullableTime) Scan(value interface{}) error {
	self.Time, self.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (self NullableTime) Value() (driver.Value, error) {
	if !self.Valid {
		return nil, nil
	}
	return self.Time, nil
}

type NullableDatetime struct {
	Datetime time.Time
	Valid    bool // Valid is true if Datetime is not NULL
}

// Scan implements the Scanner interface.
func (self *NullableDatetime) Scan(value interface{}) error {
	self.Datetime, self.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (self NullableDatetime) Value() (driver.Value, error) {
	if !self.Valid {
		return nil, nil
	}
	return self.Datetime, nil
}
