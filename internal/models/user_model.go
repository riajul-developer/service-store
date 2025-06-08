package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement" json:"id"`
	Name      string    `bun:",notnull" json:"name"`
	Email     string    `bun:",unique,notnull" json:"email"`
	Password  string    `bun:",notnull" json:"-"`
	CreatedAt time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",default:current_timestamp" json:"updated_at"`
}
