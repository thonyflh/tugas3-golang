package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseStruct struct {
	ID        uint           `json:"id" gorm:"primaryKey;type:not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
