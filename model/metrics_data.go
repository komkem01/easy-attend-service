package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// MetricsData table structure
type MetricsData struct {
	bun.BaseModel `bun:"table:metrics_data,alias:md"`

	ID         uuid.UUID  `json:"id" bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	MetricType string     `json:"metric_type" bun:"metric_type,notnull,type:metric_type"`
	MetricName string     `json:"metric_name" bun:"metric_name,notnull"`
	Value      float64    `json:"value" bun:"value,notnull"`
	Unit       *string    `json:"unit" bun:"unit"`
	Tags       *string    `json:"tags" bun:"tags,type:jsonb"`
	Context    *string    `json:"context" bun:"context,type:jsonb"`
	Timestamp  time.Time  `json:"timestamp" bun:"timestamp,notnull,default:now()"`
	CreatedAt  time.Time  `json:"created_at" bun:"created_at,notnull,default:now()"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty" bun:"deleted_at,soft_delete"`
}

// TableName returns the table name
func (md *MetricsData) TableName() string {
	return "metrics_data"
}
