package main

import (
	"fmt"
	"math"
)

// INTERFACES are named collections of method signatures

// Here is a basic interface for geometric shapes
type geometry interface {
	area() float64
	perim() float64
}

// For this expample, we will implement this interfact on rect and circle types
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// To implement an interface in Go, all the methods in the interface need implementing
// Implementing geometry on rects
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implementing geometry on circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
