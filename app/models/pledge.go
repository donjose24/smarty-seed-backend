package models

import (
	"time"
)

type Pledge struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	UserID    uint      `json:"-"`
	User      User      `gorm:"association_autoupdate:false" json:"user"`
	ProjectID uint      `json:"-"`
	Project   Project   `gorm:"association_autoupdate:false" json:"project"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
