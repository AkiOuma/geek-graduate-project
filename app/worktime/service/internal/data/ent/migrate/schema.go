// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// WorktimeColumns holds the columns for the "worktime" table.
	WorktimeColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user", Type: field.TypeInt64},
		{Name: "day", Type: field.TypeInt64},
		{Name: "minute", Type: field.TypeInt64},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// WorktimeTable holds the schema information for the "worktime" table.
	WorktimeTable = &schema.Table{
		Name:       "worktime",
		Columns:    WorktimeColumns,
		PrimaryKey: []*schema.Column{WorktimeColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "worktime_user_day",
				Unique:  true,
				Columns: []*schema.Column{WorktimeColumns[1], WorktimeColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		WorktimeTable,
	}
)

func init() {
	WorktimeTable.Annotation = &entsql.Annotation{
		Table: "worktime",
	}
}
