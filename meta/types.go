package meta

import (
	"reflect"
)

type TypeInfo struct {
	Prefix  string
	Literal string
	Type    reflect.Type
}

var Types = []TypeInfo{
	TypeInfo{Prefix: "Bool", Literal: "bool"},
	TypeInfo{Prefix: "Date", Literal: "time.Time"},     // TODO(shutej): test
	TypeInfo{Prefix: "Datetime", Literal: "time.Time"}, // TODO(shutej): test
	TypeInfo{Prefix: "Float32", Literal: "float32"},
	TypeInfo{Prefix: "Float64", Literal: "float64"},
	TypeInfo{Prefix: "Int", Literal: "int"},
	TypeInfo{Prefix: "Int64", Literal: "int64"},
	TypeInfo{Prefix: "NullBool", Literal: "sql.NullBool"},
	TypeInfo{Prefix: "NullDate", Literal: "NullableDate"},         // TODO(shutej): test
	TypeInfo{Prefix: "NullDatetime", Literal: "NullableDatetime"}, // TODO(shutej): test
	TypeInfo{Prefix: "NullFloat32", Literal: "sql.NullFloat64"},   // TODO(shutej): test
	TypeInfo{Prefix: "NullFloat64", Literal: "sql.NullFloat64"},
	TypeInfo{Prefix: "NullInt", Literal: "sql.NullInt64"}, // TODO(shutej): test
	TypeInfo{Prefix: "NullInt64", Literal: "sql.NullInt64"},
	TypeInfo{Prefix: "NullString", Literal: "sql.NullString"},
	TypeInfo{Prefix: "NullTime", Literal: "NullableTime"}, // TODO(shutej): test
	TypeInfo{Prefix: "String", Literal: "string"},
	TypeInfo{Prefix: "Time", Literal: "time.Time"}, // TODO(shutej): test
}

type FunctionInfo struct {
	Name string
	Expr string
}

var Funcs = []FunctionInfo{
	FunctionInfo{Name: "Avg", Expr: "AVG(%s)"},
	FunctionInfo{Name: "Max", Expr: "MAX(%s)"},
	FunctionInfo{Name: "Min", Expr: "MIN(%s)"},
	FunctionInfo{Name: "Ceil", Expr: "CEIL(%s)"},
	FunctionInfo{Name: "Div", Expr: "%s / %v"},
	FunctionInfo{Name: "Cast", Expr: "CAST(%s AS %s)"},
	FunctionInfo{Name: "Md5", Expr: "MD5(%s)"},
	FunctionInfo{Name: "Lower", Expr: "LOWER(%s)"},
	FunctionInfo{Name: "Hex", Expr: "HEX(%s)"},
	FunctionInfo{Name: "Substr2", Expr: "SUBSTR(%s, %v)"},
	FunctionInfo{Name: "Substr3", Expr: "SUBSTR(%s, %v, %v)"},
}
