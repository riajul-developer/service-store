package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Vendor struct {
	bun.BaseModel `bun:"table:vendors"`

	ID             int       `bun:",pk,autoincrement" json:"id"`
	UserID         int       `bun:",notnull,unique" json:"user_id"`
	HubID          int       `bun:",notnull" json:"hub_id"`
	BusinessName   string    `bun:",notnull" json:"business_name"`
	BusinessType   string    `bun:",notnull" json:"business_type"`
	TradeLicense   string    `bun:",unique" json:"trade_license"`
	NIDNo          string    `bun:",notnull,unique" json:"nid_no"`
	BankAccount    string    `json:"bank_account"`
	Status         string    `bun:",type:text,notnull" json:"status"`
	CommissionRate float64   `bun:",notnull" json:"commission_rate"`
	CreatedAt      time.Time `bun:",default:current_timestamp" json:"created_at"`
	UpdatedAt      time.Time `bun:",default:current_timestamp" json:"updated_at"`

	User *User `bun:"rel:belongs-to,join:user_id=id" json:"user,omitempty"`
	Hub  *Hub  `bun:"rel:belongs-to,join:hub_id=id" json:"hub,omitempty"`
}
