package main

import (
	"context"
	"time"

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

var hostname = "mongo"

func getClient(c *gin.Context) *mongo.Client {
	customLog("ask for DB client")
	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://"+hostname+":27017"))

	if err != nil {
		customLog("try local database")
		hostname = "localhost"
		client, err = mongo.Connect(c, options.Client().ApplyURI("mongodb://"+hostname+":27017"))
		if err != nil {
			panic(err)
		}
	}
	customLog("DB client granted")
	return client
}

func ping(client *mongo.Client, ctx *gin.Context) error {
	customLog("pinging database with this FQDN: " + hostname)
	shortCtx, _ := context.WithTimeout(ctx, 1*time.Second)
	err := client.Ping(shortCtx, readpref.Primary())
	if err != nil {
		customLog("Switching to local database")
		hostname = "localhost"
		client = getClient(ctx)
		err = ping(client, ctx)
		if err != nil {
			panic(err)
		}
	}
	customLog("pong")
	return nil
}

func getDatabase(c *mongo.Client) *mongo.Database {
	name := "GoSmartSearchDatabase"
	customLog("Looking for " + name)
	database := c.Database(name)
	return database
}
