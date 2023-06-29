package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primaryKey" `
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Deleted   bool           `json:"deleted" gorm:"index;default:false"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index" `
}
