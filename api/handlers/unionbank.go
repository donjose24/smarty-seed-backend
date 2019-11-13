package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	pledgeService "github.com/jmramos02/smarty-seed-backend/app/services/pledge"
	"github.com/jmramos02/smarty-seed-backend/app/services/unionbank"
)

type SuccessData struct {
	ProjectID uint
}

func HandleUnionbankCallback(c *gin.Context) {
	errors := c.Query("error")
	db, _ := c.Get("db")

	if dbObj, success := db.(*gorm.DB); success {

		// authentication has errors
		if errors != "" {
			c.Writer.WriteString("An error has occured. Please try again later")
			return
		}

		code := c.Query("code")
		authorization, err := unionbank.GetAuthorizationCode(code)
		if err != nil {
			c.Writer.WriteString("An error has occured. Please try again later")
			return
		}

		if authorization.Error != "" {
			c.Writer.WriteString("An error has occured. Please try again later: " + authorization.Error)
			return
		}

		state := c.Query("state")
		pledge, err := services.DecodePledge(state)

		if err != nil {
			c.Writer.WriteString(fmt.Sprintf("Request expired. Please try again. %v", err.Error))
			return
		}

		_, err = unionbank.ExecutePayment(pledge.Amount, authorization.AccessToken)
		pledgeDb := models.Pledge{
			UserID:    pledge.UserID,
			Amount:    pledge.Amount,
			ProjectID: pledge.ProjectID,
		}

		successData := SuccessData{
			ProjectID: pledge.ProjectID,
		}

		pledgeService.Create(pledgeDb, dbObj)

		c.HTML(200, "success.html", successData)
	}
}
