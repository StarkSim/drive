// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DirectoriesColumns holds the columns for the "directories" table.
	DirectoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "is_public", Type: field.TypeBool, Default: true},
		{Name: "parent_id", Type: field.TypeInt64, Nullable: true},
	}
	// DirectoriesTable holds the schema information for the "directories" table.
	DirectoriesTable = &schema.Table{
		Name:       "directories",
		Columns:    DirectoriesColumns,
		PrimaryKey: []*schema.Column{DirectoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "directories_directories_children",
				Columns:    []*schema.Column{DirectoriesColumns[8]},
				RefColumns: []*schema.Column{DirectoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ObjectsColumns holds the columns for the "objects" table.
	ObjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "is_public", Type: field.TypeBool, Default: true},
		{Name: "directory_objects", Type: field.TypeInt64, Nullable: true},
		{Name: "user_id", Type: field.TypeInt64},
	}
	// ObjectsTable holds the schema information for the "objects" table.
	ObjectsTable = &schema.Table{
		Name:       "objects",
		Columns:    ObjectsColumns,
		PrimaryKey: []*schema.Column{ObjectsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "objects_directories_objects",
				Columns:    []*schema.Column{ObjectsColumns[9]},
				RefColumns: []*schema.Column{DirectoriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "objects_users_objects",
				Columns:    []*schema.Column{ObjectsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "created_by", Type: field.TypeInt64, Default: 0},
		{Name: "updated_by", Type: field.TypeInt64, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_phone_deleted_at",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[8], UsersColumns[5]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DirectoriesTable,
		ObjectsTable,
		UsersTable,
	}
)

func init() {
	DirectoriesTable.ForeignKeys[0].RefTable = DirectoriesTable
	ObjectsTable.ForeignKeys[0].RefTable = DirectoriesTable
	ObjectsTable.ForeignKeys[1].RefTable = UsersTable
}
