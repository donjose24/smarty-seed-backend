package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
)

func Login(c *gin.Context) {
	var request services.LoginRequest
	c.Bind(&request)
	c.Get("db")
	db, _ := c.Get("db")
	if dbObj, success := db.(*gorm.DB); success {
		response, err := services.Login(request, dbObj)
		if err != nil {
			if merr, ok := err.(*multierror.Error); ok {
				errors := utils.ExtractErrorMessages(merr.Errors)
				c.JSON(400, gin.H{
					"errors": errors,
				})
				return
			}
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": response,
		})

		return
	}
}
