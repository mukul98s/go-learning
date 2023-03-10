package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Welcome to Slices")

	// []string -> this is slices
	var fruitList = []string{"Apple", "Kiwi", "Peach"}
	fmt.Printf("The datatype of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[0:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 345
	highScore[1] = 34
	highScore[2] = 35
	highScore[3] = 45

	highScore = append(highScore, 354,6454,656)

	sort.Ints(highScore)
	fmt.Println(highScore)

	// remove a value from slices based on index

	var index int = 2
	highScore = append(highScore[:index], highScore[index+1:]...)
	fmt.Println(highScore)
}