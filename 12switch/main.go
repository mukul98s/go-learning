package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Switch case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spot and roll the dice again")
	default:
		fmt.Println("This is an alien!")
	}
}