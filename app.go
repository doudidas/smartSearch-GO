package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB("mongo")

	router := gin.Default()

	router.GET("/ping", ping)
	router.GET("/health", func(c *gin.Context) {
		// c := startSession()
		// response := gin.H{
		// 	"database": c.LiveServers(),
		// }
		// c.JSON(200, response)
	})

	userGroup := router.Group("/user")
	{
		userGroup.GET("", getUser)
		userGroup.GET(":userID", getUserbyID)
		userGroup.DELETE(":userID", deleteUserByID)
		userGroup.PUT(":id", modifyUserbyID)
		userGroup.POST("", createUser)
	}
	topicGroup := router.Group("/topic")
	{
		topicGroup.GET("", func(c *gin.Context) {
		})
		topicGroup.GET("/:id", func(c *gin.Context) {
		})
	}
	destinationGroup := router.Group("/destination")
	{
		destinationGroup.GET("", func(c *gin.Context) {
		})
		destinationGroup.GET("random", func(c *gin.Context) {
		})
		// destinationGroup.GET("/:id", func(c *gin.Context) {
		// })
		destinationGroup.GET("user/:id", func(c *gin.Context) {
		})
	}
	router.Run(":9000")
}
