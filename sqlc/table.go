package sqlc

type table struct {
	name   string
	fields []Field
}

func (t table) IsSelectable() {}

func (t table) Name() string {
	return t.name
}

func (t table) Fields() []Field {
	return t.fields
}

func Table(name string) TableLike {
	return table{name: name}
}
