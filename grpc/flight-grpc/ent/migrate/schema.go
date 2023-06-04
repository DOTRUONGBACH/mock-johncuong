// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FlightsColumns holds the columns for the "flights" table.
	FlightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "from", Type: field.TypeString},
		{Name: "to", Type: field.TypeString},
		{Name: "departure_date", Type: field.TypeTime},
		{Name: "arrival_date", Type: field.TypeTime},
		{Name: "available_first_slot", Type: field.TypeInt, Default: 30},
		{Name: "available_economy_slot", Type: field.TypeInt, Default: 70},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"Available", "Arrived", "Cancel"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FlightsTable holds the schema information for the "flights" table.
	FlightsTable = &schema.Table{
		Name:       "flights",
		Columns:    FlightsColumns,
		PrimaryKey: []*schema.Column{FlightsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FlightsTable,
	}
)

func init() {
}