package main

import "fmt"

// maps are same as Hash Tables
// it store values in key value pairs
func main()  {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["GO"] = "GoLang"

	fmt.Println(languages)
	fmt.Println("GO stands for: ", languages["GO"])

	delete(languages, "JS")
	fmt.Println(languages)

	// loops in golang
	for key, value := range languages {
		fmt.Printf("For key %v, values is %v \n", key, value)
	}

}