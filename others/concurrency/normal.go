package main

import (
	"fmt"
	"time"
)

func main() {
	// Following line will be executed
	count("sheep")

	// because count function runs forever only above function call executes and
	// following code will never be executed
	count("fish")
}

/*
** count function displays its arguments in an infinite loop
 */
func count(thing string) {
	for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
