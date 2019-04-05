package main

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(location.Default())
	initDB()

	// router.GET("/api/:uri", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/" + c.Param("uri")
	// 	username, password, ok := c.Request.BasicAuth()
	// 	customLog("username: " + username + ", password: " + password + ",ok: " + strconv.FormatBool(ok))
	// 	router.HandleContext(c)
	// })
	// router.DELETE("/api/:uri", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/" + c.Param("uri")
	// 	username, password, ok := c.Request.BasicAuth()
	// 	customLog("username: " + username + ", password: " + password + ",ok: " + strconv.FormatBool(ok))
	// 	router.HandleContext(c)
	// })
	// router.PUT("/api/:uri", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/" + c.Param("uri")
	// 	username, password, ok := c.Request.BasicAuth()
	// 	customLog("username: " + username + ", password: " + password + ",ok: " + strconv.FormatBool(ok))
	// 	router.HandleContext(c)
	// })
	// router.POST("/api/:uri", func(c *gin.Context) {
	// 	c.Request.URL.Path = "/" + c.Param("uri")
	// 	username, password, ok := c.Request.BasicAuth()
	// 	customLog("username: " + username + ", password: " + password + ",ok: " + strconv.FormatBool(ok))
	// 	router.HandleContext(c)
	// })
	router.GET("api/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	router.GET("api/healthcheck", func(c *gin.Context) {
		err := pingMongo()
		if err != nil {
			c.AbortWithError(500, err)
		}
		response := gin.H{"api": "true",
			"mongo": "true"}
		c.JSON(200, response)
	})

	userGroup := router.Group("api/user")
	{
		userGroup.GET("", getUsers)
		userGroup.GET(":userID", getUserbyID)
		userGroup.DELETE(":userID", deleteUserByID)
		userGroup.PUT(":userID/email/:email", modifyUserEmail)
		userGroup.PUT(":userID", modifyUserbyID)
		userGroup.POST("", createUser)
	}
	topicGroup := router.Group("api/topic")
	{
		topicGroup.GET("", func(c *gin.Context) {
		})
		topicGroup.GET("/:id", func(c *gin.Context) {
		})
	}
	destinationGroup := router.Group("api/destination")
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
