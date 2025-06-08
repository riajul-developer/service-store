package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Hub struct {
	bun.BaseModel `bun:"table:hubs"`

	ID            int       `bun:",pk,autoincrement" json:"id"`
	Name          string    `bun:",unique,notnull" json:"name"`
	Description   *string   `json:"description,omitempty"`
	Address       string    `bun:",notnull" json:"address"`
	Latitude      float64   `bun:",notnull" json:"latitude"`
	Longitude     float64   `bun:",notnull" json:"longitude"`
	HubInchargeID int       `bun:",notnull,column:hub_incharge_id" json:"hub_incharge_id"`
	IsActive      bool      `bun:",default:true" json:"is_active"`
	CreatedAt     time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt     time.Time `bun:",default:current_timestamp" json:"updated_at"`

	HubIncharge *User `bun:"rel:belongs-to,join:hub_incharge_id=id" json:"hub_incharge,omitempty"`
}
