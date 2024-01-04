package main

import (
	"fmt"
	"time"
)

func main() {
	// Becuase both functions are executed as goroutines the main function
	// (the application) exits immediately so none of these functions have a
	// chance to be executed
	go count("sheep")
	go count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
