package main

import (
	"fmt"
	"time"
)

/*
** A goroutine is a lightweight thread of execution
 */

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Call the method the usual way (synchronous execution)
	f("direct")

	// Call the function in (TODO: as) a goroutine. This goroutine will execute
	// concurrently (asynchronously) with the calling one
	go f("goroutine")

	// goroutines can be used with anonymous functions
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Wait a second for goroutines. Obviously this is not a clever method. A
	// robust solution is WaitGroup
	time.Sleep(time.Second)
	fmt.Println("done")
}
