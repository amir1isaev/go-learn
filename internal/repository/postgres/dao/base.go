package dao

import (
	"github.com/google/uuid"
	"gopkg.in/guregu/null.v4"
	"time"
)

var AppSchema = "app"
var ServiceGroupsTableName = AppSchema + "." + "service_groups"
var GroupUsersTableName = AppSchema + "." + "group_users"
var UserPermissionsTableName = AppSchema + "." + "user_permissions"
var GroupPermissionsTableName = AppSchema + "." + "group_permissions"

type BaseModel struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt null.Time `json:"deleted_at" db:"deleted_at"`
}

var columnForCreateBase = []string{
	"id",
	"created_at",
	"updated_at",
	"deleted_at",
}
