package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func main() {
	tr := triangle{height: 5, base: 4}
	sq := square{sideLength: 5}

	printArea(tr)
	printArea(sq)
}
