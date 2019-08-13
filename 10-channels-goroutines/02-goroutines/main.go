package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call f(s). Here's how we'd call that in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// A goroutine call can also be started for an anonymous function
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln() // this scan requires we press a key before the program exits
	fmt.Println("done")
}
