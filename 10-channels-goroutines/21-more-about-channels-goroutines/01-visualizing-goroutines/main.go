// Source: https://www.sohamkamani.com/blog/2017/08/24/golang-channels-explained/

package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3

	// We want to run a goroutine to multiply n by 2
	go multiplyByTwo(n)

	// We pause the program so that the `multiplyByTwo` goroutine
	// can finish and print the output before the code exits
	time.Sleep(time.Second)
}

func multiplyByTwo(num int) int {
	result := num * 2
	fmt.Println(result)
	return result
}

// The two parts of our code (`main` and `multiplyByTwo`) are rarther disconnected. As a consequence:
// 1. We cannot access the result of `multiplyByTwo` in the `main` function
// 2. We have no way to know when the `multiplyByTwo` goroutine completes.
//	  As a result of this, we have to pause the `main` function by calling `time.Sleep`, which is a hacky solution at best.
