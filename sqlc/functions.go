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

type groupConcat struct {
	stringField
}

// This may indicate that the rendering pipeline needs to get adjusted so that things like can be less stateful
func (g *groupConcat) OrderBy(f Field) *groupConcat {
	al := resolveParentAlias(f.Alias(), f)
	g.stringField.fun.Expr = "GROUP_CONCAT(%s ORDER BY %s.%s ASC)"
	g.stringField.fun.Args = append(g.stringField.fun.Args, al, f.Name())
	return g
}

func (g *groupConcat) Separator(s string) *groupConcat {
	if len(g.stringField.fun.Args) > 0 {
		g.stringField.fun.Expr = "GROUP_CONCAT(%s ORDER BY %s.%s ASC SEPARATOR '%s')" // TODO ASC is hard coded
	} else {
		g.stringField.fun.Expr = "GROUP_CONCAT(%s, '%s')" // TODO sqlite specific (i.e. no SEPARATOR keyword)
	}

	g.stringField.fun.Args = append(g.stringField.fun.Args, s)
	return g
}

func GroupConcat(field Field) *groupConcat {

	var s Selectable
	if tf, ok := field.(TableField); ok {
		s = tf.Parent()
	}

	return &groupConcat{
		stringField: stringField{
			name:      field.Name(),
			selection: s,
			fun: FieldFunction{
				Name: "GroupConcat",
				Expr: "GROUP_CONCAT(%s)",
			},
		},
	}
}
