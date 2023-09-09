package main // package declaration. main is a special package name which tells the Go compiler that the package should compile as an executable program instead of a shared library.

import "fmt" // import declaration. fmt is a standard library package which contains functions for printing formatted output and scanning input from the console.

func main() {
	fmt.Println("Hello, World!")
	fmt.Print("I am learning go")
}