package main

import "fmt"

func main() {
	var pi float32 = 3.1416
	fmt.Printf("pi = %f\n", pi) // 3.141600

	// we can also specify the precision of the float
	fmt.Printf("pi = %.2f\n", pi) // 3.14
}