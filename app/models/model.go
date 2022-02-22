package models

import (
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type CommonTimeStampsField struct {
	CreateAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdateAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}
