package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",notnull"`
	Email     string    `bun:",unique,notnull"`
	Password  string    `bun:",notnull"`
	CreatedAt time.Time `bun:",default:current_timestamp"`
	UpdatedAt time.Time `bun:",default:current_timestamp"`
}
