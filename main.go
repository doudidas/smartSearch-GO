package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

const defaultPageValue = 20

func init() {
	setMongoParameters()
}

func main() {
	router := gin.Default()
	router.Use(location.Default())

	adminGroup := router.Group("/api", gin.BasicAuth(getAdmins()))
	{
		adminGroup.GET("", func(c *gin.Context) {
			c.Request.URL.Path = "/swagger/index.html"
			router.HandleContext(c)
		})
		adminGroup.GET("healthcheck", checkHealth)
		userGroup := adminGroup.Group("user")
		{
			userGroup.GET("", getUsers)
			userGroup.GET(":userID", getUserbyID)
			userGroup.DELETE(":userID", deleteUserByID)
			userGroup.PUT(":userID/email/:email", modifyUserEmail)
			userGroup.PUT(":userID", modifyUserByID)
			userGroup.POST("", createUser)
		}
		topicGroup := adminGroup.Group("topic")
		{
			topicGroup.GET("", getTopics)
			topicGroup.GET(":topicID", getTopicbyID)
			topicGroup.DELETE(":topicID", deleteTopicByID)
			topicGroup.PUT(":topicID", modifyTopicByID)
			topicGroup.POST("", createTopic)
		}
	}
	router.GET("/", func(c *gin.Context) { c.String(200, "") })
	router.GET("ping", func(c *gin.Context) { c.String(200, "pong") })
	router.Run(":9000")
}
