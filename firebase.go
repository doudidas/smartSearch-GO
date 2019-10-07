package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Sets your Google Cloud Platform project ID.
const projectID = "spacelama"
const credantialsPath = "./credantials/GCP-service-account.json"

var credantials FirebaseServiceAccount
var clientOptions option.ClientOption

// FirebaseServiceAccount structure from GCP
type FirebaseServiceAccount struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUR        string `json:"client_x509_cert_url"`
}

func initFirebase() error {
	// Define Client Options from credantials file
	jsonFile, err := os.Open(credantialsPath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	// read our opened json file as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'credantials' which we defined above
	err = json.Unmarshal(byteValue, &credantials)
	if err != nil {
		return err
	}
	clientOptions = option.WithCredentialsJSON(byteValue)
	return nil
}

func init() {
	// Define Client Options from credantials file
	jsonFile, err := os.Open(credantialsPath)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	// read our opened json file as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'credantials' which we defined above
	err = json.Unmarshal(byteValue, &credantials)
	if err != nil {
		panic(err)
	}
	clientOptions = option.WithCredentialsJSON(byteValue)
	// return nil
}

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
	client, err := getFireBaseClient(c)
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
	client, err := getFireBaseClient(c)
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
	client, err := getFireBaseClient(c)
	if err != nil {
		return err
	}
	// Close client when done.
	defer client.Close()

	_, err = client.Collection("users").Doc(id).Delete(c)
	return err
}
func getUsersFromFirebase(c *gin.Context, pageNumber int, pageSize int) ([]User, error) {
	client, err := getFireBaseClient(c)
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
	client, err := getFireBaseClient(c)
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

func getFireBaseClient(c *gin.Context) (*firestore.Client, error) {
	client, err := firestore.NewClient(c, projectID, clientOptions)
	if err != nil {
		return nil, errors.New("Failed to generate Firebase Client")
	}
	return client, nil
}

func testFirebase() error {
	ctx, _ := gin.CreateTestContext(nil)
	_, err := getFireBaseClient(ctx)
	return err
}
