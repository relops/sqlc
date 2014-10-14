package sqlc

func Count() IntField {
	return &intField{name: "*", fun: FieldFunction{Name: "Count", Expr: "COUNT(*)"}}
}
