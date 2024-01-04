package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// We're going to run 1 goroutine
	wg.Add(2)

	go func() {
		count("sheep")

		// Each wg.Done() call decrement the waitgroup counter by 1
		wg.Done()
	}()

	go func() {
		count("dog")

		// Each wg.Done() call decrement the waitgroup counter by 1
		wg.Done()
	}()

	// Wait until waitgroup counter is zero
	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
