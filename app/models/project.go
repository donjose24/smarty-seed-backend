package models

import (
	"time"
)

type Project struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `json:"title"`
	Goal        int       `json:"goal"`
	Beneficiary string    `json:"beneficiary"`
	Description string    `json:description`
	ImageUrl    string    `json:"image_url"`
	Category    string    `gorm:"default:'Educational Assistance'" json:"category"`
	Current     int       `gorm:"-", json:"current"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"column:deleted_at;default:null"`
}
