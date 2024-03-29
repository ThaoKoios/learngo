// Source: https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/

package main

import (
	"fmt"
)

func main() {
	out := make(chan int)
	in := make(chan int)

	// Create 3 `multiplyByTwo` go routines.
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)
	go multiplyByTwo(in, out)

	// Up till this point, none of the created goroutines actually do
	// anything, since they are all waiting for the `in` channel to
	// receive some data
	in <- 1
	in <- 2
	in <- 3

	// Now we wait for each result to come in
	fmt.Println(<-out)
	fmt.Println(<-out)
	fmt.Println(<-out)
}

func multiplyByTwo(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	num := <-in
	result := num * 2
	out <- result
}
