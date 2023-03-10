package main

import "fmt"

func main()  {
	fmt.Println("Welcome to a class on pointer")

	// var ptr *int
	// fmt.Println("Value of Pointer is ", ptr)

	myNumber := 24

	var numPtr = &myNumber

	// pointer is reference to direct memory addres
	fmt.Println("Value of actual pointer is ", numPtr)
	// the value stored inside is the actual value
	fmt.Println("Value of actual pointer is ", *numPtr)

	*numPtr = *numPtr *2
	fmt.Println("New value is: ", myNumber)
}