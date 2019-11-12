package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/api/handlers"
	"github.com/jmramos02/smarty-seed-backend/app/services"
	"strings"
)

func Initialize(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.Use(addContextMiddleware(db))
	router.LoadHTMLGlob("web/*.html")

	//Service router will be here below
	api := router.Group("/api/v1")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register)
		api.GET("/projects", handlers.ListProjects)
		api.GET("/projects/:id", handlers.ShowProject)
		api.GET("/unionbank/callback", handlers.HandleUnionbankCallback)

		protectedRoutes := api.Group("")
		{
			protectedRoutes.Use(authenticationMiddleware())
			protectedRoutes.GET("/user", handlers.GetUser)
			protectedRoutes.GET("/payments/unionbank", handlers.GenerateUnionbankRedirectString)
		}
	}

	return router
}

//Temporarily. i think we can use context here.
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
		if len(splitToken) != 2 {
			c.JSON(401, gin.H{
				"error": "Invalid Authorization Header",
			})
			c.Abort()
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		fmt.Println(splitToken)
		user, err := services.DecodeUserInfo(reqToken)
		if err != nil {
			c.JSON(401, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
