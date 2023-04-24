package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func basicLock() {
	fmt.Println("Basic Mutex")
	n := 0
	// defining a Mutex
	var m sync.Mutex

	go func() {
		// Starting a lock
		m.Lock()
		// ending a lock at the end of the program
		defer m.Unlock()
		nIsEven := isEven(n)

		time.Sleep(200 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, "is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	// if any other go routine call m.Lock, it will block the thread until m.Unlock is called
	go func() {
		m.Lock()
		n++
		fmt.Println(n, "from the second goroutine")
		m.Unlock()
	}()

	time.Sleep(time.Second)
}

func readWriteMutex() {
	fmt.Println("Read Write Mutex")
	n := 0
	var m sync.RWMutex

	// now, both goroutines call m.Lock() before accessing `n`
	// and call m.Unlock once they are done
	go func() {
		m.RLock()
		defer m.RUnlock()
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, "is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		m.RLock()
		defer m.RUnlock()
		nIsPositive := n > 0
		time.Sleep(5 * time.Millisecond)
		if nIsPositive {
			fmt.Println(n, "is positive")
			return
		}
		fmt.Println(n, "is not positive")
	}()

	go func() {
		m.Lock()
		n++
		m.Unlock()
	}()

	time.Sleep(time.Second)
}

func main() {
	// basicLock()
	// readWriteMutex()
}
