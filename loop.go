package main

import "fmt"

func main() {
	// type 1
	var i int = 0
	for i < 10 {
		fmt.Printf("i = %d\n", i)
		i++
	}

	// type 2
	for j := 0; j < 10; j++ {
		fmt.Printf("j = %d\n", j)
	}

	// type 3
	for {
		fmt.Printf("infinite loop\n")
		break // if no break, it will be infinite loop
	}

	// type 4
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // skip the rest of the loop for this iteration only
		}
		fmt.Printf("i = %d\n", i)
	}

	// type 5
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto breakHere // after goto, it will not execute the rest of the loop. It will go to breakHere label
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Printf("This will not be printed\n") // because of goto

	breakHere:
	fmt.Printf("goto here\n")
}
