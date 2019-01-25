package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
)

// User structure from mongo database
type User struct {
	id        string
	firstname string
	lastname  string
	email     string
	topics    []Topic
}

func getUser(c *gin.Context) {
	s := startSession()
	cur, err := s.Database(defaultDB).Collection("userCollection").Find(nil, nil)

	if err != nil {
		log.Fatal(err)
	}
	var results []*bson.Document

	for cur.Next(nil) {
		user := bson.NewDocument()
		err = cur.Decode(user)
		results = append(results, user)
	}
	fmt.Println(results)
	c.JSON(200, results)
}
func getUserbyID(c *gin.Context) {
	s := startSession()
	id := c.Param("userID")
	user := bson.NewDocument()
	bsonID, err := objectid.FromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", bsonID))
	//sort, err := Opt.Sort(bson.NewDocument(bson.EC.ObjectID("_id", bsonID)))
	err = s.Database(defaultDB).Collection("userCollection").FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		panic(err)
	} else {
		c.JSON(200, gin.H{"data": user})
	}
	endSession(s)
}
func deleteUserByID(c *gin.Context) {
	s := startSession()
	id := c.Param("userID")
	bsonID, err := objectid.FromHex(id)
	if err != nil {
		panic(err)
	}
	filter := bson.NewDocument(bson.EC.ObjectID("_id", bsonID))
	//sort, err := Opt.Sort(bson.NewDocument(bson.EC.ObjectID("_id", bsonID)))

	result, err := s.Database(defaultDB).Collection("userCollection").DeleteOne(c, filter, nil)
	if err != nil {
		panic(err)
	} else {
		c.JSON(200, result.DeletedCount)
	}
	endSession(s)
}
func modifyUserbyID(c *gin.Context) {
}

// FOO...
type FOO struct {
	firstname string `json:"firstname"`
	lastname  string `json:"lastname"`
	email     string `json:"email"`
}

func createUser(c *gin.Context) {
	s := startSession()
	bsonID := objectid.New()
	user := bson.NewDocument(
		bson.EC.ObjectID("_id", bsonID),
		bson.EC.String("firstname", "jean guy"),
	)
	var input *FOO
	err := c.Bind(input)
	fmt.Println(*input)
	result, err := s.Database(defaultDB).Collection("userCollection").InsertOne(c, user)
	if err != nil {
		endSession(s)
		panic(err)
	}
	c.JSON(http.StatusCreated, result.InsertedID)
	// s.Database(defaultDB).Collection("userCollection").Find(nil, nil)
	// 	user.firstname = "john"
	// 	user.lastname = "Doe"
	// 	user.email = "test@gmail.com"
	// 	fmt.Println(user)

	// 	err := s.DB(defaultDB).C("userCollection").Insert(user)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	c.JSON(200, user)
	//
}
func convertDoc(d *bson.Document) gin.H {
	fmt.Println(d.LookupElement("_id").Value())
	user := gin.H{
		"id":        d.LookupElement("_id").String(),
		"firstname": d.Lookup("firstname").StringValue(),
		"lastname":  d.Lookup("lastname").StringValue(),
		"email":     d.Lookup("email").StringValue(),
		"topics":    nil,
	}
	fmt.Println(user)
	return user
}
