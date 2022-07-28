// Code generated by ent, DO NOT EDIT.

package ent

import (
	"drive/ent/directory"
	"drive/ent/object"
	"drive/ent/schema"
	"drive/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	directoryMixin := schema.Directory{}.Mixin()
	directoryMixinFields0 := directoryMixin[0].Fields()
	_ = directoryMixinFields0
	directoryFields := schema.Directory{}.Fields()
	_ = directoryFields
	// directoryDescCreatedBy is the schema descriptor for created_by field.
	directoryDescCreatedBy := directoryMixinFields0[1].Descriptor()
	// directory.DefaultCreatedBy holds the default value on creation for the created_by field.
	directory.DefaultCreatedBy = directoryDescCreatedBy.Default.(int64)
	// directoryDescUpdatedBy is the schema descriptor for updated_by field.
	directoryDescUpdatedBy := directoryMixinFields0[2].Descriptor()
	// directory.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	directory.DefaultUpdatedBy = directoryDescUpdatedBy.Default.(int64)
	// directoryDescCreatedAt is the schema descriptor for created_at field.
	directoryDescCreatedAt := directoryMixinFields0[3].Descriptor()
	// directory.DefaultCreatedAt holds the default value on creation for the created_at field.
	directory.DefaultCreatedAt = directoryDescCreatedAt.Default.(func() time.Time)
	// directoryDescUpdatedAt is the schema descriptor for updated_at field.
	directoryDescUpdatedAt := directoryMixinFields0[4].Descriptor()
	// directory.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	directory.DefaultUpdatedAt = directoryDescUpdatedAt.Default.(func() time.Time)
	// directory.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	directory.UpdateDefaultUpdatedAt = directoryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// directoryDescDeletedAt is the schema descriptor for deleted_at field.
	directoryDescDeletedAt := directoryMixinFields0[5].Descriptor()
	// directory.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	directory.DefaultDeletedAt = directoryDescDeletedAt.Default.(time.Time)
	// directoryDescIsPublic is the schema descriptor for is_public field.
	directoryDescIsPublic := directoryFields[1].Descriptor()
	// directory.DefaultIsPublic holds the default value on creation for the is_public field.
	directory.DefaultIsPublic = directoryDescIsPublic.Default.(bool)
	// directoryDescID is the schema descriptor for id field.
	directoryDescID := directoryMixinFields0[0].Descriptor()
	// directory.DefaultID holds the default value on creation for the id field.
	directory.DefaultID = directoryDescID.Default.(func() int64)
	objectMixin := schema.Object{}.Mixin()
	objectMixinFields0 := objectMixin[0].Fields()
	_ = objectMixinFields0
	objectFields := schema.Object{}.Fields()
	_ = objectFields
	// objectDescCreatedBy is the schema descriptor for created_by field.
	objectDescCreatedBy := objectMixinFields0[1].Descriptor()
	// object.DefaultCreatedBy holds the default value on creation for the created_by field.
	object.DefaultCreatedBy = objectDescCreatedBy.Default.(int64)
	// objectDescUpdatedBy is the schema descriptor for updated_by field.
	objectDescUpdatedBy := objectMixinFields0[2].Descriptor()
	// object.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	object.DefaultUpdatedBy = objectDescUpdatedBy.Default.(int64)
	// objectDescCreatedAt is the schema descriptor for created_at field.
	objectDescCreatedAt := objectMixinFields0[3].Descriptor()
	// object.DefaultCreatedAt holds the default value on creation for the created_at field.
	object.DefaultCreatedAt = objectDescCreatedAt.Default.(func() time.Time)
	// objectDescUpdatedAt is the schema descriptor for updated_at field.
	objectDescUpdatedAt := objectMixinFields0[4].Descriptor()
	// object.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	object.DefaultUpdatedAt = objectDescUpdatedAt.Default.(func() time.Time)
	// object.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	object.UpdateDefaultUpdatedAt = objectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// objectDescDeletedAt is the schema descriptor for deleted_at field.
	objectDescDeletedAt := objectMixinFields0[5].Descriptor()
	// object.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	object.DefaultDeletedAt = objectDescDeletedAt.Default.(time.Time)
	// objectDescIsPublic is the schema descriptor for is_public field.
	objectDescIsPublic := objectFields[1].Descriptor()
	// object.DefaultIsPublic holds the default value on creation for the is_public field.
	object.DefaultIsPublic = objectDescIsPublic.Default.(bool)
	// objectDescID is the schema descriptor for id field.
	objectDescID := objectMixinFields0[0].Descriptor()
	// object.DefaultID holds the default value on creation for the id field.
	object.DefaultID = objectDescID.Default.(func() int64)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedBy is the schema descriptor for created_by field.
	userDescCreatedBy := userMixinFields0[1].Descriptor()
	// user.DefaultCreatedBy holds the default value on creation for the created_by field.
	user.DefaultCreatedBy = userDescCreatedBy.Default.(int64)
	// userDescUpdatedBy is the schema descriptor for updated_by field.
	userDescUpdatedBy := userMixinFields0[2].Descriptor()
	// user.DefaultUpdatedBy holds the default value on creation for the updated_by field.
	user.DefaultUpdatedBy = userDescUpdatedBy.Default.(int64)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescDeletedAt is the schema descriptor for deleted_at field.
	userDescDeletedAt := userMixinFields0[5].Descriptor()
	// user.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	user.DefaultDeletedAt = userDescDeletedAt.Default.(time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields0[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() int64)
}
