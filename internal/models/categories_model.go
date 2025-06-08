package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID          int       `bun:",pk,autoincrement" json:"id"`
	Name        string    `bun:",notnull,unique" json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	ParentID    *int      `json:"parent_id,omitempty"`
	IsActive    bool      `bun:",default:true" json:"is_active"`
	SortOrder   int       `json:"sort_order"`
	CreatedAt   time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt   time.Time `bun:",default:current_timestamp" json:"updated_at"`

	Parent        *Category  `bun:"rel:belongs-to,join:parent_id=id" json:"parent,omitempty"`
	Subcategories []Category `bun:"rel:has-many,join:id=parent_id" json:"subcategories,omitempty"`
}
