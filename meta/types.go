package meta

type TypeInfo struct {
	Prefix  string
	Literal string
}

var Types = []TypeInfo{
	TypeInfo{Prefix: "String", Literal: "string"},
	TypeInfo{Prefix: "Int", Literal: "int"},
	TypeInfo{Prefix: "Int64", Literal: "int64"},
	TypeInfo{Prefix: "Float32", Literal: "float32"},
	TypeInfo{Prefix: "Float64", Literal: "float64"},
	TypeInfo{Prefix: "Time", Literal: "time.Time"},
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
