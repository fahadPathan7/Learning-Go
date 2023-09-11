
// 3 ways to get input from the user.
// these are scan, scanln, and scanf
// scan - reads a single word
// scanln - reads a line
// scanf - reads a formatted string

package main

import "fmt"

func main3() {
	var name string
	var age int

	// using scan
	fmt.Scan(&name, &age)

	fmt.Printf("Name: %v\nAge: %v\n", name, age)

	// // using scanln
	// fmt.Scanln(&name)
	// fmt.Scanln(&age)

	// fmt.Printf("Name: %v\nAge: %v\n", name, age)

	// // using scanf
	// fmt.Scanf("%v %v", &name, &age)

	// fmt.Printf("Name: %v\nAge: %v\n", name, age)
}