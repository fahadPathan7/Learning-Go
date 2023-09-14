package main

import "fmt"

func main4() {
	var decimalNumber int

	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&decimalNumber)

	fmt.Printf("Binary: %b\n", decimalNumber)
	fmt.Printf("Octal: %o\n", decimalNumber)
	fmt.Printf("Hexadecimal: %x\n", decimalNumber)
}