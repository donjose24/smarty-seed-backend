package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
)

func Register(c *gin.Context) {
	var request services.RegisterRequest
	c.Bind(&request)
	db, _ := c.Get("db")
	if dbObj, success := db.(*gorm.DB); success {
		response, err := services.Register(request, dbObj)
		if err != nil {
			if merr, ok := err.(*multierror.Error); ok {
				errors := utils.ExtractErrorMessages(merr.Errors)
				c.JSON(400, gin.H{
					"error": errors,
				})
				return
			}
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"data": response,
		})
	}

	return
}
