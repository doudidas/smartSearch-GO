package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// User structure from Database
type User struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Picture   struct {
		Large     string `json:"large"`
		Medium    string `json:"medium"`
		Thumbnail string `json:"thumbnail"`
	} `json:"picture"`
	Username string `json:"username"`
}

func getUserbyID(c *gin.Context) {
	user, err := getUserFromFirebaseByID(c, c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, user)
}

func deleteUserByID(c *gin.Context) {
	err := deleteUserFromFirebase(c, c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(500, err)
	}
	c.Status(201)
	c.Done()
}
func modifyUserEmail(c *gin.Context) {
	// 	// client, err := getClient(c)
	// 	// if err != nil {
	// 	// 	c.AbortWithStatusJSON(500, err.Error())
	// 	// }
	// 	// defer client.Disconnect(c)
	// 	// collection := getUserCollection(client)

	// 	// value := bson.M{
	// 	// 	"$set": bson.M{
	// 	// 		"email": c.Param("email"),
	// 	// 	},
	// 	// }
	// 	// objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	// 	// if err != nil {
	// 	// 	c.AbortWithStatusJSON(500, err.Error())
	// 	// }
	// 	// filter := bson.M{"_id": objectID}
	// 	// fmt.Println(value)
	// 	// var output bson.M
	// 	// collection.FindOneAndUpdate(c, filter, value).Decode(&output)
	// 	// if err != nil {
	// 	// 	c.AbortWithStatusJSON(500, err.Error())
	// 	// }
	// 	// c.JSON(200, output)
}
func modifyUserByID(c *gin.Context) {
	var newUser User
	c.ShouldBindJSON(&newUser)
	err := modifyUserOnFirebase(c, c.Param("userID"), newUser)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.Done()
}

func createUsers(c *gin.Context) {
	var values []User
	err := c.ShouldBindJSON(&values)
	if err != nil {
		c.AbortWithStatusJSON(500, "Please provide an array of JSON files")
	}
	result, err := addUsersOnFirebase(c, values)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, result)
}

func getUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	pageSize := c.DefaultQuery("size", strconv.Itoa(defaultPageValue))
	pageAsNumber, _ := strconv.Atoi(page)
	sizeAsNumber, _ := strconv.Atoi(pageSize)
	result, err := getUsersFromFirebase(c, pageAsNumber, sizeAsNumber)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, result)
}
