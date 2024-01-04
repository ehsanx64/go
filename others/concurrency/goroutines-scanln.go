package main

import (
	"fmt"
	"time"
)

func main() {
	// Run two instance of count function in goroutines
	go count("sheep")
	go count("fish")

	// This function call waits for user input, thus blocks the main function's
	// exit. Until user presses enter key above functions would be executed
	// concurrently
	fmt.Scanln()
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
