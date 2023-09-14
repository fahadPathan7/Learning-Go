package main

import "fmt"

func main() {
	var a int = 100
	// var b int = 200

	switch a {
	case 10:
		fmt.Printf("a = 10\n")
		break
	case 100:
		fmt.Printf("a = 100\n")
		break
	default:
		fmt.Printf("a = %d\n", a)
	}
}