package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v \n", index, day)
	}

	for rougeValue := 1; rougeValue < 10; {
		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}
}
