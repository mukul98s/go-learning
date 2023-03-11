package main

import "fmt"

func main() {
	defer fmt.Println("two")
	defer fmt.Println("one")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("defer", i)
	}
}
