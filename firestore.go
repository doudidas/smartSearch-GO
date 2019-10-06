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

func addUsersOnFirebase(c *gin.Context, users []User) ([]string, error) {
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

func addUserOnFirebase(c *gin.Context, user User) (string, error) {

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

func getUserFromFirebaseByID(c *gin.Context, id string) (User, error) {
	var user User
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return user, err
	}
	// Close client when done.
	defer client.Close()

	doc, err := client.Collection("users").Doc(id).Get(c)
	_ = doc.DataTo(&user)
	user.ID = doc.Ref.ID
	if err != nil {
		return user, err
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
func getUsersFromFirebase(c *gin.Context, pageNumber int, pageSize int) ([]User, error) {
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return nil, err
	}
	// Close client when done.
	defer client.Close()

	iter := client.Collection("users").Limit(pageSize).Offset(pageSize * pageNumber).Documents(c)
	var output []User
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var user User
		_ = doc.DataTo(&user)
		user.ID = doc.Ref.ID
		output = append(output, user)
	}
	return output, nil
}

func modifyUserOnFirebase(c *gin.Context, userID string, user User) error {
	var err error
	client, err := firestore.NewClient(c, projectID, option.WithCredentialsFile(credantialsPath))
	if err != nil {
		return err
	}
	// Close client when done.
	defer client.Close()
	_, err = client.Collection("users").Doc(userID).Set(c, user)
	if err != nil {
		return err
	}
	return nil
}
