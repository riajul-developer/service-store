package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Name      string    `bun:",unique,notnull" json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",default:current_timestamp" json:"updated_at"`

	Users           []User           `bun:"rel:has-many" json:"users,omitempty"`
	RolePermissions []RolePermission `bun:"rel:has-many" json:"role_permissions,omitempty"`
}
