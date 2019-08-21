// Source: https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/

package main

import (
	"fmt"
)

func main() {
	n := 3
	in := make(chan int)
	out := make(chan int)

	// We now supply 2 channels to the `multiplyByTwo` function
	// One for sending data and one for receiving
	go multiplyByTwo(in, out)

	// We then send it data through the channel and wait for the result
	in <- n
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	// This line is just to illustrate that there is code that is
	// excuted before we have to wait on the `in` channel
	fmt.Println("Initializing goroutine...")

	// The goroutine does not proceed until data is received on the `in` channel
	num := <-in

	// The rest is unchanged
	result := num * 2
	out <- result
}
