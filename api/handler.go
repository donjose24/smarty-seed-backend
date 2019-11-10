package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	//Service router will be here below

	//Run the webapp

	return router
}
