// Code generated by ent, DO NOT EDIT.

package object

import (
	"time"
)

const (
	// Label holds the string label denoting the object type in the database.
	Label = "object"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldIsPublic holds the string denoting the is_public field in the database.
	FieldIsPublic = "is_public"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeDirectory holds the string denoting the directory edge name in mutations.
	EdgeDirectory = "directory"
	// Table holds the table name of the object in the database.
	Table = "objects"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "objects"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_objects"
	// DirectoryTable is the table that holds the directory relation/edge.
	DirectoryTable = "objects"
	// DirectoryInverseTable is the table name for the Directory entity.
	// It exists in this package in order to avoid circular dependency with the "directory" package.
	DirectoryInverseTable = "directories"
	// DirectoryColumn is the table column denoting the directory relation/edge.
	DirectoryColumn = "directory_objects"
)

// Columns holds all SQL columns for object fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldURL,
	FieldIsPublic,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "objects"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"directory_objects",
	"user_objects",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt time.Time
	// DefaultIsPublic holds the default value on creation for the "is_public" field.
	DefaultIsPublic bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
