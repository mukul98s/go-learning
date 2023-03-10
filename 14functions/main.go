package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in GoLang")
	greeter()

	result := adder(4, 5)
	fmt.Println("Result is: ", result)

	massResult := proAdder(1, 3, 4, 5, 6, 8)
	fmt.Println("Mass Result is: ", massResult)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

func greeter() {
	fmt.Println("Hello, Mr. Potter!")
}
