package main

import (
	"encoding/json"
	"fmt"
)

var birdJson = `{
  "species": "pigeon",
  "description": "likes to perch on rocks"
}`

var birdJsonArray = `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`

var noStructureBird = `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`

type Bird struct {
	Species     string `json:"type"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("Json Marshall & Unmarshall in Go")

	// single json value
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Println(bird)

	// deals with the arrays of json value
	var birds []Bird
	json.Unmarshal([]byte(birdJsonArray), &birds)
	fmt.Printf("Birds : %+v", birds)

	// handling json data without any structer
	var noStrBird map[string]any
	json.Unmarshal([]byte(noStructureBird), &noStrBird)
	fmt.Println(noStrBird)

	testBird := &Bird{
		Species:     "Pigeon",
		Description: "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat.",
	}

	data, _ := json.Marshal(testBird)
	fmt.Println(string(data))

}
