package main

import "fmt"

// struct is a collection of fields
type Person struct {
	name string
	age int
	id int
}

// function to print struct
func printPerson(p Person) {
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.id)
}

// function to increase age
func (p *Person) increaseAge(increament int) {
	p.age += increament
}

func main13() {
	// creating a struct object
	fahad := Person{"Fahad", 23, 1}

	// print the struct object
	fmt.Println("Printing whole struct object")
	fmt.Println(fahad)

	// specific field
	fmt.Println("\nPrinting using specific field")
	fmt.Println(fahad.name)
	fmt.Println(fahad.age)
	fmt.Println(fahad.id)

	// changing the value of a field
	fahad.increaseAge(1) // call by reference

	// printing using function
	fmt.Println("\nPrinting using function")
	printPerson(fahad)
}