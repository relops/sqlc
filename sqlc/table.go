package sqlc

type table struct {
	name   string
	fields []Field
	alias  string
}

func (t table) IsSelectable() {}

func (t table) Name() string {
	return t.name
}

func (t table) Fields() []Field {
	return t.fields
}

func (t table) As(alias string) Selectable {
	t.alias = alias
	return t
}

func (t table) Alias() string {
	return t.alias
}

func Table(name string) TableLike {
	return table{name: name}
}
