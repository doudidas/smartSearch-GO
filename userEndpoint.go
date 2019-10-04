package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
	// client, err := getClient(c)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, err.Error())
	// }
	// defer client.Disconnect(c)
	// collection := getUserCollection(client)

	// value := bson.M{
	// 	"$set": bson.M{
	// 		"email": c.Param("email"),
	// 	},
	// }
	// objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, err.Error())
	// }
	// filter := bson.M{"_id": objectID}
	// fmt.Println(value)
	// var output bson.M
	// collection.FindOneAndUpdate(c, filter, value).Decode(&output)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, err.Error())
	// }
	// c.JSON(200, output)
}
func modifyUserByID(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getUserCollection(client)

	var value bson.M
	err = c.ShouldBindJSON(&value)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	update := bson.M{
		"$set": value,
	}

	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := bson.M{"_id": objectID}
	var output bson.M

	opt := options.FindOneAndUpdateOptions{}
	opt.SetReturnDocument(options.After)

	collection.FindOneAndUpdate(c, filter, update, &opt).Decode(&output)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, output)
}

func createUsers(c *gin.Context) {
	var values []map[string]interface{}
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
func getUserCollection(client *mongo.Client) *mongo.Collection {
	return getDatabase(client).Collection("userCollection")
}
