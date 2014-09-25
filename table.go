package sqlc

type table struct {
	name string
	cols []Column
}

func (t table) isSelectable() {}

func (t table) Name() string {
	return t.name
}

func (t table) Columns() []Column {
	return t.cols
}

func Table(name string) TableLike {
	return table{name: name}
}
