package main

import (
	"fmt"
	"net/url"
)

const URL string = "https://mukul98s.dev:8000/learn?coursename=golang&backend=true"

func main() {
	fmt.Println(URL)

	// parsing the url
	result, _ := url.Parse(URL)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)
	// fmt.Println(result.Port())

	params := result.Query()
	fmt.Printf("Type of query params is %T\n", params)

	for _, val := range params {
		fmt.Println("Param is:", val)
	}

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "mukul98s.dev",
		Path:     "/watch",
		RawQuery: "user=mukul98s",
	}

	fmt.Println(partsOfURL.String())
}
