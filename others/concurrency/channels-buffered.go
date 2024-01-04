package main

import (
	"fmt"
)

func main() {
	// Following line causes a deadlock
	// c := make(chan string)

	// So make it buffered so it buffer as many messages specified
	// Buffer size must be the mininum number of messages that are going to be
	// sent over it to prevent deadlocks
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}
