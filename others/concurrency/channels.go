package main

import (
	"fmt"
	"time"
)

/*
** Send/receive on channels is blocking
**
** When you try to receive something you have to wait for it be there a value to
** receive.
**
** And when you are sending a message it will wait until a receiver is ready to
** receive
 */

func main() {
	// Create a string channel
	c := make(chan string)

	// Call the count function as a goroutine, passing the string channel
	go count("sheep", c)

	// Receive a message from the channel and put it in a variable
	msg := <-c

	// Print the variable
	fmt.Println(msg)
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// Send the function argument to the channel
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}
}
