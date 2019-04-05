package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// var c gin.Context
	router := gin.Default()
	initDB()
	router.GET("/api/:uri", func(c *gin.Context) {
		c.Request.URL.Path = "/" + c.Param("uri")
		println(c.Request.URL.Path)
		println(c.Request.URL.Path)
		router.HandleContext(c)
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("/health", func(c *gin.Context) {
	})

	userGroup := router.Group("/user")
	{
		userGroup.GET("", getUsers)
		userGroup.GET(":userID", getUserbyID)
		userGroup.DELETE(":userID", deleteUserByID)
		userGroup.PUT(":userID/email/:email", modifyUserEmail)
		userGroup.PUT(":userID", modifyUserbyID)
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
