/*
The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools. There are a few other options for managing state though.package main
Here we'll look at using the 'sync/atomic' package for atomic counters accessed by multiple goroutines.
*/

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	// We will us an unsigned integer to represent our (always-positive) counter
	var ops uint64

	// To simulate concurrent updates, we will start 50 goroutines that each increment the counter about the once a millisecond.
	for i := 0; i < 50; i++ {
		go func() {
			for {

				// To atomically increment the counter we use AddUint64, giving it the memory address of our 'ops' counter with the '&' syntax
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// Wait a second to allow some ops to accumulate
	time.Sleep(time.Second)

	/* In order to safely use the counter while it is still being updated by other goroutines,
	we extract a copy of the current value into 'opsFinal' via 'LoadUint64'.
	As above we need to give this function the memory address '&ops' from which to fetch the value.
	*/
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

// Running the program shows that we executed about 40000 operations
// Next we will look at mutexes, another tool for managing state
