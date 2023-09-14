package main

import "fmt"

func main() {
	var a int = 10;
	var p *int = &a; // p is a pointer to an integer variable
	fmt.Printf("a = %d\n", a) // value of a
	fmt.Printf("p = %d\n", p) // address of a

	a = 15 // change value of a
	fmt.Printf("Value of p = %v\n", *p) // value of p will change as well

	*p = 20 // change value of a using pointer
	fmt.Printf("Value of a = %v\n", a) // value of a will change as well
}