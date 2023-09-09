package main

import "fmt"

/*
go has 4 types of variables:
1. bool
2. string
3. int - int8, int16, int32, int64
4. float - float32, float64

Variables are declared using the var keyword followed by the variable name and type.
*/

func main() {
	var name string = "John Doe"
	var age int = 25
	var isCool bool = true
	const size float32 = 2.3

	// Alternative
	// var name = "John Doe" // type is inferred
	// var age = 25 // type is inferred
	// var isCool = true // type is inferred
	// var size = 2.3 // type is inferred

	// Shorthand
	// these can only be used inside a function. not globally.
	// name := "John Doe"
	// age := 25
	// isCool := true
	// size := 2.3

	// another way to print
	fmt.Printf("Name: %v\nAge: %v\nIsCool: %v\nSize: %v\n", name, age, isCool, size)
}