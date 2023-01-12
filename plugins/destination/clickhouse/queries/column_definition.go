package queries

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

type (
	ColumnDefinition struct {
		Name string
		Type string
	}
	ColumnDefinitions []*ColumnDefinition
)

func (c ColumnDefinitions) Get(name string) *ColumnDefinition {
	for _, col := range c {
		if col.Name == name {
			return col
		}
	}

	return nil
}

func getColumnDefinition(column *schema.Column) *ColumnDefinition {
	typ := chType(column.Type)
	if column.Name == schema.CqIDColumn.Name {
		typ = "UUID" // the only non-nullable for now
	}

	return &ColumnDefinition{
		Name: column.Name,
		Type: typ,
	}
}
