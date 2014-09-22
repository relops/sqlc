package sqlc

type VarcharColumn struct {
	Name string
}

func (c *VarcharColumn) ColumnName() string {
	return c.Name
}

func Field(name string) Column {
	return &VarcharColumn{Name: name}
}
