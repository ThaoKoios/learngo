package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Le"}
	// fmt.Println(alex)

	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Le"

	fmt.Println(alex)
	fmt.Printf("%+v\n", alex)
}
