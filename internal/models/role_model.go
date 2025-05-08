package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Role struct {
	bun.BaseModel `bun:"table:roles"`

	ID          int    `bun:",pk,autoincrement"`
	Name        string `bun:",unique"`
	Description *string
	CreatedAt   time.Time `bun:",default:current_timestamp"`
	UpdatedAt   time.Time `bun:",default:current_timestamp"`

	Users           []User           `bun:"rel:has-many"`
	RolePermissions []RolePermission `bun:"rel:has-many"`
}
