package sqlc

func Count() IntField {
	return &intField{name: "*", fun: FieldFunction{Name: "Count", Expr: "COUNT(*)"}}
}

func Trunc(field TimeField, format string) TimeField {
	return &timeField{
		name:      field.Name(), // TODO a lot of this is quite boilerplate, consider auto-generating
		selection: field.Parent(),
		fun: FieldFunction{
			Name: "Trunc",
			Expr: "STRFTIME('%[2]s', %[1]s)", // TODO this is SQLite specific and won't work elsewhere - need to feed in dialect
			Args: []interface{}{format},
		},
	}
}
