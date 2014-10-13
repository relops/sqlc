package meta

type TypeInfo struct {
	Prefix  string
	Literal string
}

var Types = []TypeInfo{
	TypeInfo{Prefix: "String", Literal: "string"},
	TypeInfo{Prefix: "Int", Literal: "int"},
	TypeInfo{Prefix: "Int64", Literal: "int64"},
	TypeInfo{Prefix: "Time", Literal: "time.Time"},
}

type Function int

const (
	None Function = iota
	Avg
	Max
	Min
	Count
)

type FunctionInfo struct {
	Name string
	Expr string
}

var Funcs = []FunctionInfo{
	FunctionInfo{Name: "Avg", Expr: "AVG(%s)"},
	FunctionInfo{Name: "Max", Expr: "MAX(%s)"},
	FunctionInfo{Name: "Min", Expr: "MIN(%s)"},
	FunctionInfo{Name: "Div", Expr: "%s / %v"},
}
