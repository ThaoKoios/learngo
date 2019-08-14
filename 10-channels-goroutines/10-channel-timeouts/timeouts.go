// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.package timeouts.
// Implementing timeouts in Go is easy and elegant thanks to channels and select.

package main

import (
	"fmt"
	"time"
)

func main() {
	// For our example, suppose we are executing an external call that returns its result on a channel c1 after 2s
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// Here is the select implementing a timeout, res := <- c1 awaits the result and <- Time.
	// After awaits a value to be sent after the timeout of 1s. since select proceeds with the first receive that's ready, we will take the timeout case if the operation takes more than the allowed 1s
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we will print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// Running this program shows that first operation timing out and the second succeeding

/*
Using this select timeout pattern requires communicationg results over channels.
This is a good idea in general because other important Go features are based on channels and select.
Other two examples of this: timers & tickers
*/
