package models

import (
	"time"

	"github.com/uptrace/bun"
)

type RolePermission struct {
	bun.BaseModel `bun:"table:role_permissions"`

	ID           int       `bun:",pk,autoincrement" json:"id"`
	RoleID       int       `json:"role_id"`
	PermissionID int       `json:"permission_id"`
	CreatedAt    time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `bun:",default:current_timestamp" json:"updated_at"`

	Role       *Role       `bun:"rel:belongs-to,join:role_id=id" json:"role,omitempty"`
	Permission *Permission `bun:"rel:belongs-to,join:permission_id=id" json:"permission,omitempty"`
}
