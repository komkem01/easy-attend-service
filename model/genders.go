package model

import (
	"time"

	"github.com/uptrace/bun"
)

// Genders table structure
type Genders struct {
	bun.BaseModel `bun:"table:genders,alias:g"`

	ID           int       `json:"id" bun:"id,pk,autoincrement"`
	Code         string    `json:"code" bun:"code,notnull,unique"`
	NameTH       string    `json:"name_th" bun:"name_th,notnull"`
	NameEN       string    `json:"name_en" bun:"name_en,notnull"`
	Abbreviation string    `json:"abbreviation" bun:"abbreviation"`
	IsActive     bool      `json:"is_active" bun:"is_active,notnull,default:true"`
	SortOrder    int       `json:"sort_order" bun:"sort_order,default:0"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,notnull,default:now()"`
}

// TableName returns the table name
func (g *Genders) TableName() string {
	return "genders"
}
