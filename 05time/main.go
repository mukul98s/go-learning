package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Travel with GO")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// time format is specific. Use the given date for
	// values formatting
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.July, 14, 10, 10, 10, 10, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
