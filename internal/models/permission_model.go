package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Permission struct {
	bun.BaseModel `bun:"table:permissions"`

	ID        int    `bun:",pk,autoincrement"`
	Name      string `bun:",unique"`
	Desc      *string
	CreatedAt time.Time `bun:",default:current_timestamp"`
	UpdatedAt time.Time `bun:",default:current_timestamp"`

	RolePermissions []RolePermission `bun:"rel:has-many"`
}
