package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines. Values can be sent into channels from one goroutine and receive those values into another goroutine

func main() {

	// Create a new channel with make(chan val-type). Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here "ping" is sent to the messages channel made above, from a new goroutine.
	go func() { messages <- "ping" }()

	// The <- channel syntax receives a value from the channel. He we will recive the "ping" message we sent above and print it out
	msg := <-messages
	fmt.Println(msg)
}
