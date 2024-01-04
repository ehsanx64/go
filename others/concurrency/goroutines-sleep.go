package main

import (
	"fmt"
	"time"
)

func main() {
	// Run two instance of count function in goroutines
	go count("sheep")
	go count("fish")

	// This causes the main function (application) wait for two seconds before
	// exiting thus, above function calls would be executed for two seconds
	time.Sleep(time.Second * 2)
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
