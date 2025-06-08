package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Permission struct {
	bun.BaseModel `bun:"table:permissions"`

	ID        int       `bun:",pk,autoincrement" json:"id"`
	Name      string    `bun:",unique" json:"name"`
	Desc      *string   `json:"desc,omitempty"`
	CreatedAt time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",default:current_timestamp" json:"updated_at"`

	RolePermissions []RolePermission `bun:"rel:has-many" json:"role_permissions,omitempty"`
}
