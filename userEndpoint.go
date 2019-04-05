package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User structure from mongo database
type User struct {
	id        primitive.ObjectID `bson:"id"`
	firstName string             `bson:"firstName"`
	lastName  string             `bson:"lastName"`
	email     string             `bson:"email"`
}

func getUserbyID(c *gin.Context) {
	client := getClient()
	defer client.Disconnect(context.Background())
	collection := getUserCollection(client)
	customLog(c.Param("userID"))
	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := primitive.M{"_id": objectID}
	var result primitive.M
	collection.FindOne(c, filter).Decode(&result)
	fmt.Println(result)

	c.JSON(200, result)
}

func getUsers(c *gin.Context) {
	client := getClient()
	defer client.Disconnect(context.Background())
	collection := getUserCollection(client)
	filter := primitive.M{}
	cur, err := collection.Find(c, filter)
	if err != nil {
		c.AbortWithError(500, err)
	}
	defer cur.Close(c)
	var result []primitive.M

	for cur.Next(c) {
		var tmp primitive.M
		err := cur.Decode(&tmp)
		if err != nil {
			c.AbortWithError(500, err)
		}
		result = append(result, tmp)
	}
	if err := cur.Err(); err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, result)
}
func deleteUserByID(c *gin.Context) {}
func modifyUserEmail(c *gin.Context) {
	client := getClient()
	defer client.Disconnect(context.Background())
	collection := getUserCollection(client)

	value := primitive.M{
		"$set": primitive.M{
			"email": c.Param("email"),
		},
	}
	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := primitive.M{"_id": objectID}
	fmt.Println(value)
	var output primitive.M
	collection.FindOneAndUpdate(context.Background(), filter, value).Decode(&output)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, output)
}
func modifyUserbyID(c *gin.Context) {
	client := getClient()
	defer client.Disconnect(context.Background())
	collection := getUserCollection(client)

	var value primitive.M
	err := c.ShouldBindJSON(&value)
	if err != nil {
		log.Fatal(err)
	}
	update := primitive.M{
		"$set": value,
	}

	objectID, err := primitive.ObjectIDFromHex(c.Param("userID"))
	if err != nil {
		c.String(500, "Invalid input. Please check format")
		return
	}
	filter := primitive.M{"_id": objectID}
	var output primitive.M

	opt := options.FindOneAndUpdateOptions{}
	opt.SetReturnDocument(options.After)

	collection.FindOneAndUpdate(c, filter, update, &opt).Decode(&output)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, output)
}

func createUser(c *gin.Context) {
	client := getClient()
	defer client.Disconnect(context.Background())
	collection := getUserCollection(client)

	var value primitive.M
	err := c.ShouldBindJSON(&value)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := collection.InsertOne(context.Background(), value)
	c.JSON(200, result.InsertedID)
}

func getUserCollection(client *mongo.Client) *mongo.Collection {
	return getDatabase(client).Collection("userCollection")
}
