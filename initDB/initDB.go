package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var fqdn string

func main() {

	if len(os.Args) == 1 {
		fmt.Println("no fqdn provided: checking localhost…")
		fqdn = "localhost"
	} else {
		fqdn = os.Args[1]
	}

	var i int
	for !apiAvailable(fqdn) {
		fmt.Println("connexion attempt #" + strconv.Itoa(i+1) + " failed ...")
		time.Sleep(2 * time.Second)
		i++
		if i == 100 {
			panic("TIMEOUT")
		}
	}

	url := "http://" + fqdn + ":9000/api/user"

	f, err := os.Open("assets/users.json")
	check(err)

	req, error := http.NewRequest("POST", url, f)
	check(error)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46Vk13YXJlMSE=")

	res, error := http.DefaultClient.Do(req)
	check(error)
	defer res.Body.Close()

	fmt.Println(res)
}

// check will interrupt current program if an error occur.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func apiAvailable(fqdn string) bool {
	url := "http://" + fqdn + ":9000/ping"
	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		fmt.Println(error)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic YWRtaW46Vk13YXJlMSE=")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		fmt.Println(error)
		return false
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body) == "pong")
	return (string(body) == "pong")
}
