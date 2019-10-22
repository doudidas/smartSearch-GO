package main

import (
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
)

const defaultPageValue = 20

func init() {
	setMongoParameters()
}

func main() {
	router := gin.Default()
	router.Use(location.Default())

	if _, err := os.Stat("./favicon.ico"); err == nil {
		router.Use(favicon.New("./favicon.ico"))
	}

	admin := getAdmins()

	router.GET("ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	adminGroup := router.Group("/api", gin.BasicAuth(admin))
	{
		adminGroup.GET("", func(c *gin.Context) {
			c.Request.URL.Path = "/swagger/"
			router.HandleContext(c)
		})
		adminGroup.GET("healthcheck", func(c *gin.Context) {
			var response gin.H
			client, err := getClient(c)
			defer client.Disconnect(c)
			if err != nil || client == nil {
				response = gin.H{"api": "true", "mongo": "true"}
			} else {
				response = gin.H{"api": "true", "mongo": "true"}
			}
			c.JSON(200, response)
		})
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

	router.Run(":9000")
}
