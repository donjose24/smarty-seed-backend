package project

import (
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
)

func List(page int, db *gorm.DB) []models.Project {
	var projects []models.Project
	perPage := 10

	db.Offset(perPage * page).Limit(perPage).Find(&projects)

	return projects
}
