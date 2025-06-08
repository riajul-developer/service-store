package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID           int64     `bun:",pk,autoincrement" json:"id"`
	Name         string    `bun:",notnull" json:"name"`
	Email        string    `bun:",unique,notnull" json:"email"`
	Phone        string    `bun:",unique" json:"phone"`
	Password     string    `bun:",notnull" json:"-"`
	Role         string    `bun:",notnull" json:"role"`
	Address      string    `json:"address"`
	ProfileImage string    `json:"profile_image"`
	IsActive     bool      `bun:",default:true" json:"is_active"`
	CreatedAt    time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt    time.Time `bun:",default:current_timestamp" json:"updated_at"`
}
