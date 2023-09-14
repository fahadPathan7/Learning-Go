package main

import "fmt"

func main8() {
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


	// range switch
	switch {
	case a >= 0 && a <= 50:
		fmt.Printf("a >= 0 && a <= 50\n")
		break
	case a >= 51 && a <= 100:
		fmt.Printf("a >= 51 && a <= 100\n")
		break
	default:
		fmt.Printf("a = %d\n", a)
	}
}