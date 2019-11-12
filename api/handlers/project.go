package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	projectService "github.com/jmramos02/smarty-seed-backend/app/services/project"
	"strconv"
)

func ListProjects(c *gin.Context) {
	page := 0
	stringOffset := c.Query("page")
	if stringOffset != "" {
		page, _ = strconv.Atoi(stringOffset)
	}

	db, _ := c.Get("db")
	if dbObj, success := db.(*gorm.DB); success {
		projects := projectService.List(page, dbObj)

		c.JSON(200, gin.H{
			"data": projects,
		})
	}
}

func ShowProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(404, gin.H{
			"error": "Resource Not found",
		})
		return
	}
	db, _ := c.Get("db")
	if dbObj, success := db.(*gorm.DB); success {
		project := projectService.Show(id, dbObj)
		c.JSON(200, gin.H{
			"data": project,
		})
	}
}
