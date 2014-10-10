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
