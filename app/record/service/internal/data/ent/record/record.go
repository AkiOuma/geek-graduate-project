// Code generated by entc, DO NOT EDIT.

package record

import (
	"time"
)

const (
	// Label holds the string label denoting the record type in the database.
	Label = "record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUser holds the string denoting the user field in the database.
	FieldUser = "user"
	// FieldDay holds the string denoting the day field in the database.
	FieldDay = "day"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the record in the database.
	Table = "reoord"
)

// Columns holds all SQL columns for record fields.
var Columns = []string{
	FieldID,
	FieldUser,
	FieldDay,
	FieldType,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)