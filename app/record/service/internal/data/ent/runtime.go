// Code generated by entc, DO NOT EDIT.

package ent

import (
	"clock-in/app/record/service/internal/data/ent/record"
	"clock-in/app/record/service/internal/data/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	recordFields := schema.Record{}.Fields()
	_ = recordFields
	// recordDescUpdatedAt is the schema descriptor for updated_at field.
	recordDescUpdatedAt := recordFields[4].Descriptor()
	// record.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	record.DefaultUpdatedAt = recordDescUpdatedAt.Default.(func() time.Time)
}
