package main

import "fmt"

// first capital letter mean PUBLIC variable
const LoginToken string = "random variable"

func main() {
	// no var syntax
	// this will not work outside the method
	username := "Mukul Sharma"
	fmt.Println(username)
	fmt.Printf("Variable is of types: %T \n", username)

	// implicit data types
	var isLoggedIn = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of types: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of types: %T \n", smallVal)

	var smallfloat float32 = 255.4555555999999999999999999999
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of types: %T \n", smallfloat)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("Variable is of types: %T \n", anothervariable)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of types: %T \n", LoginToken)

}
