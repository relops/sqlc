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
	TypeInfo{Prefix: "String", Literal: "string"},
	TypeInfo{Prefix: "Bool", Literal: "bool"},
	TypeInfo{Prefix: "Int", Literal: "int"},
	TypeInfo{Prefix: "Int64", Literal: "int64"},
	TypeInfo{Prefix: "Float32", Literal: "float32"},
	TypeInfo{Prefix: "Float64", Literal: "float64"},
	TypeInfo{Prefix: "Time", Literal: "time.Time"},

	TypeInfo{Prefix: "NullString", Literal: "sql.NullString"},
	TypeInfo{Prefix: "NullBool", Literal: "sql.NullBool"},
	TypeInfo{Prefix: "NullInt", Literal: "sql.NullInt64"}, // TODO(shutej): test
	TypeInfo{Prefix: "NullInt64", Literal: "sql.NullInt64"},
	TypeInfo{Prefix: "NullFloat32", Literal: "sql.NullFloat64"}, // TODO(shutej): test
	TypeInfo{Prefix: "NullFloat64", Literal: "sql.NullFloat64"},
	TypeInfo{Prefix: "NullTime", Literal: "NullableTime"},
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
}
