package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingElse(ctx context.Context, printCh <-chan int) {
	for {
		select {
		// checks wheater context is canceled or not
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doSomethingElse err: %s\n", err)
			}
			fmt.Println("doSomethingElse is finished")
			return
		case num := <-printCh:
			fmt.Printf("doSomethingElse: %d\n", num)
		}
	}
}

// it is recommended to put the ctx as the first parameter
func doSomething(ctx context.Context) {
	// create a context with cancel ability
	ctx, cancelCtx := context.WithCancel(ctx)

	printCh := make(chan int)
	go doSomethingElse(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}
	cancelCtx()
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Do Something is finished")
}

func doWithDeadline(ctx context.Context) {
	deadlineTime := time.Now().Add(1500 * time.Millisecond)
	// runs the context with a deadline
	ctx, cancelCtx := context.WithDeadline(ctx, deadlineTime)
	defer cancelCtx()

	printCh := make(chan int)
	go doSomethingElse(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("doWithDeadline is finished")
}

func doWithTimeout(ctx context.Context) {
	// runs the context for the given time
	ctx, cancelCtx := context.WithTimeout(ctx, 1500*time.Millisecond)
	defer cancelCtx()

	printCh := make(chan int)
	go doSomethingElse(ctx, printCh)

	for num := 1; num <= 3; num++ {
		select {
		case printCh <- num:
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			break
		}
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)
	fmt.Println("doWithTimeout is finished")
}

// context are immutable
func main() {
	// defining a new context
	// ctx := context.TODO()
	ctx := context.Background()

	ctx = context.WithValue(ctx, "key", "This is the value of the key")

	fmt.Println("WithCanel Function")
	doSomething(ctx)

	fmt.Println("WithDeadline Function")
	doWithDeadline(ctx)

	fmt.Println("WithTimeout Function")
	doWithTimeout(ctx)
}
