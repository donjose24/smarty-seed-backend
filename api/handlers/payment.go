package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	"github.com/jmramos02/smarty-seed-backend/app/services/unionbank"
)

func GenerateUnionbankRedirectString(c *gin.Context) {
	var ub unionbank.GenerateUrlRequest
	userContext, _ := c.Get("user")
	if user, success := userContext.(models.User); success {
		c.Bind(&ub)
		pledge := models.Pledge{
			Amount:    ub.Amount,
			ProjectID: ub.ProjectID,
			User:      user,
		}
		state := services.EncodePledge(pledge)

		url := unionbank.GenerateUnionbankString(ub, state)

		c.JSON(200, gin.H{
			"data": url,
		})
	}
}
