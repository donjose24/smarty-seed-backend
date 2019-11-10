package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/smarty-seed-backend/services"
)

func Login(c *gin.Context) {
	var request services.LoginRequest
	c.Bind(&request)
}
