package sqlc

type table struct {
	name string
	cols []Column
}

func (t *table) TableName() string {
	return t.name
}

func (t *table) ColumnDefinitions() []Column {
	return t.cols
}

func Table(name string) TableLike {
	return &table{name: name}
}
