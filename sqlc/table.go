package sqlc

type table struct {
	name     string
	fields   []Field
	fieldMap map[string]Field
	alias    string
}

func (t table) IsSelectable() {}

func (t table) Name() string {
	return t.name
}

func (t table) Fields() []Field {
	return t.fields
}

func (t table) As(alias string) Selectable {
	return table{name: t.name, fields: t.fields, alias: alias, fieldMap: make(map[string]Field)}
}

func (t table) Alias() string {
	return t.alias
}

func (t table) MaybeAlias() string {
	if t.alias == "" {
		return t.name
	} else {
		return t.alias
	}
}

func (t table) StringField(name string) StringField {
	return &stringField{name: name, table: t}
}

func Table(name string) TableLike {
	return table{name: name, fieldMap: make(map[string]Field)}
}
