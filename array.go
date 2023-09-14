package main

import "fmt"

func main() {
	// create an array of 3 integers
	var arr[3] int

	// assign values to each index position
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr) // [1 2 3]

	
	// another way to declare array
	arr2 := [3] int {1, 2, 3}
	fmt.Println(arr2)
}