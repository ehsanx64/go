package main

import (
	"fmt"
	"time"
)

func main() {
	// Following function is executed as a goroutine and keeps the application
	// running so ultimately both function calls are keep executing simultaneously
	go count("sheep")
	count("fish")
}

func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
