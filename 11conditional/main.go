package main

import "fmt"

func main() {
	fmt.Println("If Else in golang")

	loginCount := 23

	var result string

	if loginCount < 10 {
		result = "Less than 10"
	} else if loginCount > 10 {
		result = "Greater than 10"
	} else {
		result = "Exact 10"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is greater than 10")
	}

	
}
