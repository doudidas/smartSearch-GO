package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getUserbyID(c *gin.Context) {
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getUserCollection(client)
	customLog(c.Param("userID"))
	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := bson.M{"_id": objectID}
	var result bson.M
	collection.FindOne(c, filter).Decode(&result)
	fmt.Println(result)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}

	c.JSON(200, result)
}

func getUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	pageSize := c.DefaultQuery("size", strconv.Itoa(defaultPageValue))
	findOptions := options.Find()
	pageAsNumber, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	sizeAsNumber, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	findOptions.SetSkip(pageAsNumber * sizeAsNumber)
	findOptions.SetLimit(sizeAsNumber)
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getUserCollection(client)
	cur, err := collection.Find(c, bson.D{{}}, findOptions)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer cur.Close(c)
	var result []bson.M
	for cur.Next(c) {
		var tmp bson.M
		err = cur.Decode(&tmp)
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
		}
		result = append(result, tmp)
	}
	if err := cur.Err(); err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	if result == nil {
		result = []bson.M{}
	}
	c.JSON(200, result)
}
func deleteUserByID(c *gin.Context) {
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getUserCollection(client)

	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid input. Please check format"})
		return
	}
	filter := bson.M{"_id": objectID}
	var result bson.M
	collection.FindOneAndDelete(c, filter).Decode(&result)

	if result == nil {
		message := "Failed to remove user with this ID " + c.Param("userID")
		customLog(message)
		c.AbortWithStatusJSON(500, gin.H{"error": message})
		return
	}
	c.JSON(200, result)

}
func modifyUserEmail(c *gin.Context) {
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getUserCollection(client)

	value := bson.M{
		"$set": bson.M{
			"email": c.Param("email"),
		},
	}
	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	filter := bson.M{"_id": objectID}
	fmt.Println(value)
	var output bson.M
	collection.FindOneAndUpdate(c, filter, value).Decode(&output)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, output)
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

func createUser(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getUserCollection(client)

	var values bson.A
	err = c.ShouldBindJSON(&values)
	if err != nil {
		c.AbortWithStatusJSON(500, "Please provide an array of JSON files")
	}
	result, err := collection.InsertMany(c, values)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, result.InsertedIDs)
}

func getUserCollection(client *mongo.Client) *mongo.Collection {
	return getDatabase(client).Collection("userCollection")
}
