package models

import (
	"time"

	"github.com/uptrace/bun"
)

type RolePermission struct {
	bun.BaseModel `bun:"table:role_permissions"`

	ID           int `bun:",pk,autoincrement"`
	RoleID       int
	PermissionID int
	CreatedAt    time.Time `bun:",default:current_timestamp"`
	UpdatedAt    time.Time `bun:",default:current_timestamp"`

	Role       *Role       `bun:"rel:belongs-to,join:role_id=id"`
	Permission *Permission `bun:"rel:belongs-to,join:permission_id=id"`
}
