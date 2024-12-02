package main

import "fmt"

func main7() {
	// If statement
	var a, b int = 10, 20

	if a > b {
		fmt.Printf("a > b\n")
	} else if a < b {
		fmt.Printf("a < b\n")
	} else {
		fmt.Printf("a == b\n")
	}
	fmt.Print("\n")
}