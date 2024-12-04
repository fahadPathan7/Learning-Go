package main

import "fmt"

// a function that takes a string and prints it
func printCountry(country string) {
	fmt.Printf("Country = %s\n", country)
}

// a function that takes a string and returns it
func getSum(a int, b int) int {
	return a + b
}

func main() {
	var country string
	fmt.Print("Enter country: ")
	fmt.Scan(&country)

	printCountry(country)


	var a, b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)

	fmt.Printf("Sum = %d\n", getSum(a, b))
}