package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string) // create an empty map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	// colors[10] = "#ffffff"

	// delete(colors, 10)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
