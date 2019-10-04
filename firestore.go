package main

import (
	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Sets your Google Cloud Platform project ID.
const projectID = "spacelama"
const credantialsPath = "./GCP-service-account.json"

func addUsersOnFirebase(c *gin.Context, users []map[string]interface{}) ([]string, error) {
	var output []string
	for _, user := range users {
		id, err := addUserOnFirebase(c, user)
		if err != nil {
			return nil, err
		}
		output = append(output, id)
	}
	return output, nil
}

func addUserOnFirebase(c *gin.Context, user map[string]interface{}) (string, error) {

	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return "", err
	}
	// Close client when done.
	defer client.Close()
	docRef, _, err := client.Collection("users").Add(c, user)
	if err != nil {
		return "", err
	}
	return docRef.ID, nil
}

func getUserFromFirebaseByID(c *gin.Context, id string) (map[string]interface{}, error) {
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return nil, err
	}
	// Close client when done.
	defer client.Close()

	doc, err := client.Collection("users").Doc(id).Get(c)
	user := doc.Data()
	user["id"] = doc.Ref.ID
	if err != nil {
		return nil, err
	}
	return user, nil
}

func deleteUserFromFirebase(c *gin.Context, id string) error {
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return err
	}
	// Close client when done.
	defer client.Close()
	
	_, err = client.Collection("users").Doc(id).Delete(c)
	return err
}
func getUsersFromFirebase(c *gin.Context, pageNumber int, pageSize int) ([]map[string]interface{}, error) {
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return nil, err
	}
	// Close client when done.
	defer client.Close()

	iter := client.Collection("users").Limit(pageSize).Offset(pageSize * pageNumber).Documents(c)
	var output []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		user := doc.Data()
		user["id"] = doc.Ref.ID
		output = append(output, user)
	}
	return output, nil
}
