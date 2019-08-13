// Sources: https://tour.golang.org/concurrency/2

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0}
	s := []int{0, 4, -9, 8, 2, 7}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	//a := &x
	//b := &y
	// x := <-c

	// y := <-c

	fmt.Println(x, y, x+y)

	// fmt.Println(x)
	// fmt.Println(y)
}
