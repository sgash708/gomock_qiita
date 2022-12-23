package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	UUID      string         `json:"uuid" gorm:"column:uuid"`
	Name      string         `json:"name" gorm:"column:name"`
	CreatedAt time.Time      `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at"`
}
