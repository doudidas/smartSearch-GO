package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var databaseFQDN string
var databasePort string
var defaultDB = "SmartSearchDatabase"
var client *mongo.Client

func initMongo() {

	if len(os.Args) != 2 {
		log.Fatal("please provide input. app [databaseFQDN]")
	}
	databaseFQDN = os.Args[1]
	databasePort = "27017"
	c := startSession()
	getDatabaseInformations(c)
	endSession(c)
}
func startSession() *mongo.Client {
	fullPath := "mongodb://" + databaseFQDN + ":" + databasePort
	fmt.Println("connexion asked on ", fullPath)
	session, err := mongo.NewClient(fullPath)
	if err != nil {
		databaseFQDN = "localhost"
		fmt.Println("wrong endpoint: switching to ", fullPath)
		fullPath := "mongodb://" + databaseFQDN + ":" + databasePort
		fmt.Println("connexion asked on ", fullPath)
		session, err = mongo.NewClient(fullPath)
		if err != nil {
			panic(err)
		}
	}
	session.Connect(nil)
	// Optional. Switch the session to a monotonic behavior.
	fmt.Println("connexion granted")
	return session
}

func endSession(c *mongo.Client) {
	err := c.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("connexion closed")
}
func showCollection(c *mongo.Client, name string) {
	result, err := c.Database(defaultDB).Collection(name).Count(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database", name, ": ", result)
}
func getDatabaseInformations(c *mongo.Client) {
	fmt.Println("******DATABASE CHECK*******")
	cur, err := c.Database(defaultDB).ListCollections(nil, nil)

	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(nil) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Fatal(err)
		}
		collectionName := elem.Lookup("name").StringValue()
		showCollection(c, collectionName)
		if err != nil {
			log.Fatal("Decode error ", err)
		}
	}
	fmt.Println("***************************")
}
