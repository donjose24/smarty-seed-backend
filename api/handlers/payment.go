package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/jmramos02/smarty-seed-backend/app/models"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	"github.com/jmramos02/smarty-seed-backend/app/services/unionbank"
	"github.com/jmramos02/smarty-seed-backend/app/utils"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

func GenerateUnionbankRedirectString(c *gin.Context) {
	projectId, _ := strconv.ParseUint(c.Query("project_id"), 10, 32)
	amount, _ := strconv.ParseInt(c.Query("amount"), 10, 32)

	ub := unionbank.GenerateUnionBankURLRequest{
		ProjectID: uint(projectId),
		Amount:    int(amount),
	}

	v := validator.New()
	err := v.Struct(ub)
	var result error

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			result = multierror.Append(result, errors.New(utils.FormatErrors(e.ActualTag(), e.Field(), e.Param())))
		}

		if merr, ok := result.(*multierror.Error); ok {
			errors := utils.ExtractErrorMessages(merr.Errors)
			c.JSON(400, gin.H{
				"error": errors,
			})
			return
		}
	}

	userContext, _ := c.Get("user")
	if user, success := userContext.(models.User); success {
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
