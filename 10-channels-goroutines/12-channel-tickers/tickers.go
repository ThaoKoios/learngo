/*
Timers are for when you want to do something once in the future.
Tickers are for when you wnat to do something repeatedly at regular intervals.
Here is an example of a ticker that ticks periodically until we stop it.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		Tickers use a similar mechanism to timers: a channel that is sent values.
		Here we will use the range builtin on the channel to iterate over the values as thery arrive every 500ms
	*/
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	/*
		Tickers can be stopped like timers.
		Once a ticker is stopped it will not receive any more values on its channel.
		We will stop ours after 1600ms.
	*/
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// When we run this program the ticker should tick 3 times before we stop it
