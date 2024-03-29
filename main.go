package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

const defaultPageValue = 20

func init() {
	setMongoParameters()
}

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.AllowAllOrigins = true
	router.Use(location.Default(), cors.New(config))
	router.GET("/", func(c *gin.Context) { c.String(200, "") })
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	adminGroup := router.Group("/api", gin.BasicAuth(getAdmins()))
	{
		adminGroup.GET("healthcheck", checkHealth)

		userGroup := adminGroup.Group("user")
		{
			userGroup.GET("", getUsers)
			// userGroup.GET("/count", getUserNumber)
			// userGroup.POST("/filter/", getUserWithFilter)
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
		adminGroup.GET("countTopic", getTopicNumber)
		adminGroup.GET("countUser", getUserNumber)

	}
	// Hello
	router.Run(":9000")
}
