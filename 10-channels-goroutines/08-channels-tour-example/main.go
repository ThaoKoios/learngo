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
	// d := make(chan int)

	go sum(s[:len(s)/2], c)
	// x := <-c
	go sum(s[len(s)/2:], c)
	// y := <-c

	// go sum(s[len(s)/2:], d)
	x, y := <-c, <-c
	//x, y := <-c, <-d // receive from c

	//a := &x
	//b := &y

	fmt.Println(x, y, x+y)
	// fmt.Println(x, y, x+y)

	// fmt.Println(x)
	// fmt.Println(y)
}
