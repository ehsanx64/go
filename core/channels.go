package main

import "fmt"

/*
** Channels are pipes that connect concurrent goroutines. Values can be sent from
** a goroutine and receive those values into another goroutine.
 */

func main() {
	// Create a new channel. Channels are typed by the values they send/receieve
	messages := make(chan string)

	// Send a (string) value into the channel from a goroutine
	// 'channel <- ' syntax is used for sending values to a channel
	go func() { messages <- "ping" }()

	// Recive the value from messages channel
	// '<-channel' syntax is used for receiving values from a channel
	msg := <-messages

	// By default send/receive block until both the sender and the receiver are
	// ready. This property causes the program termination to wait for the "ping"
	// message without any synchronizaiton mechanism in place
	fmt.Println(msg)
}
