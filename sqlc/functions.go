package sqlc

func Count() IntField {
	return &intField{name: "*", fun: "Count", expr: "COUNT(*)"}
}
