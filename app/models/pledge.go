package models

import (
	"time"
)

type Pledge struct {
	ID        uint `gorm:"primary_key" json:"id"`
	UserID    uint
	User      User
	ProjectId uint
	Project   Project
	Amount    int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
