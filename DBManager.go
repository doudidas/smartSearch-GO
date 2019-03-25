package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initDB(c *gin.Context) {
	client := getClient(c)
	ping(client, c)
	getDatabase(client)
}

func getClient(c *gin.Context) *mongo.Client {
	customLog("ask for DB client")

	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	customLog("DB client granted")
	return client
}
func ping(client *mongo.Client, ctx *gin.Context) {
	customLog("ping DB")
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	customLog("pong")
}

func getDatabase(c *mongo.Client) *mongo.Database {
	name := "GoSmartSearchDatabase"
	customLog("Looking for " + name)
	database := c.Database(name)
	return database
}
