package model

import (
	"time"

	"github.com/uptrace/bun"
)

// Prefixes table structure (คำนำหน้าชื่อ)
type Prefixes struct {
	bun.BaseModel `bun:"table:prefixes,alias:p"`

	ID           int        `json:"id" bun:"id,pk,autoincrement"`
	Code         string     `json:"code" bun:"code,notnull,unique"`
	NameTH       string     `json:"name_th" bun:"name_th,notnull"`
	NameEN       string     `json:"name_en" bun:"name_en,notnull"`
	Abbreviation string     `json:"abbreviation" bun:"abbreviation"`
	GenderCode   *string    `json:"gender_code" bun:"gender_code"` // เชื่อมกับ gender (optional)
	IsActive     bool       `json:"is_active" bun:"is_active,notnull,default:true"`
	SortOrder    int        `json:"sort_order" bun:"sort_order,default:0"`
	CreatedAt    time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time  `json:"updated_at" bun:"updated_at,notnull,default:now()"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`

	// Relations
	Gender *Genders `json:"gender,omitempty" bun:"rel:belongs-to,join:gender_code=code"`
}

// TableName returns the table name
func (p *Prefixes) TableName() string {
	return "prefixes"
}
