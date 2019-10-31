package main

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DBTimeout is the maximum response time from DB
const DBTimeout = 5000

var mongoURL string

func setMongoParameters() {
	// Use one line url first
	if os.Getenv("MONGO_ONELINE_URL") != "" {
		mongoURL = os.Getenv("MONGO_ONELINE_URL")
	} else {
		//  Else compose url
		var hostname, port string

		if os.Getenv("MONGO_HOSTNAME") != "" {
			hostname = os.Getenv("MONGO_HOSTNAME")
		} else {
			customWarn("USING LOCAL DATABASE")
			hostname = "localhost"
		}
		if os.Getenv("MONGO_PORT") != "" {
			port = os.Getenv("MONGO_PORT")
		} else {
			customWarn("USING DEFAULT DATABASE")
			port = "27017"
		}
		mongoURL = "mongodb://" + hostname + ":" + port
	}

	customLog("DB: {name: mongo, url: " + mongoURL + "}")
}

func getClient(c *gin.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(c, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, errors.New("Failed to generate Mongo Client: " + err.Error())
	}

	customLog("pinging database with this url: " + mongoURL)

	// Short timeout to test mongo connection
	shortCtx, cancelFunc := context.WithTimeout(c, DBTimeout*time.Millisecond)
	defer cancelFunc()
	err = client.Ping(shortCtx, readpref.Primary())
	if err != nil {
		return nil, errors.New("Unable to reach database within " + strconv.Itoa(DBTimeout) + "ms")
	}
	customLog("Acces granted !")
	return client, nil
}

func getDatabase(c *mongo.Client) *mongo.Database {
	name := "smartsearch"
	database := c.Database(name)
	return database
}

func checkHealth(c *gin.Context) {
	var response gin.H
	client, err := getClient(c)

	if err != nil || client == nil {
		response = gin.H{"api": "true", "mongo": "false"}
		customErr("Failed to connect with Database !")
	} else {
		defer client.Disconnect(c)
		response = gin.H{"api": "true", "mongo": "true"}
	}
	c.JSON(200, response)
}
