package main

import "fmt"

const (
	min = 0
	max = 10
)

func main() {
	var nums []int
	for i := min; i <= max; i++ {
		nums = append(nums, i)
	}

	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%d is even\n", num)
		} else {
			fmt.Printf("%d is odd\n", num)
		}
	}
}
