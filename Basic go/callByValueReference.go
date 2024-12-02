package main

import "fmt"

func callByValue(a int) {
	a = 20
}

func callByReference(a *int) {
	*a = 20
}

func main() {
	// call by value
	var a int = 10;
	fmt.Println("Before calling cbv, a = ", a)
	callByValue(a) // pass value of a
	fmt.Println("After calling cbv, a = ", a)

	// call by reference
	var b int = 10;
	fmt.Println("Before calling cbr, b = ", b)
	callByReference(&b) // pass address of b
	fmt.Println("After calling cbr, b = ", b)
}