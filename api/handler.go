package api

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/api/handlers"
	"strings"
)

func Initialize(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.Use(addContextMiddleware(db))

	//Service router will be here below
	api := router.Group("/api/v1")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		api.GET("/projects", handlers.ListProjects)
	}

	return router
}

//Temporarily. this is bad practice
func addContextMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		ctx := context.WithValue(c.Request.Context(), "GinContextKey", c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func authenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		reqToken = splitToken[1]
		if len(splitToken) != 2 {
			c.JSON(401, gin.H{
				"error": "Invalid Authorization Header",
			})
		}
	}
}
