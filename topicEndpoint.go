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

func getTopicbyID(c *gin.Context) {
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getTopicCollection(client)
	customLog(c.Param("TopicID"))
	objectID, err := primitive.ObjectIDFromHex(c.Param("TopicID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := bson.M{"_id": objectID}
	var result bson.M
	collection.FindOne(c, filter).Decode(&result)
	fmt.Println(result)

	c.JSON(200, result)
}

func getTopicNumber(c *gin.Context) {
	client, err := getClient(c)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	collection := getTopicCollection(client)
	result, err := collection.CountDocuments(c, bson.D{})
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer client.Disconnect(c)
	c.JSON(200, result)
}

func getTopics(c *gin.Context) {
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
	collection := getTopicCollection(client)
	cur, err := collection.Find(c, bson.D{{}}, findOptions)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	defer cur.Close(c)
	var result []bson.M
	for cur.Next(c) {
		var tmp bson.M
		err := cur.Decode(&tmp)
		if err != nil {
			c.AbortWithStatusJSON(500, err.Error())
		}
		result = append(result, tmp)
	}
	if err := cur.Err(); err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, result)
}
func deleteTopicByID(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getTopicCollection(client)

	objectID, err := primitive.ObjectIDFromHex(c.Param("TopicID"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Invalid input. Please check format"})
		return
	}
	filter := bson.M{"_id": objectID}
	var result bson.M
	collection.FindOneAndDelete(c, filter).Decode(&result)

	if result == nil {
		message := "Failed to remove Topic with this ID " + c.Param("TopicID")
		customLog(message)
		c.AbortWithStatusJSON(500, gin.H{"error": message})
		return
	}
	c.JSON(200, result)

}
func modifyTopicEmail(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getTopicCollection(client)

	value := bson.M{
		"$set": bson.M{
			"email": c.Param("email"),
		},
	}
	objectID, err := primitive.ObjectIDFromHex(c.Param("TopicID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
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
func modifyTopicByID(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getTopicCollection(client)

	var value bson.M
	err = c.ShouldBindJSON(&value)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	update := bson.M{
		"$set": value,
	}

	objectID, err := primitive.ObjectIDFromHex(c.Param("TopicID"))
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
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, output)
}

func createTopic(c *gin.Context) {
	client, err := getClient(c)
	defer client.Disconnect(c)
	collection := getTopicCollection(client)

	var value bson.M
	err = c.ShouldBindJSON(&value)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	result, err := collection.InsertOne(c, value)
	if err != nil {
		c.AbortWithStatusJSON(500, err.Error())
	}
	c.JSON(200, result.InsertedID)
}

func getTopicCollection(client *mongo.Client) *mongo.Collection {
	return getDatabase(client).Collection("TopicCollection")
}
