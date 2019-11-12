package project

import (
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
)

func List(page int, db *gorm.DB) []models.Project {
	var projects []models.Project
	perPage := 10

	//db.Offset(perPage * page).Limit(perPage).Find(&projects)
	db.Raw("SELECT p.*, SUM(pl.amount) as current FROM projects p LEFT JOIN pledges as pl ON p.id=pl.project_id GROUP BY p.id OFFSET ? LIMIT ?", (perPage * page), perPage).Scan(&projects)

	return projects
}

func Show(id int, db *gorm.DB) models.Project {
	var project models.Project

	db.Debug().Raw("SELECT p.*, SUM(pl.amount) as current FROM projects p LEFT JOIN pledges as pl ON p.id=pl.project_id  WHERE p.id = ? GROUP BY p.id", id).Scan(&project)
	return project
}
