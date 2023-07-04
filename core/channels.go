package main

import (
	"fmt"
	"time"
)

/*
** Channels are pipes that connect concurrent goroutines. Values can be sent from
** a goroutine and receive those values into another goroutine.
 */
func demoBasic() {
	// Create a new channel. Channels are typed by the values they send/receive
	messages := make(chan string)

	// Send a (string) value into the channel from a goroutine
	// 'channel <- ' syntax is used for sending values to a channel
	go func() { messages <- "ping" }()

	// Receive the value from messages channel
	// '<-channel' syntax is used for receiving values from a channel
	msg := <-messages

	// By default send/receive block until both the sender and the receiver are
	// ready. This property causes the program termination to wait for the "ping"
	// message without any synchronizaiton mechanism in place
	fmt.Println(msg)
}

/*
** By default channels are unbuffered, meaning they will only accept sends
** (chan <-) if there is a corresponding receive (<-chan) ready to receive the
** sent value. Buffered channels accept a limited number of values without a
** corresponding receiver for those values.
 */
func demoChannelBuffering() {
	// Make a string channel buffered up to 2 values
	messages := make(chan string, 2)

	// Since the channel is buffered, we can send values into it without a
	// corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Receive the values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
** Channels can be used to synchronize execution across goroutines. Here's an
** example of using a blocking receive to wait for a goroutine to finish. When
** waiting for multiple goroutines, a WaitGroup may be the preferred solution.
 */
func demoChannelSynchronization() {
	// Run the worker in a goroutine giving it the channel to notify back
	done := make(chan bool, 1)
	go worker(done)

	// This blocks until we receive a message on the channel (from the worker)
	// If this line removed program would terminate even before worker's start
	<-done
	fmt.Println("Quitting ...")
}

// This function will be used in a goroutine. The done channel will be used to
// notifity the other goroutine that this function execution is done
func worker(done chan bool) {
	// Let's pretend the working process takes 2 seconds
	fmt.Println("Working ...")
	time.Sleep(time.Second * 2)
	fmt.Println("Done!")

	// Send the signal
	done <- true
}

/*
** When using channels as function parameters, you can specify if a channel is
** meant to only send or receive values. This specifity increases the type-safety
 */
func demoChannelDirections() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// The ping function only accepts a channel for sending values. Trying to receive
// on the channel would cause a compile-time error.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function only accepts the pings channel for receiving values and
// the pongs channel for sending values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// Basic channel usage
	//demoBasic()

	// Buffered channel demo
	//demoChannelBuffering()

	// Channel synchronization demo
	//demoChannelSynchronization()

	// Channel directions demo
	demoChannelDirections()
}
