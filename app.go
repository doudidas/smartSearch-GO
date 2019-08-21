package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(location.Default())
	initDB()
	admin := gin.Accounts{"admin": "VMware1!"}

	router.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	adminGroup := router.Group("/api", gin.BasicAuth(admin))
	{
		adminGroup.GET("healthcheck", func(c *gin.Context) {
			err := pingMongo()
			if err != nil {
				response := gin.H{"api": "true", "mongo": "false"}
				c.JSONP(200, response)
			}
			response := gin.H{"api": "true", "mongo": "true"}
			c.JSON(200, response)
		})
		userGroup := adminGroup.Group("user")
		{
			userGroup.GET("", getUsers)
			userGroup.GET(":userID", getUserbyID)
			userGroup.DELETE(":userID", deleteUserByID)
			userGroup.PUT(":userID/email/:email", modifyUserEmail)
			userGroup.PUT(":userID", modifyUserbyID)
			userGroup.POST("", createUser)
		}
		topicGroup := adminGroup.Group("topic")
		{
			topicGroup.GET("", func(c *gin.Context) {})
			topicGroup.GET("/:topicID", func(c *gin.Context) {})
		}
	}
	router.Run(":80")
}
