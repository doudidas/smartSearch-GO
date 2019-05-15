package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var hostname string

func initDB() {
	setHostname()
	err := pingMongo()
	if err != nil {
		panic(err)
	}
	monitor()
}

func setHostname() {
	if os.Getenv("MONGO-HOSTNAME") != "" {
		hostname = os.Getenv("MONGO-HOSTNAME")
	} else {
		hostname = "localhost"
	}
	customLog("mongo DB hostname set to " + hostname)
}
func getClient() *mongo.Client {
	customLog("ask for DB client")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+hostname+":27017"))
	if err != nil {
		panic(err)
	}
	customLog("DB client granted")
	return client
}

func pingMongo() error {
	client := getClient()
	ctx := gin.Context{}
	customLog("pinging database with this FQDN: " + hostname)
	shortCtx, cancelFunc := context.WithTimeout(&ctx, 1*time.Second)
	defer cancelFunc()

	err := client.Ping(shortCtx, readpref.Primary())
	defer client.Disconnect(&ctx)
	if err != nil {
		return err
	}
	customLog("pong")
	return nil
}

func getDatabase(c *mongo.Client) *mongo.Database {
	name := "GoSmartSearchDatabase"
	database := c.Database(name)
	return database
}

func monitor() {
	var o string
	c := getClient()
	d := c.Database("admin")
	d.RunCommand(context.Background(), "db.enableFreeMonitoring()").Decode(&o)
	fmt.Println(o)
}
