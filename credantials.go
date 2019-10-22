package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
)

const pathToFile = "/etc/smartsearch/apiAdmins.json"

func getAdmins() gin.Accounts {
	// Open our jsonFile
	jsonFile, err := os.Open(pathToFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println("Unable to reach " + pathToFile + " file: " + err.Error())
		fmt.Println("Using random Credantials instead...")
		return generateAdmin()
	}
	fmt.Println("Successfully Opened " + pathToFile)
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Using random Credantials instead...")
		return generateAdmin()
	}
	var result gin.Accounts
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Using random Credantials instead...")
		return generateAdmin()
	}
	return result
}

func generateAdmin() gin.Accounts {
	fmt.Println("Generating new Admin...")
	res, err := password.Generate(64, 10, 10, false, false)
	if err != nil {
		panic(err)
	}
	output := gin.Accounts{"admin": res}
	fmt.Println("Admin generated!")
	fmt.Println(output)
	return output
}
