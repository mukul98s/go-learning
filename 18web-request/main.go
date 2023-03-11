package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://api.github.com/users/mukul98s"

func main() {
	fmt.Println("Welcome to network request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the response %T \n", response)

	// always close this
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string((dataBytes))
	fmt.Printf("recieved data %s\n", content)
}
