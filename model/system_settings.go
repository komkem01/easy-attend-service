package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// SystemSettings table structure
type SystemSettings struct {
	bun.BaseModel `bun:"table:system_settings,alias:ss"`

	ID           uuid.UUID `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	SettingKey   string    `json:"setting_key" bun:"setting_key,notnull,unique"`
	SettingName  string    `json:"setting_name" bun:"setting_name,notnull"`
	Description  *string   `json:"description" bun:"description"`
	DataType     string    `json:"data_type" bun:"data_type,notnull,type:data_type_setting"`
	Value        *string   `json:"value" bun:"value"`
	DefaultValue *string   `json:"default_value" bun:"default_value"`
	IsPublic     bool      `json:"is_public" bun:"is_public,notnull,default:false"`
	IsEditable   bool      `json:"is_editable" bun:"is_editable,notnull,default:true"`
	Category     *string   `json:"category" bun:"category"`
	SortOrder    int       `json:"sort_order" bun:"sort_order,default:0"`
	CreatedAt    time.Time `json:"created_at" bun:"created_at,notnull,default:now()"`
	UpdatedAt    time.Time `json:"updated_at" bun:"updated_at,notnull,default:now()"`
}

// TableName returns the table name
func (ss *SystemSettings) TableName() string {
	return "system_settings"
}
