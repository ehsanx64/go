package main

import (
	"fmt"
	"time"
)

/*
** Select gives us the ability to wait on multiple channel operations. Combining
** goroutines and channels with select is a powerful feature of Go.
 */

func main() {
	// We're going to select across two channels
	ch1 = make(chan string)
	ch2 = make(chan string)

	// Each channel will receive a value after some amount of time, simulating
	// blocking I/O (or RPC) operations executing concurrently
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "two"
	}()

	// We'll use select to await both of these values simultaneously, printing
	// each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}
}
