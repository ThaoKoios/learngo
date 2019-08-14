// Go's select lets you wait on multiple channel operations.
// Combining goroutines and channels with select is a powerful feature of GO

package main

import (
	"fmt"
	"time"
)

func main() {

	// For this example, we will select across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate e.g. blocking RPC(remote procedure call) operations executing in concurrent goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// We will use "select" to await both of these values simulataneously, printing each one as it arrives
	// We receive the values "one" and then "two" as expected
	// Note that the total execution time is only ~2 seconds since both the 1 and 2 second "Sleeps" execute concurrently
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
