package models

import (
	"time"
)

type Pledge struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UserID    uint
	User      User `gorm:"association_autoupdate:false"`
	ProjectID uint
	Project   Project `gorm:"association_autoupdate:false"`
	Amount    int
	Category  string    `gorm:"default:'Educational Assistance'" json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
