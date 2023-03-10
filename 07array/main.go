package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Kiwi"
	// fruitList[2] = "Apple"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veggie List is: ", vegList)
	fmt.Println("Length of Veggie List is: ", len(vegList))



}
